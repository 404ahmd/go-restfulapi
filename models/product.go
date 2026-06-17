package models

import "time"

type Product struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type varchar(100);not null"`
	Price     float64   `json:"price" gorm:"type decimal(10,2);not null"`
	Stock     int       `json:"stock" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
