package espn

type Client struct {
	LeagueID int
	Year     int
	ESPNS2   string // ESPN's main session cookie for private leagues
	SWID     string // unique ESPN user identifier
}

func NewClient(leagueID, year int) *Client {
	return &Client{
		LeagueID: leagueID,
		Year:     year,
	}
}

func NewPrivateClient(leagueID, year int, espnS2, swid string) *Client {
	return &Client{
		LeagueID: leagueID,
		Year:     year,
		ESPNS2:   espnS2,
		SWID:     swid,
	}
}
