package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"hostess-service/internal/model"
	"hostess-service/internal/service"
	"log"
	"net/http"
	"strconv"
)

// CreateAuthor godoc
// @Summary Create a new author
// @Description Create a new author with the provided information.
// @Tags author
// @Accept  json
// @Produce  json
// @Param   author   body    model.Author   true  "Create Author"
// @Success 200 {object}  model.Author
// @Failure 400 {string} string "Invalid input"
// @Router  /author [post]
func CreateAuthor(service service.AuthorService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var author model.Author
		err := json.NewDecoder(request.Body).Decode(&author)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			log.Printf(err.Error())
			return
		}

		err = service.CreateAuthor(&author)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			log.Printf(err.Error())
			return
		}

		err = json.NewEncoder(writer).Encode(author)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			log.Printf(err.Error())
			return
		}
	}
}

// GetAllAuthors
// @Summary Get all authors
// @Description get all authors
// @Tags authors
// @Accept json
// @Produce json
// @Success 200 {object} model.Author
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Authors not found"
// @Router /authors [get]
func GetAllAuthors(service service.AuthorService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		authors, err := service.GetAllAuthors()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			log.Printf(err.Error())
			return
		}
		err = json.NewEncoder(writer).Encode(authors)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			log.Printf(err.Error())
		}
	}
}

// UpdateAuthor godoc
// @Summary Update author
// @Description Update an existing author with information
// @Tags author
// @Accept json
// @Produce json
// @Param id path int true "Author ID"
// @Param author body model.Author ture "Update Author"
// @Success 200 {object} model.Author
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Author not found"
// @Router /author/{id} [put]
func UpdateAuthor(service service.AuthorService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			log.Printf(err.Error())
			return
		}

		var author model.Author
		err = json.NewDecoder(request.Body).Decode(&author)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			log.Printf(err.Error())
			return
		}

		err = service.UpdateAuthor(int64(id), &author)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusNotFound)
			log.Printf(err.Error())
			return
		}

		err = json.NewEncoder(writer).Encode(author)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			log.Printf(err.Error())
		}
	}
}

// DeleteAuthor
// @Summary Delete a author
// @Description Delete a author by its unique identifier
// @Tags author
// @Produce json
// @Param id path int true "Author ID"
// @Success 204 "Author deleted"
// Failure 404 {string} string "Author not found"
// @Router /author/{id} [delete]
func DeleteAuthor(service service.AuthorService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			log.Printf(err.Error())
			return
		}

		err = service.DeleteAuthor(int64(id))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusNotFound)
			log.Printf(err.Error())
			return
		}

		writer.WriteHeader(http.StatusNoContent)
	}
}
