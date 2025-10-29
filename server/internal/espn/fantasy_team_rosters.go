package espn

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchFantasyTeamRosters() (*LeagueData, error) {
	rawJSON, err := c.FetchRawFantasyTeamRosters()
	if err != nil {
		return nil, err
	}

	var leagueData LeagueData
	if err := json.Unmarshal(rawJSON, &leagueData); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &leagueData, nil
}

func (c *Client) FetchRawFantasyTeamRosters() ([]byte, error) {
	url := fmt.Sprintf("%s/seasons/%d/segments/0/leagues/%d?view=mTeam&view=mRoster", c.EspnApiBaseUrl, c.Year, c.LeagueID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if c.ESPNS2 != "" && c.SWID != "" {
		req.AddCookie(&http.Cookie{Name: "espn_s2", Value: c.ESPNS2})
		req.AddCookie(&http.Cookie{Name: "SWID", Value: c.SWID})
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}

func (c *Client) FetchFantasyTeamRostersClean() (*RosterResponse, error) {
	leagueData, err := c.FetchFantasyTeamRosters()
	if err != nil {
		return nil, err
	}

	return convertToCleanRosterResponse(leagueData), nil
}

func convertToCleanRosterResponse(rawData *LeagueData) *RosterResponse {
	cleanTeams := make([]CleanFantasyTeam, 0, len(rawData.FantasyTeams))

	for _, team := range rawData.FantasyTeams {
		cleanPlayers := make([]CleanRosterPlayer, 0, len(team.Roster.Entries))

		for _, entry := range team.Roster.Entries {
			player := entry.PlayerPoolEntry.Player
			cleanPlayer := CleanRosterPlayer{
				ID:                player.ID,
				Name:              player.FullName,
				Position:          GetPositionName(player.DefaultPositionId),
				EligiblePositions: getEligiblePositions(player.EligibleSlots),
				Team:              GetNBATeam(player.ProTeamId),
				Jersey:            player.Jersey,
				AverageStats:      convertPlayerStats(player.Stats),
				Ownership:         player.Ownership,
				InjuryStatus:      player.InjuryStatus,
				IsActive:          player.Active,
			}

			cleanPlayers = append(cleanPlayers, cleanPlayer)
		}

		cleanTeam := CleanFantasyTeam{
			ID:      team.ID,
			Name:    team.Name,
			Abbrev:  team.Abbrev,
			Players: cleanPlayers,
		}

		cleanTeams = append(cleanTeams, cleanTeam)
	}

	return &RosterResponse{
		Teams: cleanTeams,
	}
}
