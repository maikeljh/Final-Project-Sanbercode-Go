package structs

import "time"

type Cart struct {
	ID        int64     `json:"id"`
	ProductID int64     `json:"product_id"`
	Product   Product   `json:"product"`
	Count     int64     `json:"count"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
