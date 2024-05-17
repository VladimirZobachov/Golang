package router

import (
	socketio "github.com/googollee/go-socket.io"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"hostess-service/internal/controller"
	"hostess-service/internal/middleware"
	"hostess-service/internal/service"
	"log"
	"net/http"
)

func InitRouters(rs service.ReservationService, hs service.HallService, ds service.DepartmentService, as service.AuthorService, is *socketio.Server) {
	router := mux.NewRouter()
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	router.HandleFunc("/api/v1/reservations", controller.CreateReservation(rs, is)).Methods("POST")
	router.HandleFunc("/api/v1/reservations", controller.GetAllReservations(rs)).Methods("GET")
	router.HandleFunc("/api/v1/reservations/{id}", controller.GetReservation(rs)).Methods("GET")
	router.HandleFunc("/api/v1/reservations/{id}", controller.UpdateReservation(rs)).Methods("PUT")
	router.HandleFunc("/api/v1/reservations/status/{id}", controller.UpdateReservationStatus(rs)).Methods("PUT")
	router.HandleFunc("/api/v1/reservations/{id}", controller.DeleteReservation(rs)).Methods("DELETE")
	router.HandleFunc("/api/v1/department/{id}", controller.GetAllWorkTimesByDepartment(ds)).Methods("GET")
	router.HandleFunc("/api/v1/department/settings/{id}", controller.GetDepartmentSettings(ds)).Methods("GET")
	router.HandleFunc("/api/v1/halls/map", controller.GetHallsMap(hs)).Methods("GET")
	router.HandleFunc("/api/v1/halls/update", controller.HallsUpdate(hs)).Methods("GET")
	router.HandleFunc("/api/v1/authors", controller.GetAllAuthors(as)).Methods("GET")
	router.HandleFunc("/api/v1/author", controller.CreateAuthor(as)).Methods("POST")
	router.HandleFunc("/api/v1/author/{id}", controller.UpdateAuthor(as)).Methods("POST")
	router.HandleFunc("/api/v1/author/{id}", controller.DeleteAuthor(as)).Methods("POST")

	socketRouter := router.PathPrefix("/socket.io/").Subrouter()
	socketRouter.Use(middleware.Cors)
	socketRouter.Handle("/", is)

	go func() {
		is.Serve()
		defer is.Close()
	}()

	log.Fatal(http.ListenAndServe(":8080", router))
}
