package structures

import "time"

type Team struct {
	ID                  uint      `json:"id"`
	Name                string    `json:"name"`
	Owner               int64     `json:"owner"` // if anon user(check) -> create
	Contacts            string    `json:"contacts"`
	Topic               string    `json:"topic"` // Choose?? Topic model
	MinSubPrice         float64   `json:"min_sub_price"`
	MaxSubPrice         float64   `json:"max_sub_price"`
	Description         string    `json:"description"`
	BotLink             string    `json:"bot_link"` // link
	IsScummer           bool      `json:"is_scammer"`
	TeamSize            int       `json:"team_size"`
	SponsorCount        int       `json:"sponsor_count"`
	MinWithdrawalAmount int       `json:"min_withdrawal_amount"` // float $
	IsVerified          bool      `json:"is_verified"`
	RegDate             time.Time `json:"reg_date"`
}
