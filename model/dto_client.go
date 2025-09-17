package model

type Client struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Document string `json:"document"`
}

type ClientRequest struct {
	ID       int    `json:"id"` // ID can be optional in requests for creation
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Document string `json:"document" binding:"required"`
}

// Directly convert ClientRequest to Client
func (cr ClientRequest) Client() Client {
	return Client{
		ID:       cr.ID,
		Name:     cr.Name,
		Email:    cr.Email,
		Document: cr.Document,
	}
}
