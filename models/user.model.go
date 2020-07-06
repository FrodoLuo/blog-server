package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"createdAt"`

	Blocked   bool   `json:"blocked"`
	Confirmed bool   `json:"confirmed"`
	Email     string `json:"email"`
	Username  string `json:"username"`
}
