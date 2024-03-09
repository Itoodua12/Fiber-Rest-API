package model

import "time"

type Product struct {
	ID           uint `json:"id" gorm:"PrimaryKey"`
	CretedAt     time.Time
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}
