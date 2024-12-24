package storage

import "github.com/ujjwalprajapati16/students-api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentByid(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
	UpdateStudent(id int64, name string, email string, age int) (types.Student, error)
	DeleteStudent(id int64) (types.Student, error)
}
