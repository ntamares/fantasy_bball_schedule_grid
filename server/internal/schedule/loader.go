package schedule

import (
	"encoding/json"
	"fantasy-bball-schedule-grid/internal/models"
	"fmt"
	"log"
	"os"
	"slices"
	"time"
)

func GetMondayDate() time.Time {
	currentDate := time.Now()
	currentDay := int(currentDate.Weekday())
	daysFromMonday := (currentDay + 6) % 7
	mondayDate := currentDate.AddDate(0, 0, -daysFromMonday)

	return mondayDate
}

func GetWeeklyScheduleDates() []time.Time {
	dates := make([]time.Time, 0, 7)
	monday := GetMondayDate()

	for i := 0; i < 7; i++ {
		nextDate := monday.AddDate(0, 0, i)
		year, month, day := nextDate.Date()
		dateTrunc := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		dates = append(dates, dateTrunc)
	}

	return dates
}

func FetchGames() ([]models.Game, error) {
	data, err := os.ReadFile("data/nba_schedule_2025_2026.json")

	if err != nil {
		log.Printf("Error loading JSON: %v", err)
		return nil, fmt.Errorf("failed to load JSON: %w", err)
	}

	var schedule models.Schedule
	err = json.Unmarshal(data, &schedule)

	if err != nil {
		log.Printf("Error parsing JSON: %v", err)
		return nil, fmt.Errorf("failed to parse schedule from JSON: %w", err)
	}

	weeklyScheduleDates := GetWeeklyScheduleDates()
	games := make([]models.Game, 0)

	for _, game := range schedule.Games {
		gameDate, err := time.Parse(time.RFC3339, game.Date)

		if err != nil {
			log.Printf("Error parsing JSON date: %v", err)
			return nil, fmt.Errorf("failed to parse game date '%s': %w", game.Date, err)
		}

		year, month, day := gameDate.Date()
		gameDateTrunc := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

		if slices.Contains(weeklyScheduleDates, gameDateTrunc) {
			games = append(games, game)
		}
	}
	return games, nil
}
