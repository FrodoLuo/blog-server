package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"createdAt"`

	Blocked   bool `json:"blocked"`
	Confirmed bool `json:"confirmed"`

	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"-"`
	Token    string `json:"-"`

	Type UserRole `json:"type"`
}

/*
UserRole users roles
*/
type UserRole = int

/*
user roles for the user
*/
const (
	EVERYONE   UserRole = iota
	REGISTERED UserRole = iota
	ADMIN      UserRole = iota
)
