package espn

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetRosters() (*LeagueData, error) {
	url := fmt.Sprintf("https://lm-api-reads.fantasy.espn.com/apis/v3/games/fba/seasons/%d/segments/0/leagues/%d?view=mTeam&view=mRoster", c.Year, c.LeagueID)

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

	var leagueData LeagueData
	if err := json.NewDecoder(resp.Body).Decode(&leagueData); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &leagueData, nil
}

func (c *Client) GetRawRosters() ([]byte, error) {
	url := fmt.Sprintf("https://lm-api-reads.fantasy.espn.com/apis/v3/games/fba/seasons/%d/segments/0/leagues/%d?view=mTeam&view=mRoster", c.Year, c.LeagueID)

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
