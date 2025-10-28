package espn

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var (
	PositionMap = map[int]string{
		1: "PG",
		2: "SG",
		3: "SF",
		4: "PF",
		5: "C",
	}

	// stat mappings from ESPN API
	StatsMap = map[string]string{
		"0":  "PTS",  // Points
		"1":  "BLK",  // Blocks
		"2":  "STL",  // Steals
		"3":  "AST",  // Assists
		"4":  "OREB", // Offensive Rebounds
		"5":  "DREB", // Defensive Rebounds
		"6":  "REB",  // Total Rebounds
		"7":  "EJ",   // Ejections
		"8":  "FF",   // Flagrant Fouls
		"9":  "PF",   // Personal Fouls
		"10": "TF",   // Technical Fouls
		"11": "TO",   // Turnovers
		"12": "DQ",   // Disqualifications
		"13": "FGM",  // Field Goals Made
		"14": "FGA",  // Field Goals Attempted
		"15": "FTM",  // Free Throws Made
		"16": "FTA",  // Free Throws Attempted
		"17": "3PM",  // Three Pointers Made
		"18": "3PA",  // Three Pointers Attempted
		"19": "FG%",  // Field Goal Percentage
		"20": "FT%",  // Free Throw Percentage
		"21": "3PT%", // Three Point Percentage
		"22": "AFG%", // Adjusted Field Goal Percentage
		"23": "FGMI", // Field Goals Missed
		"24": "FTMI", // Free Throws Missed
		"25": "3PMI", // Three Pointers Missed
		"26": "APG",  // Assists Per Game
		"27": "BPG",  // Blocks Per Game
		"28": "MPG",  // Minutes Per Game
		"29": "PPG",  // Points Per Game
		"30": "RPG",  // Rebounds Per Game
		"31": "SPG",  // Steals Per Game
		"32": "TOPG", // Turnovers Per Game
		"33": "3PG",  // Three Pointers Per Game
		"34": "PPM",  // Points Per Minute
		"35": "A/TO", // Assist to Turnover Ratio
		"36": "STR",  // Strength
		"37": "DD",   // Double Doubles
		"38": "TD",   // Triple Doubles
		"39": "QD",   // Quadruple Doubles
		"40": "MIN",  // Minutes Played
		"41": "GS",   // Games Started
		"42": "GP",   // Games Played
		"43": "TW",   // Team Wins
		"44": "FTR",  // Free Throw Rate
		"45": "UNK",  // Unknown/Unused
	}

	NBATeamsMap = map[int]NBATeam{
		0:  {ID: 0, Abbreviation: "FA", City: "Free", Name: "Agent", FullName: "Free Agent"},
		1:  {ID: 1, Abbreviation: "ATL", City: "Atlanta", Name: "Hawks", FullName: "Atlanta Hawks"},
		2:  {ID: 2, Abbreviation: "BOS", City: "Boston", Name: "Celtics", FullName: "Boston Celtics"},
		3:  {ID: 3, Abbreviation: "NOP", City: "New Orleans", Name: "Pelicans", FullName: "New Orleans Pelicans"},
		4:  {ID: 4, Abbreviation: "CHI", City: "Chicago", Name: "Bulls", FullName: "Chicago Bulls"},
		5:  {ID: 5, Abbreviation: "CLE", City: "Cleveland", Name: "Cavaliers", FullName: "Cleveland Cavaliers"},
		6:  {ID: 6, Abbreviation: "DAL", City: "Dallas", Name: "Mavericks", FullName: "Dallas Mavericks"},
		7:  {ID: 7, Abbreviation: "DEN", City: "Denver", Name: "Nuggets", FullName: "Denver Nuggets"},
		8:  {ID: 8, Abbreviation: "DET", City: "Detroit", Name: "Pistons", FullName: "Detroit Pistons"},
		9:  {ID: 9, Abbreviation: "GSW", City: "Golden State", Name: "Warriors", FullName: "Golden State Warriors"},
		10: {ID: 10, Abbreviation: "HOU", City: "Houston", Name: "Rockets", FullName: "Houston Rockets"},
		11: {ID: 11, Abbreviation: "IND", City: "Indiana", Name: "Pacers", FullName: "Indiana Pacers"},
		12: {ID: 12, Abbreviation: "LAC", City: "LA", Name: "Clippers", FullName: "LA Clippers"},
		13: {ID: 13, Abbreviation: "LAL", City: "Los Angeles", Name: "Lakers", FullName: "Los Angeles Lakers"},
		14: {ID: 14, Abbreviation: "MIA", City: "Miami", Name: "Heat", FullName: "Miami Heat"},
		15: {ID: 15, Abbreviation: "MIL", City: "Milwaukee", Name: "Bucks", FullName: "Milwaukee Bucks"},
		16: {ID: 16, Abbreviation: "MIN", City: "Minnesota", Name: "Timberwolves", FullName: "Minnesota Timberwolves"},
		17: {ID: 17, Abbreviation: "BKN", City: "Brooklyn", Name: "Nets", FullName: "Brooklyn Nets"},
		18: {ID: 18, Abbreviation: "NYK", City: "New York", Name: "Knicks", FullName: "New York Knicks"},
		19: {ID: 19, Abbreviation: "ORL", City: "Orlando", Name: "Magic", FullName: "Orlando Magic"},
		20: {ID: 20, Abbreviation: "PHL", City: "Philadelphia", Name: "76ers", FullName: "Philadelphia 76ers"},
		21: {ID: 21, Abbreviation: "PHO", City: "Phoenix", Name: "Suns", FullName: "Phoenix Suns"},
		22: {ID: 22, Abbreviation: "POR", City: "Portland", Name: "Trail Blazers", FullName: "Portland Trail Blazers"},
		23: {ID: 23, Abbreviation: "SAC", City: "Sacramento", Name: "Kings", FullName: "Sacramento Kings"},
		24: {ID: 24, Abbreviation: "SAS", City: "San Antonio", Name: "Spurs", FullName: "San Antonio Spurs"},
		25: {ID: 25, Abbreviation: "OKC", City: "Oklahoma City", Name: "Thunder", FullName: "Oklahoma City Thunder"},
		26: {ID: 26, Abbreviation: "UTA", City: "Utah", Name: "Jazz", FullName: "Utah Jazz"},
		27: {ID: 27, Abbreviation: "WAS", City: "Washington", Name: "Wizards", FullName: "Washington Wizards"},
		28: {ID: 28, Abbreviation: "TOR", City: "Toronto", Name: "Raptors", FullName: "Toronto Raptors"},
		29: {ID: 29, Abbreviation: "MEM", City: "Memphis", Name: "Grizzlies", FullName: "Memphis Grizzlies"},
		30: {ID: 30, Abbreviation: "CHA", City: "Charlotte", Name: "Hornets", FullName: "Charlotte Hornets"},
	}
)

func GetPositionName(positionID int) string {
	if name, exists := PositionMap[positionID]; exists {
		return name
	}
	return "Unknown"
}

func GetNBATeam(teamID int) NBATeam {
	if team, exists := NBATeamsMap[teamID]; exists {
		return team
	}
	return NBATeam{ID: teamID, Abbreviation: "UNK", City: "Unknown", Name: "Team", FullName: "Unknown Team"}
}

func convertStats(rawStats map[string]interface{}) map[string]float64 {
	convertedStats := make(map[string]float64)
	for key, value := range rawStats {
		if statName, exists := StatsMap[key]; exists && statName != "" {
			if floatVal, ok := value.(float64); ok {
				convertedStats[statName] = floatVal
			}
		}
	}
	return convertedStats
}

func convertToCleanResponse(rawData *freeAgentsData) *FreeAgentResponse {
	cleanPlayers := make([]CleanFreeAgent, 0, len(rawData.Players))

	for _, entry := range rawData.Players {
		player := entry.Player
		cleanPlayer := CleanFreeAgent{
			ID:                player.ID,
			Name:              player.FullName,
			Position:          GetPositionName(player.DefaultPositionId),
			EligiblePositions: getEligiblePositions(player.EligibleSlots),
			Status:            formatStatus(entry.Status),
			Team:              GetNBATeam(player.ProTeamId),
			Jersey:            player.Jersey,
			AverageStats:      convertPlayerStats(player.Stats),
			Ownership:         player.Ownership,
			InjuryStatus:      player.InjuryStatus,
			IsActive:          player.Active,
		}

		cleanPlayers = append(cleanPlayers, cleanPlayer)
	}

	return &FreeAgentResponse{
		Players: cleanPlayers,
	}
}

func formatStatus(status string) string {
	switch status {
	case "FREEAGENT":
		return "Free Agent"
	case "WAIVERS":
		return "Waivers"
	default:
		return status
	}
}

func convertPlayerStats(rawStats []playerStats) CleanStats {
	stats := CleanStats{}

	if len(rawStats) == 0 {
		return stats
	}

	avgStats := rawStats[0].FetchAverageStats()
	statMappings := []struct {
		espnKey   string
		statField *float64
	}{
		{"PTS", &stats.Points},
		{"PPG", &stats.PointsPerGame},
		{"FGM", &stats.FieldGoalsMade},
		{"FGA", &stats.FieldGoalsAttempted},
		{"FG%", &stats.FieldGoalPercentage},
		{"FTM", &stats.FreeThrowsMade},
		{"FTA", &stats.FreeThrowsAttempted},
		{"FT%", &stats.FreeThrowPercentage},
		{"3PM", &stats.ThreePointersMade},
		{"3PA", &stats.ThreePointersAttempted},
		{"3PT%", &stats.ThreePointPercentage},
		{"REB", &stats.Rebounds},
		{"RPG", &stats.ReboundsPerGame},
		{"OREB", &stats.OffensiveRebounds},
		{"DREB", &stats.DefensiveRebounds},
		{"AST", &stats.Assists},
		{"APG", &stats.AssistsPerGame},
		{"TO", &stats.Turnovers},
		{"TOPG", &stats.TurnoversPerGame},
		{"A/TO", &stats.AssistToTurnoverRatio},
		{"STL", &stats.Steals},
		{"SPG", &stats.StealsPerGame},
		{"BLK", &stats.Blocks},
		{"BPG", &stats.BlocksPerGame},
		{"MIN", &stats.Minutes},
		{"MPG", &stats.MinutesPerGame},
		{"PF", &stats.PersonalFouls},
		{"GP", &stats.GamesPlayed},
		{"GS", &stats.GamesStarted},
		{"DD", &stats.DoubleDoubles},
		{"TD", &stats.TripleDoubles},
	}

	for _, mapping := range statMappings {
		if val, exists := avgStats[mapping.espnKey]; exists {
			*mapping.statField = val
		}
	}

	return stats
}

func getEligiblePositions(eligibleSlots []int) []string {
	positionSet := make(map[string]bool)

	for _, slot := range eligibleSlots {
		switch slot {
		case 0:
			positionSet["PG"] = true
		case 1:
			positionSet["SG"] = true
		case 2:
			positionSet["SF"] = true
		case 3:
			positionSet["PF"] = true
		case 4:
			positionSet["C"] = true
		}
	}

	positions := []string{}
	order := []string{"PG", "SG", "SF", "PF", "C"}
	for _, pos := range order {
		if positionSet[pos] {
			positions = append(positions, pos)
		}
	}

	return positions
}
func (c *Client) FetchFreeAgentsClean() (*FreeAgentResponse, error) {
	rawData, err := c.FetchFreeAgents()
	if err != nil {
		return nil, err
	}

	return convertToCleanResponse(rawData), nil
}

func (c *Client) FetchFreeAgents() (*freeAgentsData, error) {
	rawJSON, err := c.FetchRawFreeAgents()
	if err != nil {
		return nil, err
	}

	var data freeAgentsData
	if err := json.Unmarshal(rawJSON, &data); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &data, nil
}

func (c *Client) FetchRawFreeAgents() ([]byte, error) {
	baseURL := fmt.Sprintf("https://lm-api-reads.fantasy.espn.com/apis/v3/games/fba/seasons/%d/segments/0/leagues/%d", c.Year, c.LeagueID)

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	q := u.Query()
	q.Add("view", "kona_player_info")
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if c.ESPNS2 != "" && c.SWID != "" {
		req.AddCookie(&http.Cookie{Name: "espn_s2", Value: c.ESPNS2})
		req.AddCookie(&http.Cookie{Name: "SWID", Value: c.SWID})
	}

	filters := map[string]interface{}{
		"players": map[string]interface{}{
			"filterStatus": map[string]interface{}{
				"value": []string{"FREEAGENT", "WAIVERS"},
			},
			"filterProTeamIds": map[string]interface{}{
				"value": func() []int {
					ids := make([]int, 30)
					for i := 0; i < 30; i++ {
						ids[i] = i + 1
					}
					return ids
				}(),
			},
			"sortPercOwned": map[string]interface{}{
				"sortPriority": 1,
				"sortAsc":      false,
			},
			"sortDraftRanks": map[string]interface{}{
				"sortPriority": 100,
				"sortAsc":      true,
				"value":        "STANDARD",
			},
		},
	}

	filterJSON, err := json.Marshal(filters)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal filters: %w", err)
	}
	req.Header.Set("x-fantasy-filter", string(filterJSON))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	return body, nil
}
