package services

import (
	"fantasy-bball-schedule-grid/internal/config"
	"fantasy-bball-schedule-grid/internal/espn"
)

func GetFantasyTeamRosters() (*espn.RosterResponse, error) {
	cfg := config.Load()

	client := espn.NewPrivateClient(
		cfg.LeagueID,
		cfg.Year,
		cfg.ESPNS2,
		cfg.SWID,
	)

	return client.FetchFantasyTeamRostersClean()
}
