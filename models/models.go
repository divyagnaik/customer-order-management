package models

import "time"

type CustomerRequest struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Address     string `json:"address" validate:"required"`
}

type CustomerResponse struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type OrderRequest struct {
	ProductId string `json:"product_id" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required,min=1"`
}

type OrderResponse struct {
	Id          string    `json:"id"`
	ProductId   string    `json:"product_id"`
	UserId      string    `json:"user_id"`
	Quantity    int       `json:"quantity"`
	Status      string    `json:"status"`
	TotalPrice  float64   `json:"total_price"`
	OrderedAt   time.Time `json:"ordered_at"`
	LastUpdated time.Time `json:"last_updated"`
}

type ProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Price       float64 `json:"price" validate:"required,min=0"`
	Description string  `json:"description,omitempty"`
}

type ProductResponse struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
