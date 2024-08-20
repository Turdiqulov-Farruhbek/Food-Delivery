package postgres

import (
	"database/sql"
	"fmt"

	// "log"
	"strconv"
	"strings"

	pb "delivery/product_service/genproto"

	"github.com/google/uuid"
)

type KitchenRepo struct {
	db *sql.DB
}

func NewKitchenRepo(db *sql.DB) *KitchenRepo {
	return &KitchenRepo{db: db}
}
func (r *KitchenRepo) CreateKitchen(req *pb.KitchenCreateReq) (*pb.Void, error) {
	query := `INSERT INTO kitchens (id, 
									manager_id,
									name, 
									phone_number,
									description,
									address,
									latitude,
									longitude)
						 VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := r.db.Exec(query,
		uuid.NewString(),
		req.ManagerId,
		req.Body.Name,
		req.Body.PhoneNumber,
		req.Body.Description,
		req.Body.Address,
		req.Body.Latitude,
		req.Body.Longitude)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (r *KitchenRepo) UpdateKitchen(req *pb.KitchenCreateReq) (*pb.Void, error) {
	query := "UPDATE kitchens SET "
	var cons []string
	var args []interface{}

	// Dynamically build the query
	if req.Body.Name != "" && req.Body.Name != "string" {
		cons = append(cons, fmt.Sprintf("name=$%d", len(args)+1))
		args = append(args, req.Body.Name)
	}
	if req.Body.PhoneNumber != "" && req.Body.PhoneNumber != "string" {
		cons = append(cons, fmt.Sprintf("phone_number=$%d", len(args)+1))
		args = append(args, req.Body.PhoneNumber)
	}
	if req.Body.Description != "" && req.Body.Description != "string" {
		cons = append(cons, fmt.Sprintf("description ILIKE $%d", len(args)+1))
		args = append(args, req.Body.Description)
	}
	if req.Body.Address != "" && req.Body.Address != "string" {
		cons = append(cons, fmt.Sprintf("address=$%d", len(args)+1))
		args = append(args, req.Body.Address)
	}
	if req.Body.Latitude != 0 {
		cons = append(cons, fmt.Sprintf("latitude=$%d", len(args)+1))
		args = append(args, req.Body.Latitude)
	}
	if req.Body.Longitude != 0 {
		cons = append(cons, fmt.Sprintf("longitude=$%d", len(args)+1))
		args = append(args, req.Body.Longitude)
	}

	// Ensure there's at least one field to update
	if len(cons) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query += strings.Join(cons, ", ")
	query += fmt.Sprintf(" WHERE deleted_at = 0 and id=$%d", len(args)+1)
	args = append(args, req.KitchenId)

	// Execute the query
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}
func (r *KitchenRepo) DeleteKitchen(id *pb.ById) (*pb.Void, error) {
	query := `UPDATE kitchens SET deleted_at = EXTRACT(EPOCH FROM now())
                WHERE id = $1 and deleted_at = 0`
	_, err := r.db.Exec(query, id.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (r *KitchenRepo) ListKitchens(req *pb.KitchenFilter) (*pb.KitchenList, error) {
	query := `SELECT id, 
					manager_id, 
					name, 
					phone_number,
					description,
					address,
					latitude,
					longitude,
					created_at
				FROM kitchens 
				WHERE deleted_at = 0`

	// Adding filters if present
	var params []interface{}
	var conditions []string

	if req.ManagerId != "" {
		conditions = append(conditions, "manager_id = $"+strconv.Itoa(len(params)+1))
		params = append(params, req.ManagerId)
	}
	if req.Name != "" {
		conditions = append(conditions, "name ILIKE '%' || $"+strconv.Itoa(len(params)+1)+" || '%'")
		params = append(params, req.Name)
	}
	if req.PhoneNumber != "" {
		conditions = append(conditions, "phone_number ILIKE '%' || $"+strconv.Itoa(len(params)+1)+" || '%'")
		params = append(params, req.PhoneNumber)
	}
	if req.Description != "" {
		conditions = append(conditions, "description ILIKE '%' || $"+strconv.Itoa(len(params)+1)+" || '%'")
		params = append(params, req.Description)
	}
	if req.Address != "" {
		conditions = append(conditions, "address ILIKE '%' || $"+strconv.Itoa(len(params)+1)+" || '%'")
		params = append(params, req.Address)
	}
	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY created_at DESC"
	if req.Filter.Limit > 0 {
		query += " LIMIT $" + strconv.Itoa(len(params)+1)
		params = append(params, req.Filter.Limit)
	}
	if req.Filter.Offset > 0 {
		query += " OFFSET $" + strconv.Itoa(len(params)+1)
		params = append(params, req.Filter.Offset)
	}

	rows, err := r.db.Query(query, params...)
	if err != nil {
		return nil, fmt.Errorf("error listing kitchens: %v", err)
	}
	defer rows.Close()

	var kitchens []*pb.KitchenGet
	for rows.Next() {
		var kitchen pb.KitchenGet
		err := rows.Scan(&kitchen.Id,
			&kitchen.ManagerId,
			&kitchen.Name,
			&kitchen.PhoneNumber,
			&kitchen.Description,
			&kitchen.Address,
			&kitchen.Latitude,
			&kitchen.Longitude,
			&kitchen.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning kitchen: %v", err)
		}
		kitchens = append(kitchens, &kitchen)
	}

	return &pb.KitchenList{
		Kitchens:   kitchens,
		TotalCount: int32(len(kitchens)),
		Limit:      req.Filter.Limit,
		Offset:     req.Filter.Offset,
	}, nil
}

func (r *KitchenRepo) GetKitchen(id *pb.ById) (*pb.KitchenGet, error) {
	query := `SELECT id, 
                    manager_id, 
                    name, 
                    phone_number, 
                    description, 
                    address, 
                    latitude, 
                    longitude,
					created_at
    FROM kitchens
    WHERE id = $1 AND deleted_at = 0`
	var kitchen pb.KitchenGet
	row := r.db.QueryRow(query, id.Id)
	err := row.Scan(
		&kitchen.Id,
		&kitchen.ManagerId,
		&kitchen.Name,
		&kitchen.PhoneNumber,
		&kitchen.Description,
		&kitchen.Address,
		&kitchen.Latitude,
		&kitchen.Longitude,
		&kitchen.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("kitchen not found")
	} else if err != nil {
		return nil, err
	}
	return &kitchen, nil
}
