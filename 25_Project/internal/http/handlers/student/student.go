package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

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
