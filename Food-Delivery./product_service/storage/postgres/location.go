package postgres

import (
	"database/sql"
	"fmt"
	"strings"

	pb "delivery/product_service/genproto"

	"github.com/google/uuid"
)

type LocationRepo struct {
	db *sql.DB
}

func NewLocationRepo(db *sql.DB) *LocationRepo {
	return &LocationRepo{db: db}
}
func (r *LocationRepo) CreateLocation(req *pb.LocationCreateReq) (*pb.Void, error) {
	query := `INSERT INTO locations (id,
									courier_id, 
									latitude, 
									longitude) 
							VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query,
		uuid.NewString(),
		req.CourierId,
		req.Body.Latitude,
		req.Body.Longitude)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (r *LocationRepo) GetLocation(id *pb.ById) (*pb.Location, error) {
	query := `SELECT id, 
                    courier_id, 
                    latitude, 
                    longitude,
                    updated_at
    FROM locations
    WHERE courier_id = $1 AND deleted_at = 0`
	var location pb.Location
	row := r.db.QueryRow(query, id.Id)
	err := row.Scan(
		&location.Id,
		&location.CourierId,
		&location.Latitude,
		&location.Longitude,
		&location.Time,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("location not found")
	} else if err != nil {
		return nil, err
	}
	return &location, nil
}

func (r *LocationRepo) UpdateLocation(req *pb.LocationCreateReq) (*pb.Void, error) {
	query := "UPDATE locations SET "
	var cons []string
	var args []interface{}

	// Dynamically build the query
	if req.CourierId != "" && req.CourierId != "string" {
		cons = append(cons, fmt.Sprintf("courier_id = $%d", len(args)))
		args = append(args, req.CourierId)
	}
	if req.Body.Latitude != 0 {
		cons = append(cons, fmt.Sprintf("latitude = $%d", len(args)))
		args = append(args, req.Body.Latitude)
	}
	if req.Body.Longitude != 0 {
		cons = append(cons, fmt.Sprintf("longitude = $%d", len(args)))
		args = append(args, req.Body.Longitude)
	}

	if len(cons) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}
	query += strings.Join(cons, ", ")
	query += fmt.Sprintf(" WHERE courier_id = $%d AND deleted_at = 0", len(args)+1)
	args = append(args, req.CourierId)

	// Execute the query
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
