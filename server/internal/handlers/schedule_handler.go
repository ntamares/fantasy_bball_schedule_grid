package handlers

import (
	"encoding/json"
	"fantasy-bball-schedule-grid/internal/services"
	"net/http"
)

func GetWeeklyScheduleHandler(w http.ResponseWriter, r *http.Request) {
	schedule, err := services.GetWeeklySchedule()
	if err != nil {
		http.Error(w, "Failed to get schedule: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schedule)
}
