package models

type Game struct {
	HomeTeam Team   `json:"home"`
	AwayTeam Team   `json:"away"`
	Date     string `json:"scheduled"`
}
