package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/krishanu7/students-api/internal/storage"
	"strconv"
	"github.com/krishanu7/students-api/internal/types"
	"github.com/krishanu7/students-api/internal/utils/response"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student
		// Decode the request body into student struct it might throw an error
		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) { // If the request body is empty
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(errors.New("request body is empty")))
			return
		}
		if err != nil { // If there is an error while decoding the request body
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err));
		}
		// request validation
		if err := validator.New().Struct(student); err != nil {
			validatorErrs := err.(validator.ValidationErrors) // Type assertion
			response.WriteJson(w, http.StatusBadRequest, response.ValidatioinError(validatorErrs))
			return
		}

		lastId, err := storage.CreateStudent(
			student.ID,
			student.Name,
			student.Age,
			student.Email,
		)
		slog.Info("User created successfully", slog.String("userId", fmt.Sprint(lastId))) ;

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}


		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id") // id := r.URL.Query().Get("id")
		slog.Info("Getting a student by Id", slog.String("id", id));
		intId, err := strconv.ParseInt(id, 10, 64)

		if(err != nil) {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		student, err := storage.GetStudent(intId)

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		response.WriteJson(w, http.StatusOK, student);
	}
}

