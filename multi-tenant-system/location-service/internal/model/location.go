package model

type LocationRequest struct {
	TenantID  string  `json:"tenant_id"`
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}
