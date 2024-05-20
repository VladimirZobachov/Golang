package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetAllWorkTimesByDepartment
// @Summary Get all work times by department
// @Description Get work times for a given department and optionally filter by day of the week
// @Tags department
// @Accept json
// @Produce json
// @Param id path int true "Department ID"
// @Param day_of_week query string false "Day of the Week"
// @Success 200 {object} []model.WorkTimeAPI
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "No work times found"
// @Router /department/{id} [get]
func GetAllWorkTimesByDepartment(service repository.DepartmentService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		departmentIdString, ok := vars["id"]
		if !ok {
			http.Error(writer, "Invalid department ID", http.StatusBadRequest)
			return
		}

		departmentId, err := strconv.ParseInt(departmentIdString, 10, 64)
		if err != nil {
			http.Error(writer, "Invalid department ID parameter", http.StatusBadRequest)
			return
		}

		dayOfWeek := request.URL.Query().Get("day_of_week")

		workTimes, err := service.GetAllWorkTimeByDepartment(departmentId, dayOfWeek)
		if err != nil {
			http.Error(writer, "Work times not found", http.StatusNotFound)
			return
		}

		if len(workTimes) == 0 {
			http.Error(writer, "No work times found", http.StatusNotFound)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(writer).Encode(workTimes)
		if err != nil {
			log.Println("Error encoding response:", err)
		}
	}
}

// GetDepartmentSettings
// @Summary Get department settings
// @Description Get settings by department id
// @Tags department
// @Accept json
// @Produce json
// @Param id path int true "Department ID"
// @Success 200 {object} []model.DepartmentSettings
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Department settings found"
// @Router /department/settings/{id} [get]
func GetDepartmentSettings(service repository.DepartmentService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		departmentIdString, ok := vars["id"]
		if !ok {
			http.Error(writer, "Invalid department ID", http.StatusBadRequest)
		}

		departmentId, err := strconv.ParseInt(departmentIdString, 10, 64)
		if err != nil {
			http.Error(writer, "Invalid department ID parameter", http.StatusBadRequest)
		}

		departmentSettings, err := service.GetDepartmentSettings(departmentId)
		if err != nil {
			http.Error(writer, "Department settings not found", http.StatusNotFound)
		}

		writer.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(writer).Encode(departmentSettings)
		if err != nil {
			log.Println("Error encoding response:", err)
		}
	}
}
