package operator

import (
	"location-service/internal/client"
	"location-service/internal/model"
	"location-service/internal/streamer"
)

func InsertLocation(loc model.LocationRequest) error {
	query := `INSERT INTO locations (tenant_id, latitude, longitude, timestamp)
         VALUES ($1, $2, $3, CURRENT_TIMESTAMP)`

	_, err := client.DB.Exec(query, loc.TenantID, loc.Latitude, loc.Longitude)
	if err != nil {
		return err
	}

	streamer.SendLocation(streamer.LocationPayload{
		TenantID:  loc.TenantID,
		Latitude:  loc.Latitude,
		Longitude: loc.Longitude,
	})

	return nil
}
