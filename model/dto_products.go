package model

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Active      bool    `json:"active"`
}

type ProductRequest struct {
	ID          int     `json:"id"` // ID can be optional in requests for creation
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Category    string  `json:"category" binding:"required"`
	Active      bool    `json:"active" binding:"required"`
}

// Directly convert ProductRequest to Product
func (pr ProductRequest) Product() Product {
	return Product{
		ID:          pr.ID,
		Name:        pr.Name,
		Description: pr.Description,
		Price:       pr.Price,
		Category:    pr.Category,
		Active:      pr.Active,
	}
}
