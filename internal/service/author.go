package service

import (
	"gorm.io/gorm"
	"hostess-service/internal/model"
	"log"
)

type AuthorService interface {
	CreateAuthor(author *model.Author) error
	GetAllAuthors() ([]*model.Author, error)
	UpdateAuthor(id int64, author *model.Author) error
	DeleteAuthor(id int64) error
}

type authorService struct {
	db *gorm.DB
}

func NewAuthorService(db *gorm.DB) *authorService {
	return &authorService{db: db}
}

func (s *authorService) CreateAuthor(author *model.Author) error {
	return s.db.Create(author).Error
}

func (s *authorService) GetAllAuthors() ([]*model.Author, error) {
	var authors []*model.Author
	err := s.db.Find(&authors).Error
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (s *authorService) UpdateAuthor(id int64, author *model.Author) error {
	result := s.db.Model(&model.Author{}).Where("id = ?", id).Updates(author)
	log.Println(result.Error)
	return result.Error
}

func (s *authorService) DeleteAuthor(id int64) error {
	result := s.db.Delete(&model.Author{}, id)
	return result.Error
}
