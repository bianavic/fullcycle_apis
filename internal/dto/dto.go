package dto

// objeto para receber dados de fora
type CreateProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
