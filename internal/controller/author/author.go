package author

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"hostess-service/internal/model"
)

// создаёшь класс (структуру) контроллера, в нём как минимум одно поле - интрефейс репозитория ТОЛЬКО с теми
// функциями, которые тебе нужны в этом контроллере. В данном случае это работа с авторами.
type controller struct {
	service authorService
	// при необходимости потом добавишь сюда логгер и т.д.
}

func New(r authorService) *controller {
	return &controller{service: r}
}

// интерфейс описываешь именно в месте использования!!
type authorService interface {
	// в идеале, этот метод должен возвращать *model.Author вместе с ошщибкой. Если метод что-то меняет, он должен
	// явно возвращать эту сущность.
	CreateAuthor(author *model.Author) error
	GetAllAuthors() ([]*model.Author, error)
	UpdateAuthor(id int64, author *model.Author) error
	DeleteAuthor(id int64) error
}

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
func (c *controller) CreateAuthor() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var author model.Author
		err := json.NewDecoder(request.Body).Decode(&author)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			log.Printf(err.Error())
			return
		}

		err = c.service.CreateAuthor(&author)
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
func (c *controller) GetAllAuthors() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		authors, err := c.service.GetAllAuthors()
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
func (c *controller) UpdateAuthor() http.HandlerFunc {
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

		err = c.service.UpdateAuthor(int64(id), &author)
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
func (c *controller) DeleteAuthor() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			log.Printf(err.Error())
			return
		}

		err = c.service.DeleteAuthor(int64(id))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusNotFound)
			log.Printf(err.Error())
			return
		}

		writer.WriteHeader(http.StatusNoContent)
	}
}
