package espn

type Client struct {
	LeagueID       int
	Year           int
	ESPNS2         string // ESPN's main session cookie for private leagues
	SWID           string // unique ESPN user identifier
	EspnApiBaseUrl string
}

func NewClient(leagueID, year int, espnApiBaseUrl string) *Client {
	return &Client{
		LeagueID:       leagueID,
		Year:           year,
		EspnApiBaseUrl: espnApiBaseUrl,
	}
}

func NewPrivateClient(leagueID, year int, espnS2, swid string, espnApiBaseUrl string) *Client {
	return &Client{
		LeagueID:       leagueID,
		Year:           year,
		ESPNS2:         espnS2,
		SWID:           swid,
		EspnApiBaseUrl: espnApiBaseUrl,
	}
}
