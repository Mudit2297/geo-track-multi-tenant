package utils

type LoginResponse struct {
	Response struct {
		AccessToken       string      `json:"AccessToken"`
		ExpiresIn         int         `json:"ExpiresIn"`
		IDToken           string      `json:"IdToken"`
		NewDeviceMetadata interface{} `json:"NewDeviceMetadata"`
		RefreshToken      string      `json:"RefreshToken"`
		TokenType         string      `json:"TokenType"`
	} `json:"response"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type AuthPayload struct {
	Username string
	Password string
}

type CreateTenantPayload struct {
	Name         string `json:"name"`
	ContactEmail string `json:"contact_email"`
}

type TenantsResponse []struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	ContactEmail string `json:"contact_email"`
	CreatedAt    string `json:"created_at"`
}

type TenantByIDResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	ContactEmail string `json:"contact_email"`
	CreatedAt    string `json:"created_at"`
}
