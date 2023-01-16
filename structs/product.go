package structs

import "time"

type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	CategoryID  int64     `json:"category_id"`
	Price       uint64    `json:"price"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
