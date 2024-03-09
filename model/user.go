package model

import "time"

type User struct {
	ID        uint `json:"id" gorm:"PrimaryKey"`
	CreatedAt time.Time
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
