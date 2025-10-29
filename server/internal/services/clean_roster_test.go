package services

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"testing"
// 	"time"
// )

// func TestFetchCleanFantasyTeamRosters(t *testing.T) {
// 	fmt.Println("=== TESTING CLEAN FANTASY TEAM ROSTERS WITH LINEUP SLOTS ===")
// 	fmt.Println("Using GetFantasyTeamRosters with real credentials from .env")
// 	fmt.Println()

// 	// Fetch clean roster data with lineup slots
// 	response, err := GetFantasyTeamRosters()
// 	if err != nil {
// 		t.Fatalf("❌ Failed to fetch clean roster data: %v", err)
// 	}

// 	fmt.Printf("✅ Successfully fetched clean fantasy team rosters!\n")
// 	fmt.Printf("Total fantasy teams: %d\n", len(response.Teams))
// 	fmt.Println()

// 	// Save clean JSON to file with timestamp
// 	timestamp := time.Now().Format("2006-01-02_15-04-05")
// 	filename := fmt.Sprintf("clean_fantasy_team_rosters_%s.json", timestamp)
// 	testDataDir := "test_data"

// 	// Create test_data directory if it doesn't exist
// 	if err := os.MkdirAll(testDataDir, 0755); err != nil {
// 		log.Printf("⚠️ Warning: Could not create test_data directory: %v", err)
// 	}

// 	cleanJSON, err := json.MarshalIndent(response, "", "  ")
// 	if err != nil {
// 		t.Fatalf("❌ Failed to marshal clean JSON: %v", err)
// 	}

// 	filePath := filepath.Join(testDataDir, filename)
// 	if err := os.WriteFile(filePath, cleanJSON, 0644); err != nil {
// 		log.Printf("⚠️ Warning: Could not save clean roster data to file: %v", err)
// 	} else {
// 		fmt.Printf("💾 Saved clean fantasy team rosters data to: %s\n", filePath)
// 	}
// 	fmt.Println()

// 	// Test data structure
// 	if len(response.Teams) == 0 {
// 		t.Fatal("❌ No teams found in response")
// 	}

// 	fmt.Println("=== CLEAN ROSTER STRUCTURE TEST ===")

// 	// Test first team
// 	firstTeam := response.Teams[0]
// 	fmt.Printf("🏆 FIRST TEAM: %s (%s)\n", firstTeam.Name, firstTeam.Abbrev)
// 	fmt.Printf("   Team ID: %d\n", firstTeam.ID)
// 	fmt.Printf("   Roster Size: %d players\n", len(firstTeam.Players))

// 	if len(firstTeam.Players) == 0 {
// 		t.Fatal("❌ No players found in first team")
// 	}

// 	// Test lineup slot assignments
// 	fmt.Println("\n=== LINEUP SLOT ASSIGNMENTS ===")
// 	startingLineup := 0
// 	benchPlayers := 0
// 	utilityPlayers := 0
// 	injuredReserve := 0

// 	for i, player := range firstTeam.Players {
// 		if i < 10 { // Show first 10 players
// 			lineupInfo := "No Assignment"
// 			if player.LineupSlotId != nil {
// 				lineupInfo = fmt.Sprintf("Slot %d (%s)", *player.LineupSlotId, player.LineupSlot)
// 			}

// 			fmt.Printf("   %d. %s (%s) - %s %s - %s\n",
// 				i+1,
// 				player.Name,
// 				player.Position,
// 				player.Team.Abbreviation,
// 				player.Team.Name,
// 				lineupInfo)
// 		}

// 		// Count lineup assignments
// 		if player.LineupSlotId != nil {
// 			switch *player.LineupSlotId {
// 			case 0, 1, 2, 3, 4, 5, 6: // Starting lineup positions
// 				startingLineup++
// 			case 11: // Utility
// 				utilityPlayers++
// 			case 12: // Bench
// 				benchPlayers++
// 			case 13: // Injured Reserve
// 				injuredReserve++
// 			}
// 		}
// 	}

// 	fmt.Printf("\n📊 LINEUP BREAKDOWN:\n")
// 	fmt.Printf("   Starting Lineup: %d players\n", startingLineup)
// 	fmt.Printf("   Utility: %d players\n", utilityPlayers)
// 	fmt.Printf("   Bench: %d players\n", benchPlayers)
// 	fmt.Printf("   Injured Reserve: %d players\n", injuredReserve)

// 	// Test that we have lineup slot data
// 	playersWithSlots := 0
// 	for _, player := range firstTeam.Players {
// 		if player.LineupSlotId != nil {
// 			playersWithSlots++
// 		}
// 	}

// 	fmt.Printf("   Players with lineup assignments: %d/%d\n", playersWithSlots, len(firstTeam.Players))

// 	if playersWithSlots == 0 {
// 		t.Error("❌ No players have lineup slot assignments")
// 	} else {
// 		fmt.Printf("✅ Successfully retrieved lineup slot assignments!\n")
// 	}

// 	fmt.Println("\n=== TEST COMPLETED SUCCESSFULLY ===")
// }
