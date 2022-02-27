package campaign

import "time"

type campaign struct {
	ID              int
	UserId          int
	Name            string
	SortDescription string
	Description     string
	Perks           string
	BackerCount     int
	GoalAmount      int
	CurrentAmount   int
	Slug            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
