package courier

import (
	"context"
	"courier_delivery/genproto"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type CourierLocation struct {
	Db *pgx.Conn
}

func NewCourierLocation(db *pgx.Conn) *CourierLocation {
	return &CourierLocation{Db: db}
}

// CreateCourierLocation yangi kuryer lokatsiyasini yaratadi
func (c *CourierLocation) CreateCourierLocation(ctx context.Context, req *genproto.CreateCourierLocationRequest) (*genproto.CourierLocationResponse, error) {
	query := `INSERT INTO courier_locations (courier_id, latitude, longitude, updated_at) VALUES ($1, $2, $3, NOW()) RETURNING id, updated_at`
	var location genproto.CourierLocationResponse
	err := c.Db.QueryRow(ctx, query, req.CourierId, req.Latitude, req.Longitude).Scan(&location.Id, &location.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create courier location: %w", err)
	}
	location.CourierId = req.CourierId
	location.Latitude = req.Latitude
	location.Longitude = req.Longitude
	return &location, nil
}

// GetCourierLocation kuryer lokatsiyasini ID orqali qaytaradi
func (c *CourierLocation) GetCourierLocation(ctx context.Context, req *genproto.GetCourierLocationRequest) (*genproto.CourierLocationResponse, error) {
	query := `SELECT id, courier_id, latitude, longitude, updated_at FROM courier_locations WHERE id=$1`
	var location genproto.CourierLocationResponse
	err := c.Db.QueryRow(ctx, query, req.Id).Scan(&location.Id, &location.CourierId, &location.Latitude, &location.Longitude, &location.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get courier location: %w", err)
	}
	return &location, nil
}

// UpdateCourierLocation kuryer lokatsiyasini yangilaydi
func (c *CourierLocation) UpdateCourierLocation(ctx context.Context, req *genproto.UpdateCourierLocationRequest) (*genproto.CourierLocationResponse, error) {
	query := `UPDATE courier_locations SET latitude=$1, longitude=$2, updated_at=NOW() WHERE id=$3 RETURNING id, courier_id, latitude, longitude, updated_at`
	var location genproto.CourierLocationResponse
	err := c.Db.QueryRow(ctx, query, req.Latitude, req.Longitude, req.Id).Scan(&location.Id, &location.CourierId, &location.Latitude, &location.Longitude, &location.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to update courier location: %w", err)
	}
	return &location, nil
}

// DeleteCourierLocation kuryer lokatsiyasini o'chiradi
func (c *CourierLocation) DeleteCourierLocation(ctx context.Context, req *genproto.DeleteCourierLocationRequest) (*genproto.DeleteCourierLocationResponse, error) {
	query := `DELETE FROM courier_locations WHERE id=$1 RETURNING id, `
	_, err := c.Db.Exec(ctx, query, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete courier location: %w", err)
	}
	return &genproto.DeleteCourierLocationResponse{
		Success: true,
        Message: "Courier location deleted successfully",
	}, nil
}

// GetAllCourierLocations barcha kuryer lokatsiyalarini qaytaradi
func (c *CourierLocation) GetAllCourierLocations(ctx context.Context, req *genproto.GetAllCourierLocationsRequest) (*genproto.GetAllCourierLocationsResponse, error) {
	query := `SELECT id, courier_id, latitude, longitude, updated_at FROM courier_locations`
	rows, err := c.Db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all courier locations: %w", err)
	}
	defer rows.Close()

	var locations []*genproto.CourierLocationResponse
	for rows.Next() {
		var location genproto.CourierLocationResponse
		err := rows.Scan(&location.Id, &location.CourierId, &location.Latitude, &location.Longitude, &location.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan courier location: %w", err)
		}
		locations = append(locations, &location)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("failed to iterate over courier locations: %w", rows.Err())
	}

	return &genproto.GetAllCourierLocationsResponse{
		Locations: locations,
	}, nil
}
