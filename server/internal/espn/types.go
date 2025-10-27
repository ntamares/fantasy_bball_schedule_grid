package espn

type FreeAgentResponse struct {
	Players []CleanFreeAgent `json:"players"`
}

type RosterResponse struct {
	Teams []CleanFantasyTeam `json:"teams"`
}

type CleanFantasyTeam struct {
	ID      int                 `json:"id"`
	Name    string              `json:"name"`
	Abbrev  string              `json:"abbrev"`
	Players []CleanRosterPlayer `json:"players"`
}

type CleanRosterPlayer struct {
	ID                int           `json:"id"`
	Name              string        `json:"name"`
	Jersey            string        `json:"jersey"`
	Position          string        `json:"position"`
	EligiblePositions []string      `json:"eligiblePositions"`
	Team              NBATeam       `json:"team"`
	AverageStats      CleanStats    `json:"averageStats"`
	Ownership         OwnershipInfo `json:"ownership"`
	InjuryStatus      string        `json:"injuryStatus"`
	IsActive          bool          `json:"isActive"`
}

type CleanFreeAgent struct {
	ID                int           `json:"id"`
	Name              string        `json:"name"`
	Jersey            string        `json:"jersey"`
	Position          string        `json:"position"`
	EligiblePositions []string      `json:"eligiblePositions"`
	Status            string        `json:"status"`
	Team              NBATeam       `json:"team"`
	AverageStats      CleanStats    `json:"averageStats"`
	Ownership         OwnershipInfo `json:"ownership"`
	InjuryStatus      string        `json:"injuryStatus"`
	IsActive          bool          `json:"isActive"`
}

type CleanStats struct {
	Points                 float64 `json:"points"`
	PointsPerGame          float64 `json:"pointsPerGame"`
	FieldGoalsMade         float64 `json:"fieldGoalsMade"`
	FieldGoalsAttempted    float64 `json:"fieldGoalsAttempted"`
	FieldGoalPercentage    float64 `json:"fieldGoalPercentage"`
	FreeThrowsMade         float64 `json:"freeThrowsMade"`
	FreeThrowsAttempted    float64 `json:"freeThrowsAttempted"`
	FreeThrowPercentage    float64 `json:"freeThrowPercentage"`
	ThreePointersMade      float64 `json:"threePointersMade"`
	ThreePointersAttempted float64 `json:"threePointersAttempted"`
	ThreePointPercentage   float64 `json:"threePointPercentage"`
	Rebounds               float64 `json:"rebounds"`
	ReboundsPerGame        float64 `json:"reboundsPerGame"`
	OffensiveRebounds      float64 `json:"offensiveRebounds"`
	DefensiveRebounds      float64 `json:"defensiveRebounds"`
	Assists                float64 `json:"assists"`
	AssistsPerGame         float64 `json:"assistsPerGame"`
	Turnovers              float64 `json:"turnovers"`
	TurnoversPerGame       float64 `json:"turnoversPerGame"`
	AssistToTurnoverRatio  float64 `json:"assistToTurnoverRatio"`
	Steals                 float64 `json:"steals"`
	StealsPerGame          float64 `json:"stealsPerGame"`
	Blocks                 float64 `json:"blocks"`
	BlocksPerGame          float64 `json:"blocksPerGame"`
	Minutes                float64 `json:"minutes"`
	MinutesPerGame         float64 `json:"minutesPerGame"`
	PersonalFouls          float64 `json:"personalFouls"`
	GamesPlayed            float64 `json:"gamesPlayed"`
	GamesStarted           float64 `json:"gamesStarted"`
	DoubleDoubles          float64 `json:"doubleDoubles"`
	TripleDoubles          float64 `json:"tripleDoubles"`
}

type NBATeam struct {
	ID           int    `json:"id"`
	Abbreviation string `json:"abbreviation"`
	City         string `json:"city"`
	Name         string `json:"name"`
	FullName     string `json:"fullName"`
}

type PlayerStats struct {
	SeasonID        int                    `json:"seasonId"`
	SplitTypeID     int                    `json:"statSplitTypeId"`
	SourceID        int                    `json:"statSourceId"`
	ScoringPeriodID int                    `json:"scoringPeriodId"`
	AppliedAverage  float64                `json:"appliedAverage"`
	AppliedTotal    float64                `json:"appliedTotal"`
	RawAverageStats map[string]interface{} `json:"averageStats"`
	RawTotalStats   map[string]interface{} `json:"stats"`
}

func (ps *PlayerStats) GetAverageStats() map[string]float64 {
	return ConvertStats(ps.RawAverageStats)
}

func (ps *PlayerStats) GetTotalStats() map[string]float64 {
	return ConvertStats(ps.RawTotalStats)
}

type OwnershipInfo struct {
	PercentOwned   float64 `json:"percentOwned"`
	PercentStarted float64 `json:"percentStarted"`
	PercentChange  float64 `json:"percentChange"`
}

type LeagueData struct {
	FantasyTeams []FantasyTeam `json:"teams"`
}

type FantasyTeam struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Abbrev string `json:"abbrev"`
	Roster Roster `json:"roster"`
}

type Roster struct {
	Entries []RosterEntry `json:"entries"`
}

type RosterEntry struct {
	PlayerID        int             `json:"playerId"`
	PlayerPoolEntry PlayerPoolEntry `json:"playerPoolEntry"`
}

type PlayerPoolEntry struct {
	ID     int    `json:"id"`
	Player Player `json:"player"`
}

type Player struct {
	ID                int           `json:"id"`
	FullName          string        `json:"fullName"`
	FirstName         string        `json:"firstName"`
	LastName          string        `json:"lastName"`
	DefaultPositionId int           `json:"defaultPositionId"`
	EligibleSlots     []int         `json:"eligibleSlots"`
	ProTeamId         int           `json:"proTeamId"`
	Jersey            string        `json:"jersey"`
	Injured           bool          `json:"injured"`
	InjuryStatus      string        `json:"injuryStatus"`
	Active            bool          `json:"active"`
	Stats             []PlayerStats `json:"stats"`
	Ownership         OwnershipInfo `json:"ownership"`
}

type FreeAgentsData struct {
	Players []FreeAgentEntry `json:"players"`
}

type FreeAgentEntry struct {
	ID                int    `json:"id"`
	Status            string `json:"status"`
	Player            Player `json:"player"`
	OnTeamId          int    `json:"onTeamId"`
	KeepValue         int    `json:"keeperValue"`
	DraftAuctionValue int    `json:"draftAuctionValue"`
}
