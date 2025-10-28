package main

import (
	"fantasy-bball-schedule-grid/internal/config"
	"fantasy-bball-schedule-grid/internal/espn"
	"fmt"
)

func main() {
	cfg := config.Load()

	client := espn.NewPrivateClient(
		cfg.LeagueID,
		cfg.Year,
		cfg.ESPNS2,
		cfg.SWID,
	)

	response, err := client.FetchFreeAgentsClean()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Show first 20 players to demonstrate team grouping
	fmt.Println("=== FREE AGENTS SORTED BY TEAM ===")
	currentTeam := ""
	teamCount := 0
	playerCount := 0

	for i, player := range response.Players {
		if i >= 30 { // Show first 30 players
			break
		}

		if player.Team.Abbreviation != currentTeam {
			if currentTeam != "" {
				fmt.Printf("  (%d players)\n\n", playerCount)
			}
			currentTeam = player.Team.Abbreviation
			teamCount++
			playerCount = 0
			fmt.Printf("🏀 %s (%s):\n", player.Team.FullName, player.Team.Abbreviation)
		}

		playerCount++
		fmt.Printf("  %d. %s (%s)\n", playerCount, player.Name, player.Position)
	}

	fmt.Printf("  (%d players)\n\n", playerCount)
	fmt.Printf("Total teams shown: %d\n", teamCount)
	fmt.Printf("Total players: %d\n", len(response.Players))
}
