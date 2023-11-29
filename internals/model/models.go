package model

import "time"

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreateCategory struct {
	Name string `json:"name"`
}

type Product struct {
	Id              int       `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Brand           string    `json:"brand"`
	Sku             string    `json:"sku"`
	InStock         bool      `json:"in_stock"`
	Image           string    `json:"image"`
	Price           float64   `json:"price"`
	DiscountedPrice float64   `json:"discounted_price"`
	CategoryId      int       `json:"category_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
