package transaction

import (
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
	User       user.User
	CreatedAt  time.Time
	UpdatedAte time.Time
}
