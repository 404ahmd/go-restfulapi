package models

type Customer struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"varchar(100);not null"`
	Fullname string `json:"fullname" gorm:"varchar(100);not null"`
	Email    string `json:"email" gorm:"varchar(50);not null"`
	Age      int64  `json:"age" gorm:"not null"`
	Address  string `json:"address" gorm:"varchar(100);not null"`
	Gender   string `json:"gender" gorm:"not null"`
}

type CreateCustomerRequest struct {
	Username string `json:"username" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Age      int64  `json:"age" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
}

type UpdateCustomerRequest struct{
	Username string `json:"username"`
		Fullname string `json:"fullname"`
		Email    string `json:"email"`
		Age      int64  `json:"age"`
		Address  string `json:"address"`
		Gender   string `json:"gender"`
}