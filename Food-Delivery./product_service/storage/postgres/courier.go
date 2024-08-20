package postgres

import (
	"database/sql"
	"fmt"

	pb "delivery/product_service/genproto"
)

type CourierRepo struct {
	db *sql.DB
}

func NewCourierRepo(db *sql.DB) *CourierRepo {
	return &CourierRepo{db: db}
}

func (r *CourierRepo) AcceptOrder(req *pb.OrderAcept) (*pb.Void, error) {
	check_query := `select courier_id from orders where id = $1 and deleted_at = 0`
	row := r.db.QueryRow(check_query, req.OrderId)
	var courier_id string
	err := row.Scan(&courier_id)
	if err != nil || courier_id != "00000000-0000-0000-0000-000000000000" {
		return nil, fmt.Errorf("error: order is already assigned to a courier")
	}
	query := `UPDATE orders SET courier_id = $1 WHERE id = $2`
	_, err = r.db.Exec(query, req.CourierId, req.OrderId)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
