package services

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// 	"path/filepath"
// 	"testing"
// 	"time"
// )

// func TestGetFreeAgents(t *testing.T) {
// 	fmt.Printf("=== MAKING REAL ESPN API CALL ===\n")
// 	fmt.Printf("Using GetFreeAgents service with real credentials from .env\n\n")

// 	// Call the actual service which loads config and makes real API call
// 	result, err := GetFreeAgents()
// 	if err != nil {
// 		t.Fatalf("Failed to fetch free agents from ESPN API: %v", err)
// 	}

// 	// Verify we got a response
// 	if result == nil {
// 		t.Fatal("Expected non-nil result from ESPN API")
// 	}

// 	fmt.Printf("✅ Successfully fetched data from ESPN API!\n")
// 	fmt.Printf("Total players returned: %d\n\n", len(result.Players))

// 	// Create test_data directory if it doesn't exist
// 	testDataDir := filepath.Join(".", "test_data")
// 	if err := os.MkdirAll(testDataDir, 0755); err != nil {
// 		t.Fatalf("Failed to create test_data directory: %v", err)
// 	}

// 	// Generate timestamp for filename
// 	timestamp := time.Now().Format("2006-01-02_15-04-05")
// 	filename := fmt.Sprintf("real_free_agents_%s.json", timestamp)
// 	filePath := filepath.Join(testDataDir, filename)

// 	// Marshal JSON with indentation
// 	jsonData, err := json.MarshalIndent(result, "", "  ")
// 	if err != nil {
// 		t.Fatalf("Failed to marshal JSON: %v", err)
// 	}

// 	// Save JSON to file
// 	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
// 		t.Fatalf("Failed to write JSON file: %v", err)
// 	}

// 	fmt.Printf("💾 Saved real ESPN data to: %s\n\n", filePath)

// 	// Print detailed information about the first few players
// 	fmt.Println("=== REAL FREE AGENTS FROM ESPN API ===")

// 	if len(result.Players) == 0 {
// 		fmt.Println("⚠️  No players returned from ESPN API")
// 		fmt.Println("This could mean:")
// 		fmt.Println("  - Authentication cookies are invalid/expired")
// 		fmt.Println("  - League settings restrict free agent visibility")
// 		fmt.Println("  - API endpoint or filters need adjustment")
// 		return
// 	}

// 	// Print first 3 players with complete data
// 	maxPlayers := 3
// 	if len(result.Players) < maxPlayers {
// 		maxPlayers = len(result.Players)
// 	}

// 	for i := 0; i < maxPlayers; i++ {
// 		player := result.Players[i]
// 		fmt.Printf("\n🏀 PLAYER #%d\n", i+1)
// 		fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
// 		fmt.Printf("Name:                %s\n", player.Name)
// 		fmt.Printf("ID:                  %d\n", player.ID)
// 		fmt.Printf("Jersey:              #%s\n", player.Jersey)
// 		fmt.Printf("Position:            %s\n", player.Position)
// 		fmt.Printf("Eligible Positions:  %v\n", player.EligiblePositions)
// 		fmt.Printf("Status:              %s\n", player.Status)
// 		fmt.Printf("Team:                %s (%s)\n", player.Team.FullName, player.Team.Abbreviation)
// 		fmt.Printf("Injury Status:       %s\n", player.InjuryStatus)
// 		fmt.Printf("Active:              %t\n", player.IsActive)

// 		// Ownership data
// 		fmt.Printf("\n📊 OWNERSHIP INFO:\n")
// 		fmt.Printf("  Owned:             %.1f%%\n", player.Ownership.PercentOwned)
// 		fmt.Printf("  Started:           %.1f%%\n", player.Ownership.PercentStarted)
// 		fmt.Printf("  Change:            %+.1f%%\n", player.Ownership.PercentChange)

// 		// Statistics
// 		stats := player.AverageStats
// 		fmt.Printf("\n📈 SEASON AVERAGES:\n")
// 		fmt.Printf("  Points:            %.1f\n", stats.Points)
// 		fmt.Printf("  Rebounds:          %.1f\n", stats.Rebounds)
// 		fmt.Printf("  Assists:           %.1f\n", stats.Assists)
// 		fmt.Printf("  Steals:            %.1f\n", stats.Steals)
// 		fmt.Printf("  Blocks:            %.1f\n", stats.Blocks)
// 		fmt.Printf("  Turnovers:         %.1f\n", stats.Turnovers)
// 		fmt.Printf("  Field Goals:       %.1f/%.1f (%.1f%%)\n",
// 			stats.FieldGoalsMade, stats.FieldGoalsAttempted, stats.FieldGoalPercentage*100)
// 		fmt.Printf("  Free Throws:       %.1f/%.1f (%.1f%%)\n",
// 			stats.FreeThrowsMade, stats.FreeThrowsAttempted, stats.FreeThrowPercentage*100)
// 		fmt.Printf("  Three Pointers:    %.1f/%.1f (%.1f%%)\n",
// 			stats.ThreePointersMade, stats.ThreePointersAttempted, stats.ThreePointPercentage*100)
// 		fmt.Printf("  Minutes:           %.1f\n", stats.Minutes)
// 		fmt.Printf("  Games Played:      %.0f\n", stats.GamesPlayed)
// 		fmt.Println()
// 	}

// 	fmt.Println("=== FIRST PLAYER JSON STRUCTURE ===")
// 	if len(result.Players) > 0 {
// 		firstPlayerJSON, err := json.MarshalIndent(result.Players[0], "", "  ")
// 		if err == nil {
// 			fmt.Println(string(firstPlayerJSON))
// 		}
// 	}
// 	fmt.Println("=== END JSON STRUCTURE ===")

// 	// Test assertions
// 	if len(result.Players) > 0 {
// 		firstPlayer := result.Players[0]
// 		if firstPlayer.Name == "" {
// 			t.Error("Expected player name to be set")
// 		}
// 		if firstPlayer.Team.FullName == "" {
// 			t.Error("Expected team full name to be set")
// 		}
// 	}

// 	fmt.Printf("\n✅ Test completed successfully! Retrieved %d real players from ESPN\n", len(result.Players))
// }

// func TestGetFreeAgentsGrouped(t *testing.T) {
// 	fmt.Println("=== MAKING REAL ESPN API CALL FOR GROUPED DATA ===")
// 	fmt.Println("Using GetFreeAgentsGrouped service with real credentials from .env")

// 	result, err := GetFreeAgentsGrouped()
// 	if err != nil {
// 		t.Fatalf("Expected no error, got %v", err)
// 	}

// 	if result == nil {
// 		t.Fatal("Expected result to not be nil")
// 	}

// 	fmt.Printf("\n✅ Successfully fetched grouped data from ESPN API!\n")
// 	fmt.Printf("Total teams with free agents: %d\n", len(result.TeamGroups))

// 	// Calculate total players across all teams
// 	totalPlayers := 0
// 	for _, teamGroup := range result.TeamGroups {
// 		totalPlayers += len(teamGroup.Players)
// 	}
// 	fmt.Printf("Total players across all teams: %d\n", totalPlayers)

// 	// Save grouped data to JSON file
// 	timestamp := time.Now().Format("2006-01-02_15-04-05")
// 	filename := fmt.Sprintf("test_data/real_grouped_free_agents_%s.json", timestamp)

// 	jsonData, err := json.MarshalIndent(result, "", "  ")
// 	if err != nil {
// 		t.Fatalf("Failed to marshal JSON: %v", err)
// 	}

// 	err = os.WriteFile(filename, jsonData, 0644)
// 	if err != nil {
// 		t.Fatalf("Failed to write file: %v", err)
// 	}

// 	fmt.Printf("\n💾 Saved real grouped ESPN data to: %s\n", filename)

// 	fmt.Println("\n=== TEAMS WITH FREE AGENTS (Sorted Alphabetically) ===")
// 	for i, teamGroup := range result.TeamGroups {
// 		fmt.Printf("\n🏀 TEAM #%d: %s (%s)\n", i+1, teamGroup.Team.FullName, teamGroup.Team.Abbreviation)
// 		fmt.Printf("   Players: %d\n", len(teamGroup.Players))

// 		// Show first few players for this team
// 		maxPlayers := 3
// 		if len(teamGroup.Players) < maxPlayers {
// 			maxPlayers = len(teamGroup.Players)
// 		}

// 		for j := 0; j < maxPlayers; j++ {
// 			player := teamGroup.Players[j]
// 			fmt.Printf("   - %s (%s) - %s - %.1f%% owned\n",
// 				player.Name, player.Position, player.Status, player.Ownership.PercentOwned)
// 		}

// 		if len(teamGroup.Players) > maxPlayers {
// 			fmt.Printf("   - ... and %d more players\n", len(teamGroup.Players)-maxPlayers)
// 		}
// 	}

// 	// Test assertions
// 	if len(result.TeamGroups) == 0 {
// 		t.Error("Expected at least one team group")
// 	}

// 	// Verify teams are sorted alphabetically by FullName
// 	for i := 1; i < len(result.TeamGroups); i++ {
// 		prevTeam := result.TeamGroups[i-1].Team.FullName
// 		currentTeam := result.TeamGroups[i].Team.FullName
// 		if prevTeam > currentTeam {
// 			t.Errorf("Teams not sorted alphabetically: %s should come after %s", prevTeam, currentTeam)
// 		}
// 	}

// 	// Verify each team has at least one player
// 	for _, teamGroup := range result.TeamGroups {
// 		if len(teamGroup.Players) == 0 {
// 			t.Errorf("Team %s has no players", teamGroup.Team.FullName)
// 		}
// 	}

// 	fmt.Printf("\n✅ Grouped test completed successfully! Retrieved %d teams with players from ESPN\n", len(result.TeamGroups))
// }
