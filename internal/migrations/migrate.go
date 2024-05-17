package migrations

import (
	"gorm.io/gorm"
	"hostess-service/internal/model"
	"log"
)

func Migrate(db *gorm.DB) {

	err := db.AutoMigrate(
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
