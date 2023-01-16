package structs

import "time"

type User struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
