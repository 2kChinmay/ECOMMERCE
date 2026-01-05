package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/2kChinmay/students-api/internal/types"
	"github.com/2kChinmay/students-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Creating a new student")
		var student types.Student
		// Here you would typically parse the request body to populate the student struct
		// and then save it to a database.

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		}

		//validation
		if err := validator.New().Struct(student); err != nil {
			validationErr := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidaionError(validationErr))
			return
		}
	
		slog.Info("Student created", slog.String("name", student.Name), slog.Int("age", student.Age), slog.String("email", student.Email))
		result := make(map[string]string)
		result["success"] = "Ok"
		response.WriteJson(w, http.StatusCreated, result)
	}
}
