package repository

import (
	"log"

	"hostess-service/internal/model"
)

func (r *repo) CreateAuthor(author *model.Author) error {
	return r.db.Create(author).Error
}

func (r *repo) GetAllAuthors() ([]*model.Author, error) {
	var authors []*model.Author
	err := r.db.Find(&authors).Error
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (r *repo) UpdateAuthor(id int64, author *model.Author) error {
	result := r.db.Model(&model.Author{}).Where("id = ?", id).Updates(author)
	log.Println(result.Error)
	return result.Error
}

func (r *repo) DeleteAuthor(id int64) error {
	result := r.db.Delete(&model.Author{}, id)
	return result.Error
}
