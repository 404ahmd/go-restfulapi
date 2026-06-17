package models

import "time"

type Order struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	UserID     uint      `json:"user_id" gorm:"not null"`
	CustomerID uint      `json:"customer_id" gorm:"not null"`
	Status     string    `json:"status" gorm:"type:varchar(20);default:'pending'"`
	TotalPrice float64   `json:"total_price" gorm:"type:decimal(10,2)"`
	Notes      string    `json:"notes" gorm:"type:text"`
	OrderedAt  time.Time `json:"ordered_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	User     User        `json:"user" gorm:"foreignKey:UserID"`
	Customer Customer    `json:"customer" gorm:"foreignKey:CustomerID"`
	Items    []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	OrderID   uint    `json:"order_id" gorm:"not null"`
	ProductID uint    `json:"product_id" gorm:"not null"`
	Quantity  int     `json:"quantity" gorm:"not null"`
	UnitPrice float64 `json:"unit_price" gorm:"type:decimal(10,2)"`
	Subtotal  float64 `json:"subtotal" gorm:"type:decimal(10,2)"`

	Product Product `json:"product" gorm:"foreignKey:ProductID"`
}

// Request struct
type CreateOrderRequest struct {
	CustomerID uint              `json:"customer_id" binding:"required"`
	Notes      string            `json:"notes"`
	Items      []OrderItemRequest `json:"items" binding:"required,min=1,dive"`
}

type OrderItemRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=processing shipped delivered cancelled"`
}

type OrderQueryParams struct {
	Page       int    `form:"page,default=1"`
	Limit      int    `form:"limit,default=10"`
	Status     string `form:"status"`
	CustomerID uint   `form:"customer_id"`
	StartDate  string `form:"start_date"`
	EndDate    string `form:"end_date"`
}