package repository

import (
	"log"

	"gorm.io/gorm"

	"hostess-service/internal/model"
)

type repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *repo {
	return &repo{db: db}
}

// так как миграции у тебя завязаны на конкретном репозитории, то и должны быть в пакете с репозиторием
func (r *repo) Migrate() {
	err := r.db.AutoMigrate(
		model.Reservation{},
		model.Hall{},
		model.Table{},
		model.Department{},
		model.WorkTime{},
		model.Holiday{},
		model.Author{},
	)
	if err != nil {
		log.Fatalf("error migrating model: %v", err)
	}

}
