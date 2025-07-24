package model

type Tenant struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	ContactEmail string `json:"contact_email"`
	CreatedAt    string `json:"created_at"`
}
