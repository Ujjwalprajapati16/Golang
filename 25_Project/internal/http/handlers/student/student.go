package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/ujjwalprajapati16/students-api/internal/storage"
	"github.com/ujjwalprajapati16/students-api/internal/types"
	"github.com/ujjwalprajapati16/students-api/internal/utils/response"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Creating a student...")
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralErro(fmt.Errorf("request body is empty")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralErro(err))
			return
		}

		// validate request
		if err := validator.New().Struct(student); err != nil {
			validatorErr := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidateError(validatorErr))
			return
		}

		// create student in db
		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)

		slog.Info("Created a student", slog.Int64("id", lastId))

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralErro(err))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}

// student get by id
func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("Getting a student...", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Error("Invalid ID format", slog.String("id", id))
			response.WriteJson(w, http.StatusBadRequest, response.GeneralErro(fmt.Errorf("invalid ID format")))
			return
		}

		student, err := storage.GetStudentByid(intId)
		if err != nil {
			slog.Error("Failed to get student", slog.String("error", err.Error()))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralErro(err))
			return
		}

		response.WriteJson(w, http.StatusOK, student)
	}
}

// student list
func GetList(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Getting a list of students...")
		students, err := storage.GetStudents()
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralErro(err))

			return
		}

		response.WriteJson(w, http.StatusOK, students)
	}
}

// Update a student
func Update(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("Updating a student...", slog.String("id", id))

		// Parse and validate the ID from the request path
		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Error("Invalid ID format", slog.String("id", id))
			response.WriteJson(w, http.StatusBadRequest, response.GeneralErro(fmt.Errorf("invalid ID format")))
			return
		}

		// Decode the updated student details from the request body
		var updatedStudent types.Student
		err = json.NewDecoder(r.Body).Decode(&updatedStudent)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralErro(fmt.Errorf("request body is empty")))
			return
		}
		if err != nil {
			slog.Error("Failed to decode request body", slog.String("error", err.Error()))
			response.WriteJson(w, http.StatusBadRequest, response.GeneralErro(err))
			return
		}

		// Validate the updated student details
		if err := validator.New().Struct(updatedStudent); err != nil {
			validatorErr := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidateError(validatorErr))
			return
		}

		if updatedStudent.Name == "" && updatedStudent.Email == "" && updatedStudent.Age == 0 {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralErro(fmt.Errorf("no fields to update")))
			return
		}

		// Call the storage layer to update the student in the database
		student, err := storage.UpdateStudent(intId, updatedStudent.Name, updatedStudent.Email, updatedStudent.Age)
		if err != nil {
			slog.Error("Failed to update student", slog.String("error", err.Error()))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralErro(err))
			return
		}

		// Respond with the updated student details
		response.WriteJson(w, http.StatusOK, student)
	}
}

// Delete a student
func Delete(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("Deleting a student...", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Error("Invalid ID format", slog.String("id", id))
			response.WriteJson(w, http.StatusBadRequest, response.GeneralErro(fmt.Errorf("invalid ID format")))
			return
		}

		student, err := storage.DeleteStudent(intId)
		if err != nil {
			slog.Error("Failed to delete student", slog.String("error", err.Error()))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralErro(err))
			return
		}

		response.WriteJson(w, http.StatusOK, student)
	}
}
