package models

type Customer struct{
	ID uint `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"varchar(100);not null"`
	Fullname string `json:"fullname" gorm:"varchar(100);not null"`
	Email string `json:"email" gorm:"varchar(50);not null"`
	Age int64 `json:"age" gorm:"not null"`
	Address string `json:"address" gorm:"varchar(100);not null"`
	Gender string `json:"gender" gorm:"not null"`
}