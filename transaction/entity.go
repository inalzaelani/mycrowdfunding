package transaction

import (
	"latihanGo/campaign"
	"latihanGo/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	PaymentUrl string
	User       user.User
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Campaign   campaign.Campaign
}
