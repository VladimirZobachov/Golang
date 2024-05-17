package controller

import (
	"encoding/json"
	"hostess-service/internal/service"
	"log"
	"net/http"
)

// GetHallsMap
// @Summary Get all halls with their tables
// @Description Get a list of all halls and their associated tables
// @Tags halls
// @Accept json
// @Produce json
// @Success 200 {array} model.HallResponse
// @Failure 500 {string} string "Internal Server Error"
// @Router /halls/map [get]
func GetHallsMap(hs service.HallService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		halls, err := hs.GetHallsMap()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(writer).Encode(halls)
		if err != nil {
			http.Error(writer, "Failed to encode response", http.StatusInternalServerError)
			log.Println("Error encoding response:", err)
		}
	}
}

func HallsUpdate(hs service.HallService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		response := hs.ImportHallsFromGoulash()
		if !response.Success {
			log.Fatalf("Failed to update halls from API: %s", response.ErrorMessage)
		}
		log.Println("Halls and tables updated successfully!")
		writer.WriteHeader(http.StatusOK) // 200 OK
		writer.Write([]byte("Halls and tables updated successfully!"))
	}
}
