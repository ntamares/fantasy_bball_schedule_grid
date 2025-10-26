package services

import (
	"errors"
	"fantasy-bball-schedule-grid/internal/schedule"
	"time"
)

func GetGameDates() ([]time.Time, error) {
	gameDates := schedule.FetchGameDates()

	if len(gameDates) == 0 {
		return nil, errors.New("No games scheduled")
	}

	return gameDates, nil
}
