package app

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"hostess-service/config"
	"hostess-service/internal/controller/author"
	"hostess-service/internal/middleware"
	"hostess-service/internal/repository"

	"hostess-service/internal/websocket"
	"hostess-service/pkg/mysql"
)

func Run(cfg *config.Config) {
	repo := repository.New(mysql.SetupDatabase(cfg))
	repo.Migrate()

	router := mux.NewRouter()

	authorController := author.New(repo)
	router.HandleFunc("/api/v1/authors", authorController.GetAllAuthors()).Methods("GET")
	router.HandleFunc("/api/v1/author", authorController.CreateAuthor()).Methods("POST")
	router.HandleFunc("/api/v1/author/{id}", authorController.UpdateAuthor()).Methods("POST")
	router.HandleFunc("/api/v1/author/{id}", authorController.DeleteAuthor()).Methods("POST")

	// ты сюда передаёшь весь репозиторий, а там использешь только методы интерфейса.

	reservationService := repository.NewReservationService(db)
	hallService := repository.NewHallService(db, cfg.Goulash)
	departmentService := repository.NewDepartmentService(db)
	authorService := repository.NewAuthorRepo(db)

	ioServer := websocket.SetupSocketIO()
	socketRouter := router.PathPrefix("/socket.io/").Subrouter()
	socketRouter.Use(middleware.Cors)
	socketRouter.Handle("/", ioServer)

	go func() {
		ioServer.Serve()
		defer ioServer.Close()
	}()

	if err := http.ListenAndServe(":8080", router); err != nil {
		// проверяешь, что сервер не завершился штатно
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}
}

//func initRouters(rs repository.ReservationService, hs repository.HallService, ds repository.DepartmentService, as repository.AuthorService, is *socketio.Server) {
//	router := mux.NewRouter()
//	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
//	router.HandleFunc("/api/v1/reservations", controller.CreateReservation(rs, is)).Methods("POST")
//	router.HandleFunc("/api/v1/reservations", controller.GetAllReservations(rs)).Methods("GET")
//	router.HandleFunc("/api/v1/reservations/{id}", controller.GetReservation(rs)).Methods("GET")
//	router.HandleFunc("/api/v1/reservations/{id}", controller.UpdateReservation(rs)).Methods("PUT")
//	router.HandleFunc("/api/v1/reservations/status/{id}", controller.UpdateReservationStatus(rs)).Methods("PUT")
//	router.HandleFunc("/api/v1/reservations/{id}", controller.DeleteReservation(rs)).Methods("DELETE")
//	router.HandleFunc("/api/v1/department/{id}", controller.GetAllWorkTimesByDepartment(ds)).Methods("GET")
//	router.HandleFunc("/api/v1/department/settings/{id}", controller.GetDepartmentSettings(ds)).Methods("GET")
//	router.HandleFunc("/api/v1/halls/map", controller.GetHallsMap(hs)).Methods("GET")
//	router.HandleFunc("/api/v1/halls/update", controller.HallsUpdate(hs)).Methods("GET")
//	router.HandleFunc("/api/v1/authors", controller.GetAllAuthors(as)).Methods("GET")
//	router.HandleFunc("/api/v1/author", controller.CreateAuthor(as)).Methods("POST")
//	router.HandleFunc("/api/v1/author/{id}", controller.UpdateAuthor(as)).Methods("POST")
//	router.HandleFunc("/api/v1/author/{id}", controller.DeleteAuthor(as)).Methods("POST")
//

//}
