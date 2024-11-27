package structures

type Public struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Tag            string  `json:"tag"`
	Owner          string  `json:"owner"`
	Contacts       string  `json:"contacts"`
	Topic          string  `json:"topic"`
	SubcriberPrice float64 `json:"subscriber_price"`
	AdPrice        float64 `json:"ad_price"`
	WantsOP        bool    `json:"wants_op"`
	Description    string  `json:"description"`
	IsSelling      bool    `json:"is_selling"`
	MonthlyUsers   int     `json:"mothly_users"`
	SalePrice      float64 `json:"sale_price"`
	IsScammer      bool    `json:"is_scammer"`
	IsVerified     bool    `json:"is_verified"`
}
