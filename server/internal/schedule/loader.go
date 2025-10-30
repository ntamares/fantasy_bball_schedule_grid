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

func getMondayDate() time.Time {
	currentDate := time.Now().UTC()
	currentDay := int(currentDate.Weekday())
	daysFromMonday := (currentDay + 6) % 7
	mondayDate := currentDate.AddDate(0, 0, -daysFromMonday)

	return mondayDate
}

func FetchGameDates() []time.Time {
	dates := make([]time.Time, 0, 7)
	monday := getMondayDate()

	for i := 0; i < 7; i++ {
		nextDate := monday.AddDate(0, 0, i)
		year, month, day := nextDate.Date()
		dateTrunc := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		dates = append(dates, dateTrunc)
	}

	return dates
}

func FetchWeeklySchedule() ([]models.Game, error) {
	path := os.Getenv("SCHEDULE_JSON_PATH")
	data, err := os.ReadFile(path)
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

	gameDates := FetchGameDates()
	games := make([]models.Game, 0)

	for _, game := range schedule.Games {
		gameDate, err := time.Parse(time.RFC3339, game.Date)
		if err != nil {
			log.Printf("Error parsing JSON date: %v", err)
			return nil, fmt.Errorf("failed to parse game date '%s': %w", game.Date, err)
		}

		// TODO replace local JSON with source from US
		// SportRadar is based in Switzerland causing the off-by-one for teh dates
		gameDate = gameDate.AddDate(0, 0, -1) // quick and dirty solution
		year, month, day := gameDate.Date()
		gameDateTrunc := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

		if slices.Contains(gameDates, gameDateTrunc) {
			game.Date = gameDateTrunc.Format(time.RFC3339)
			games = append(games, game)
		}
	}
	return games, nil
}
