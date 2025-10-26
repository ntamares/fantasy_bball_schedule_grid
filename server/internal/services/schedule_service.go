package services

import (
	"fantasy-bball-schedule-grid/internal/models"
	"fantasy-bball-schedule-grid/internal/schedule"
)

func GetWeeklySchedule() ([]models.Game, error) {
	return schedule.FetchWeeklySchedule()
}
