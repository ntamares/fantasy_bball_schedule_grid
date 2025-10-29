package services

import (
	"fantasy-bball-schedule-grid/internal/config"
	"fantasy-bball-schedule-grid/internal/espn"
)

func GetFantasyTeamRosters() (*espn.RosterResponse, error) {
	cfg := config.Load()

	client := espn.NewClient(
		cfg.LeagueID,
		cfg.Year,
		cfg.EspnApiBaseUrl,
	)

	return client.FetchFantasyTeamRostersClean()
}
