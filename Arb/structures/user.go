package structures

import (
	"time"
)

type User struct {
	ID      uint      `json:"id"`
	TID     int64     `json:"telegram_id"`
	Name    string    `json:"name"`
	RegDate time.Time `json:"registration_date"`
	Rating  float64   `json:"rating"`
}
