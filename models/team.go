package models

type Team struct {
	CaptainID    string `json:"captainId" gorm:"primary_key"`
	TournamentID int    `json:"tournamentId" gorm:"primary_key"`
	TeamName     string `json:"teamName" gorm:"not null"`
	Members      string `json:"members"`
}

type TeamInput struct {
	CaptainID    string   `json:"captainId"`
	TournamentID int      `json:"tournamentId"`
	TeamName     string   `json:"teamName"`
	Members      []string `json:"members"`
}

type TeamLeave struct {
	CaptainID    string `json:"captainId"`
	TournamentID int    `json:"tournamentId"`
}
