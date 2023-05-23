package item

type CreateItemInput struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Sold        bool    `json:"sold" validate:"required"`
	UserID      string  `json:"user_id" validate:"required"`
}

type UpdateItemInput struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Sold        bool    `json:"sold"`
	UserID      string  `json:"user_id"`
}
