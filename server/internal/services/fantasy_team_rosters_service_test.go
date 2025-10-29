package services

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"
)

// This test exercises the service-level call that retrieves CLEANED fantasy team rosters
// and verifies we receive non-empty, well-formed data from ESPN via our backend.
func TestFetchFantasyTeamRosters(t *testing.T) {
	fmt.Println("=== FETCH CLEAN FANTASY TEAM ROSTERS (service → ESPN → service) ===")

	// Call the service wrapper which returns cleaned data
	result, err := GetFantasyTeamRosters()
	if err != nil {
		t.Fatalf("service GetFantasyTeamRosters() returned error: %v", err)
	}
	if result == nil {
		t.Fatal("service GetFantasyTeamRosters() returned nil result")
	}

	// Basic shape checks
	if len(result.Teams) == 0 {
		t.Fatal("expected at least one team in cleaned roster response; got 0")
	}

	fmt.Printf("✅ Received cleaned roster response with %d teams\n", len(result.Teams))

	// Ensure output directory exists
	_ = os.MkdirAll("test_data", 0o755)
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("test_data/clean_fantasy_team_rosters_%s.json", timestamp)

	// Persist JSON for manual inspection
	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		t.Fatalf("failed to marshal cleaned roster JSON: %v", err)
	}
	if err := os.WriteFile(filename, data, 0o644); err != nil {
		t.Fatalf("failed to write cleaned roster JSON: %v", err)
	}
	fmt.Printf("💾 Saved cleaned roster JSON to: %s\n", filename)

	// Print a brief summary for quick visibility
	first := result.Teams[0]
	fmt.Printf("\n=== TEAM SAMPLE ===\n")
	fmt.Printf("Team: %s (%s) [ID:%d] | Players: %d\n", first.Name, first.Abbrev, first.ID, len(first.Players))
	maxPlayers := 5
	if len(first.Players) < maxPlayers {
		maxPlayers = len(first.Players)
	}
	for i := 0; i < maxPlayers; i++ {
		p := first.Players[i]
		fmt.Printf("  - %s (%s) | Team: %s %s | Owned: %.2f%%\n",
			p.Name, p.Position, p.Team.Abbreviation, p.Team.Name, p.Ownership.PercentOwned)
	}
	if len(first.Players) > maxPlayers {
		fmt.Printf("  ...and %d more\n", len(first.Players)-maxPlayers)
	}

	// Stronger assertions on cleaned shape
	for _, team := range result.Teams {
		if team.ID == 0 {
			t.Errorf("team has zero ID: %+v", team)
		}
		if team.Name == "" {
			t.Errorf("team has empty name: %+v", team)
		}
		if team.Abbrev == "" {
			t.Errorf("team has empty abbrev: %+v", team)
		}
		if len(team.Players) == 0 {
			t.Errorf("team %s has no players in cleaned response", team.Name)
		}
		// Spot-check a few player fields when present
		for _, p := range team.Players {
			if p.ID == 0 || p.Name == "" {
				t.Errorf("invalid player in team %s: %+v", team.Name, p)
				break
			}
			break // only need to sample the first player per team
		}

		if team.ID == 7 {
			print(team.Players)
		}
	}

	fmt.Println("\n✅ Clean fantasy team rosters test finished successfully")
}
