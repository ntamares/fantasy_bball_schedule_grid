package handlers

import (
	"encoding/json"
	"fantasy-bball-schedule-grid/internal/services"
	"net/http"
)

func GetGameDatesHandler(w http.ResponseWriter, r *http.Request) {
	gameDates, err := services.GetGameDates()
	if err != nil {
		http.Error(w, "Failed to get game dates: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gameDates)
}
