package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"varchar(50);not null"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Role      string    `json:"role" gorm:"type:varchar(20)"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
