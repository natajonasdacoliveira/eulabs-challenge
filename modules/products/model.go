package products

import "time"

type Product struct {
	ID          uint64     `json:"id" db:"id"`
	Name        string     `json:"name" db:"name" form:"name" valid:"type(string),required"`
	Price       float64    `json:"price" db:"price" form:"price" valid:"type(float64),required"`
	Description string     `json:"description" db:"description"`
	CreatedAt   time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt   *time.Time `json:"deletedAt" db:"deleted_at"`
}

type ProductDTO struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

type BaseDTO struct {
	Data interface{} `json:"data"`
}
