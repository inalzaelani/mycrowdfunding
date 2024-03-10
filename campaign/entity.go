package campaign

import "time"

type Campaign struct {
	ID             int
	UserID         int
	Name           string
	ShortDesc      string
	Description    string
	Perk           string
	BackerCount    int
	GoalAmount     int
	CurrentAmout   int
	Slug           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	CampaignImages []CampaignImage
}

type CampaignImage struct {
	ID         int
	CampaignID int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
