package structures

import "time"

type User struct {
	ID      int       `json:"id"`
	TID     int64     `json:"telegram_id"`
	Name    string    `json:"name"`
	RegDate time.Time `json:"registration_date"`
	Rating  float64   `json:"rating"`
}

type Team struct {
	ID                  uint      `json:"id"`
	Name                string    `json:"name"`
	Owner               User      `json:"owner"`
	Contacts            string    `json:"contacts"`
	Topic               string    `json:"topic"`
	MinSubPrice         float64   `json:"min_sub_price"`
	MaxSubPrice         float64   `json:"max_sub_price"`
	Description         string    `json:"description"`
	BotLink             string    `json:"bot_link"`
	IsScummer           bool      `json:"is_scammer"`
	TeamSize            int       `json:"team_size"`
	SponsorCount        int       `json:"sponsor_count"`
	MinWithdrawalAmount int       `json:"min_withdrawal_amount"`
	IsVerified          bool      `json:"is_verified"`
	RegDate             time.Time `json:"reg_date"`
}

type Public struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	Tag            string    `json:"tag"`
	Owner          int64     `json:"owner"` //TID
	Contacts       string    `json:"contacts"`
	Topic          string    `json:"topic"` // string -> int
	SubcriberPrice float64   `json:"subscriber_price"`
	AdPrice        float64   `json:"ad_price"`
	WantsOP        bool      `json:"wants_op"`
	Description    string    `json:"description"`
	IsSelling      bool      `json:"is_selling"`
	MonthlyUsers   int       `json:"mothly_users"`
	SalePrice      float64   `json:"sale_price"`
	IsScammer      bool      `json:"is_scammer"`
	IsVerified     bool      `json:"is_verified"`
	RegDate        time.Time `json:"reg_date"`
}
