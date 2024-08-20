package postgres

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	pb "delivery/product_service/genproto"
	"delivery/product_service/kafka"

	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
)

type OrderRepo struct {
	db          *sql.DB
	kf_producer kafka.KafkaProducer
}

func NewOrderRepo(db *sql.DB, kf_producer kafka.KafkaProducer) *OrderRepo {
	return &OrderRepo{db: db, kf_producer: kf_producer}
}
func (r *OrderRepo) CreateOrder(req *pb.OrderCreateReq) (*pb.Void, error) {

	query := `insert into orders(id, 
								user_id, 
								status,
								address.
								delivery_time,
								is_paid,
								courier_id,
								total) 
						values($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.Exec(query,
		uuid.NewString(),
		req.UserId,
		"pending",
		req.Body.DeliveryAddress,
		req.Body.DeliveryTime,
		"no",
		"00000000-0000-0000-0000-000000000000",
		0)

	if err != nil {
		return nil, err
	}
	//write to kafka
	message := fmt.Sprintf("user with id %s just created new order", req.UserId)

	req2 := &pb.NotificationMessage{
		Message:     message,
		TargetGroup: "admin",
	}
	input, err := protojson.Marshal(req2) ///////////////////////////////////////////////////
	if err != nil {
		return nil, err
	}
	err = r.kf_producer.ProduceMessages("notify-all", input)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}
func (r *OrderRepo) GetOrder(id *pb.ById) (*pb.OrderGet, error) {
	query := `select id,
                    user_id,
                    status,
                    address,
                    delivery_time,
                    is_paid,
                    courier_id,
                    total,
					created_at
                from orders
                where id = $1`
	row := r.db.QueryRow(query, id.Id)

	var order pb.OrderGet
	err := row.Scan(
		&order.Id,
		&order.UserId,
		&order.Status,
		&order.Address,
		&order.DeliveryTime,
		&order.IsPaid,
		&order.CourierId,
		&order.Total,
		&order.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &order, nil
}
func (r *OrderRepo) UpdateOrder(req *pb.OrderUpdate) (*pb.Void, error) {
	query := "UPDATE orders SET "
	var cons []string
	var args []interface{}

	// Dynamically build the query
	if req.Body.CourierId != "" && req.Body.CourierId != "string" {
		cons = append(cons, fmt.Sprintf("courier_id=$%d", len(args)+1))
		args = append(args, req.Body.CourierId)
	}
	if req.Body.Status != "" && req.Body.Status != "string" {
		cons = append(cons, fmt.Sprintf("status=$%d", len(args)+1))
		args = append(args, req.Body.Status)
	}
	if req.Body.Total != 0 {
		cons = append(cons, fmt.Sprintf("total=$%d", len(args)+1))
		args = append(args, req.Body.Total)
	}
	if req.Body.Address != "" && req.Body.Address != "string" {
		cons = append(cons, fmt.Sprintf("address=$%d", len(args)+1))
		args = append(args, req.Body.Address)
	}
	if req.Body.DeliveryTime != "" && req.Body.DeliveryTime != "string" {
		cons = append(cons, fmt.Sprintf("delivery_time=$%d", len(args)+1))
		args = append(args, req.Body.DeliveryTime)
	}

	// Ensure there's at least one field to update
	if len(cons) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query += strings.Join(cons, ", ")
	query += fmt.Sprintf(" WHERE deleted_at = 0 and id=$%d", len(args)+1)
	args = append(args, req.Id)

	// Execute the query
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	if req.Body.Status != "" && req.Body.Status != "string" {
		//get user id
		query := `select user_id from orders where id = $1`
		row := r.db.QueryRow(query, req.Id)
		var userId string
		err := row.Scan(&userId)
		if err != nil {
			return nil, err
		}
		//write to kafka
		message := fmt.Sprintf("status of your order with id %s just changed to %s", req.Id, req.Body.Status)

		req := &pb.NotificationCreate{
			RecieverId: userId,
			Message:    message,
		}
		input, err := protojson.Marshal(req) /////////////////////////////////////////////////////////////////////////////////////////
		if err != nil {
			return nil, err
		}
		err = r.kf_producer.ProduceMessages("notifification-create", input)
		if err != nil {
			return nil, err
		}
	}

	return &pb.Void{}, nil
}
func (r *OrderRepo) DeleteOrder(id *pb.ById) (*pb.Void, error) {
	query := `UPDATE orders SET deleted_at = EXTRACT(EPOCH FROM now())
                WHERE id = $1 and deleted_at = 0`
	_, err := r.db.Exec(query, id.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}
func (r *OrderRepo) ListOrders(filter *pb.OrderFilter) (*pb.OrderList, error) {
	query := `SELECT id, 
					user_id, 
					courier_id, 
					status,
					total,
					address,
					created_at,
					delivery_time,
					is_paid
				FROM orders 
				WHERE deleted_at = 0`

	// Adding filters if present
	var params []interface{}
	var conditions []string

	if filter.UserId != "" {
		conditions = append(conditions, "user_id = $"+strconv.Itoa(len(params)+1))
		params = append(params, filter.UserId)
	}
	if filter.Status != "" {
		conditions = append(conditions, "status = $"+strconv.Itoa(len(params)+1))
		params = append(params, filter.Status)
	}
	if filter.Address != "" {
		conditions = append(conditions, "address ILIKE '%' || $"+strconv.Itoa(len(params)+1)+" || '%'")
		params = append(params, filter.Address)
	}
	if filter.CourierId != "" {
		conditions = append(conditions, "courier_id = $"+strconv.Itoa(len(params)+1))
		params = append(params, filter.CourierId)
	}
	if filter.MinTotal > 0 {
		conditions = append(conditions, "total >= $"+strconv.Itoa(len(params)+1))
		params = append(params, filter.MinTotal)
	}
	if filter.MaxTotal > 0 {
		conditions = append(conditions, "total <= $"+strconv.Itoa(len(params)+1))
		params = append(params, filter.MaxTotal)
	}
	if filter.TimeFrom != "" {
		conditions = append(conditions, "(delivery_time >= $"+strconv.Itoa(len(params)+1)+" OR delivery_time >= $"+strconv.Itoa(len(params)+2)+")")
		params = append(params, filter.TimeFrom, filter.TimeFrom)
	}
	if filter.TimeTo != "" {
		conditions = append(conditions, "(delivery_time <= $"+strconv.Itoa(len(params)+1)+" OR delivery_time <= $"+strconv.Itoa(len(params)+2)+")")
		params = append(params, filter.TimeTo, filter.TimeTo)
	}
	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY created_at DESC"
	if filter.Filter.Limit > 0 {
		query += " LIMIT $" + strconv.Itoa(len(params)+1)
		params = append(params, filter.Filter.Limit)
	}
	if filter.Filter.Offset > 0 {
		query += " OFFSET $" + strconv.Itoa(len(params)+1)
		params = append(params, filter.Filter.Offset)
	}

	rows, err := r.db.Query(query, params...)
	if err != nil {
		return nil, fmt.Errorf("error listing orders: %v", err)
	}
	defer rows.Close()

	var orders []*pb.OrderGet
	for rows.Next() {
		var order pb.OrderGet
		err := rows.Scan(&order.Id,
			&order.UserId,
			&order.CourierId,
			&order.Status,
			&order.Total,
			&order.Address,
			&order.CreatedAt,
			&order.DeliveryTime,
			&order.IsPaid)
		if err != nil {
			return nil, fmt.Errorf("error scanning order: %v", err)
		}
		orders = append(orders, &order)
	}

	return &pb.OrderList{
		Orders: orders,
		Count:  int32(len(orders)),
		Limit:  filter.Filter.Limit,
		Offset: filter.Filter.Offset,
	}, nil
}
func (r *OrderRepo) AssignOrder(req *pb.OrderAssign) (*pb.Void, error) {
	var courier_id string
	check_query := `select courier_id from orders where id = $1`
	row := r.db.QueryRow(check_query, req.OrderId)
	err := row.Scan(&courier_id)
	if err != nil || courier_id != "00000000-0000-0000-0000-000000000000" {
		return nil, fmt.Errorf("order is already assigned to a courier")
	}

	query := `UPDATE orders SET courier_id = $1 WHERE id = $2 AND deleted_at = 0`
	_, err = r.db.Exec(query, req.CourierId, req.OrderId)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}
