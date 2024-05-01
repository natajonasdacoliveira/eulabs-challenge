package products

import "time"

type Product struct {
	ID          int        `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Price       float64    `json:"price" db:"price"`
	Description string     `json:"description" db:"description"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
}

type ProductDTO struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type BaseDTO struct {
	Data interface{} `json:"data"`
}
