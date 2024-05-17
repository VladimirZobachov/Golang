package controller

import (
	"encoding/json"
	socketio "github.com/googollee/go-socket.io"
	"github.com/gorilla/mux"
	"hostess-service/internal/model"
	"hostess-service/internal/service"
	"log"
	"net/http"
	"strconv"
)

// CreateReservation handler
// @Summary Create a new reservation
// @Description Create a new reservation with the provided information.
// @Tags reservations
// @Accept json
// @Produce json
// @Param reservation body model.ReservationAPI true "Create Reservation"
// @Success 200 {object} model.ReservationResponse
// @Failure 400 {string} string "Invalid input"
// @Router /reservations [post]
func CreateReservation(service service.ReservationService, socketServer *socketio.Server) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		var reservation model.Reservation
		if err := json.NewDecoder(request.Body).Decode(&reservation); err != nil {
			log.Println(err)
			http.Error(writer, `{"error_message": "`+err.Error()+`"}`, http.StatusBadRequest)
			return
		}

		conflicts, err := service.IsAvailableTimeSlot(reservation.TableID, reservation.TimeStart, reservation.TimeFinish)
		if err != nil {
			http.Error(writer, `{"error_message": "`+err.Error()+`"}`, http.StatusInternalServerError)
			return
		}

		if len(conflicts) > 0 {
			response := model.ReservationResponse{
				Success:              false,
				ErrorMessage:         "Time slot conflicts exist",
				ConflictReservations: conflicts,
				Reservation:          nil,
			}

			if err := json.NewEncoder(writer).Encode(response); err != nil {
				log.Println("Error encoding conflict response:", err)
				http.Error(writer, `{"error_message": "Failed to encode conflict response"}`, http.StatusInternalServerError)
			}
			return
		}

		if err := service.CreateReservation(&reservation); err != nil {
			http.Error(writer, `{"error_message": "`+err.Error()+`"}`, http.StatusInternalServerError)
			log.Println(err)
			return
		}

		// Notify connected clients in the relevant department room about the new reservation
		departmentID := strconv.Itoa(reservation.DepartmentID)
		socketServer.BroadcastToRoom("/", departmentID, "reservations", reservation)

		response := model.ReservationResponse{
			Success:              true,
			ErrorMessage:         "",
			ConflictReservations: nil,
			Reservation:          &reservation,
		}

		if err := json.NewEncoder(writer).Encode(response); err != nil {
			log.Println("Error encoding response:", err)
			http.Error(writer, `{"error_message": "Failed to encode response"}`, http.StatusInternalServerError)
		}
	}
}

// GetReservation godoc
// @Summary Get a reservation by ID
// @Description Get a reservation by its uniq identifier.
// @Tags reservations
// @Accept  json
// @Produce  json
// @Param   id path int true "Reservation ID"
// @Success 200 {object}  model.Reservation
// @Failure 400 {string} string "Reservation not found"
// @Router  /reservations/{id} [get]
func GetReservation(service service.ReservationService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return
		}

		reservation, err := service.GetReservationByID(int64(id))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusNotFound)
			log.Println(err)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(writer).Encode(reservation)
		if err != nil {
			log.Println("Error encoding response:", err)
		}
	}
}

// GetAllReservations
// @Summary Get all reservations
// @Description Get all reservations
// @Tags reservations
// @Accept json
// @Produce json
// @Success 200 {object} model.ReservationResponse
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Reservation not found"
// @Router /reservations [get]
func GetAllReservations(service service.ReservationService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Parse date_from and date_to query parameters
		query := request.URL.Query()
		dateFromString := query.Get("date_from")
		dateToString := query.Get("date_to")

		var dateFrom, dateTo int64
		var err error

		if dateFromString != "" {
			dateFrom, err = strconv.ParseInt(dateFromString, 10, 64)
			if err != nil {
				http.Error(writer, "Invalid date_from parameter", http.StatusBadRequest)
				return
			}
		}

		if dateToString != "" {
			dateTo, err = strconv.ParseInt(dateToString, 10, 64)
			if err != nil {
				http.Error(writer, "Invalid date_to parameter", http.StatusBadRequest)
				return
			}
		}

		// Fetch reservations with the provided date range
		reservations, err := service.GetAllReservations(dateFrom, dateTo)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusNotFound)
			return
		}

		// Encode the reservations into JSON and send the response
		writer.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(writer).Encode(reservations)
		if err != nil {
			log.Println("Error encoding response:", err)
		}
	}
}

// UpdateReservation handler
// @Summary Update a reservation
// @Description Update an existing reservation with provided information
// @Tags reservations
// @Accept json
// @Produce json
// @Param id path int true "Reservation ID"
// @Param reservation body model.ReservationAPI true "Update Reservation"
// @Success 200 {object} model.ReservationResponse
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Reservation not found"
// @Router /reservations/{id} [put]
func UpdateReservation(service service.ReservationService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(request)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(writer, `{"error_message": "`+err.Error()+`"}`, http.StatusBadRequest)
			log.Println(err)
			return
		}

		var reservation model.Reservation
		if err := json.NewDecoder(request.Body).Decode(&reservation); err != nil {
			http.Error(writer, `{"error_message": "`+err.Error()+`"}`, http.StatusBadRequest)
			log.Println(err)
			return
		}

		conflicts, err := service.IsAvailableTimeSlot(reservation.TableID, reservation.TimeStart, reservation.TimeFinish, int64(id))
		if err != nil {
			http.Error(writer, `{"error_message": "`+err.Error()+`"}`, http.StatusInternalServerError)
			return
		}

		if len(conflicts) > 0 {
			response := model.ReservationResponse{
				Success:              true,
				ErrorMessage:         "Time slot conflicts exist",
				ConflictReservations: conflicts,
				Reservation:          nil,
			}

			if err := json.NewEncoder(writer).Encode(response); err != nil {
				log.Println("Error encoding conflict response:", err)
				http.Error(writer, `{"error_message": "Failed to encode conflict response"}`, http.StatusInternalServerError)
			}
			return
		}

		err = service.UpdateReservation(int64(id), &reservation)
		if err != nil {
			http.Error(writer, `{"error_message": "`+err.Error()+`"}`, http.StatusNotFound)
			log.Println(err)
			return
		}

		response := model.ReservationResponse{
			Success:              true,
			ErrorMessage:         "",
			ConflictReservations: nil,
			Reservation:          &reservation,
		}

		if err := json.NewEncoder(writer).Encode(response); err != nil {
			log.Println("Error encoding response:", err)
			http.Error(writer, `{"error_message": "Failed to encode response"}`, http.StatusInternalServerError)
		}
	}
}

// DeleteReservation
// @Summary Delete a reservation
// @Description Delete a reservation by its unique identifier
// @Tags reservations
// @Produce json
// @Param id path int true "Reservation ID"
// @Success 204 "Reservation deleted"
// Failure 404 {string} string "Reservation not found"
// @Router /reservations/{id} [delete]
func DeleteReservation(service service.ReservationService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(writer, "Invalid reservation ID", http.StatusBadRequest)
			log.Println(err)
			return
		}

		err = service.DeleteReservationByID(int64(id))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusNotFound)
			log.Println(err)
			return
		}

		writer.WriteHeader(http.StatusNoContent)
	}
}
