package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ujjwalprajapati16/students-api/internal/config"
	"github.com/ujjwalprajapati16/students-api/internal/types"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS students (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL,
			age INTEGER NOT NULL
		);`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{Db: db}, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO students (name, email, age) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (s *Sqlite) GetStudentByid(id int64) (types.Student, error) {
	// Prepare the query statement
	stmt, err := s.Db.Prepare("SELECT id, name, email, age FROM students WHERE id = ?")
	if err != nil {
		return types.Student{}, fmt.Errorf("failed to prepare query: %w", err)
	}
	defer stmt.Close()

	var student types.Student

	// Execute the query and scan the result
	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("student with id %d not found", id)
		}
		return types.Student{}, fmt.Errorf("failed to query student with id %d: %w", id, err)
	}

	return student, nil
}

func (s *Sqlite) GetStudents() ([]types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id, name, email, age FROM students")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare query: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("failed to query students: %w", err)
	}

	defer rows.Close()

	var students []types.Student

	for rows.Next() {
		var student types.Student
		if err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age); err != nil {
			return nil, fmt.Errorf("failed to scan student: %w", err)
		}
		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over students: %w", err)
	}

	return students, nil
}

// Update a student
func (s *Sqlite) UpdateStudent(id int64, name string, email string, age int) (types.Student, error) {
	// Prepare the SQL update statement
	stmt, err := s.Db.Prepare("UPDATE students SET name = ?, email = ?, age = ? WHERE id = ?")
	if err != nil {
		return types.Student{}, fmt.Errorf("failed to prepare query: %w", err)
	}
	defer stmt.Close()

	// Execute the update query with the provided values
	_, err = stmt.Exec(name, email, age, id)
	if err != nil {
		return types.Student{}, fmt.Errorf("failed to update student with id %d: %w", id, err)
	}

	// Fetch the updated student data
	updatedStudent, err := s.GetStudentByid(id)
	if err != nil {
		return types.Student{}, fmt.Errorf("failed to fetch updated student with id %d: %w", id, err)
	}

	return updatedStudent, nil
}

// Delete a student
func (s *Sqlite) DeleteStudent(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare("DELETE FROM students WHERE id = ?")
	if err != nil {
		return types.Student{}, fmt.Errorf("failed to prepare query: %w", err)
	}
	defer stmt.Close()

	student, err := s.GetStudentByid(id)
	if err != nil {
		return types.Student{}, err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return types.Student{}, fmt.Errorf("failed to delete student with id %d: %w", id, err)
	}

	return student, nil
}
