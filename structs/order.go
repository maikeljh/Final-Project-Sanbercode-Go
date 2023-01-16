package structs

import "time"

type Order struct {
	ID        int64     `json:"id"`
	CartID    int64     `json:"cart_id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
