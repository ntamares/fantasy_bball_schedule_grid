package main

import (
	"fantasy-bball-schedule-grid/internal/schedule"
	"fmt"
	"log"
)

func main() {
	games, err := schedule.FetchGames()
	if err != nil {
		log.Fatalf("Error fetching games: %v", err)
	}
	fmt.Println(games)
}
