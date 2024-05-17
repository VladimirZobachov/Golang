package app

import (
	"hostess-service/config"
	"hostess-service/internal/migrations"
	"hostess-service/internal/router"
	"hostess-service/internal/service"
	"hostess-service/internal/websocket"
	"hostess-service/pkg/mysql"
)

func Run(cfg *config.Config) {

	db := mysql.SetupDatabase(cfg)
	migrations.Migrate(db)
	reservationService := service.NewReservationService(db)
	hallService := service.NewHallService(db, cfg.Goulash)
	departmentService := service.NewDepartmentService(db)
	ioServer := websocket.SetupSocketIO()
	authorService := service.NewAuthorService(db)
	router.InitRouters(reservationService, hallService, departmentService, authorService, ioServer)
}
