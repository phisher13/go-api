package entity

type ProductDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}
