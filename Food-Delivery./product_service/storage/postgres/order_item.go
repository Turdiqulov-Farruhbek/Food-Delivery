package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	pb "delivery/product_service/genproto"

	"github.com/google/uuid"
)

type OrderItemRepo struct {
	db *sql.DB
}

func NewOrderItemRepo(db *sql.DB) *OrderItemRepo {
	return &OrderItemRepo{db: db}
}

func (r *OrderItemRepo) AddItem(req *pb.ItemCreateReq) (*pb.Void, error) {
	var orderID string

	// Fetch the order ID for the user's pending order
	query := `SELECT id FROM orders WHERE user_id = $1 AND status = 'pending'`
	err := r.db.QueryRow(query, req.UserId).Scan(&orderID)

	if err == sql.ErrNoRows {
		// If no pending order exists, create a new order
		orderID = uuid.NewString()
		query = `INSERT INTO orders(id, 
                                    user_id, 
                                    status,
                                    address,
                                    delivery_time,
                                    is_paid,
                                    courier_id,
                                    total) 
                 VALUES($1, $2, $3, $4, $5, $6, $7, $8)`
		_, err := r.db.Exec(query,
			orderID,
			req.UserId,
			"pending",
			"address",                              // Assuming address comes from the request
			"2000-01-01 14:30:00",                  // You may want to replace this with actual delivery time
			false,                                  // Assuming is_paid is false for a new order
			"00000000-0000-0000-0000-000000000000", // Assuming no courier assigned yet
			0.0)                                    // Initial total is 0.0 for a new order

		if err != nil {
			return nil, errors.New("error creating new order: " + err.Error())
		}
	} else if err != nil {
		return nil, errors.New("error finding user's pending order: " + err.Error())
	}

	// Calculate the total price for the item
	var unitPrice float64
	query = `SELECT price FROM products WHERE id = $1`
	err = r.db.QueryRow(query, req.Body.ProductId).Scan(&unitPrice)
	if err != nil {
		return nil, errors.New("error fetching product price: " + err.Error())
	}
	totalPrice := unitPrice * float64(req.Body.Quantity)

	// Insert the item into the order
	query = `INSERT INTO order_items(
                                id,
                                order_id,
                                product_id,
                                quantity,
                                price)
             VALUES($1, $2, $3, $4, $5)`
	_, err = r.db.Exec(query,
		uuid.NewString(),
		orderID,
		req.Body.ProductId,
		req.Body.Quantity,
		totalPrice)
	if err != nil {
		return nil, errors.New("error adding item to order: " + err.Error())
	}

	// Optionally, update the order's total if needed
	query = `UPDATE orders SET total = total + $1 WHERE id = $2`
	_, err = r.db.Exec(query, totalPrice, orderID)
	if err != nil {
		return nil, errors.New("error updating order total: " + err.Error())
	}

	return &pb.Void{}, nil
}

func (r *OrderItemRepo) GetItem(id *pb.ById) (*pb.ItemGet, error) {
	query := `SELECT id, 
					order_id, 
					product_id, 
					quantity,
					price,
					created_at
					FROM order_items 
					WHERE id = $1 and deleted_at =0`
	row := r.db.QueryRow(query, id.Id)

	var item pb.ItemGet
	var prod_id string
	err := row.Scan(&item.Id,
		&item.OrderId,
		&prod_id,
		&item.Quantity,
		&item.Price,
		&item.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("item not found")
	} else if err != nil {
		return nil, errors.New("error getting item:" + err.Error())
	}
	// get product
	query1 := `select id,
					name,
					description,
                    price,
                    image,
					created_at
				from products 
				where id = $1 and deleted_at = 0`
	var product pb.ProductGet
	err = r.db.QueryRow(query1, prod_id).Scan(&product.Id,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Image,
		&product.CreatedAt)
	if err != nil {
		return nil, errors.New("error getting product:" + err.Error())
	}
	item.Body = &product

	return &item, nil
}

func (r *OrderItemRepo) UpdateItem(req *pb.ItemCreateReq) (*pb.Void, error) {
	log.Println(req)
	// fetch item credentials to compare
	query1 := `select product_id, quantity, price, order_id 
						from order_items 
						where id = $1 and deleted_at = 0`
	var old_prod_id string
	var old_quantity int
	var old_price float64
	var order_id string

	err := r.db.QueryRow(query1, req.ItemId).Scan(&old_prod_id, &old_quantity, &old_price, &order_id)
	if err == sql.ErrNoRows {
		return nil, errors.New("item not found")
	} else if err != nil {
		return nil, errors.New("error getting item credentials:" + err.Error())
	}

	//update

	query2 := "UPDATE order_items SET "
	var cons []string
	var args []interface{}

	if req.Body.ProductId != "" && req.Body.ProductId != "string" {
		cons = append(cons, fmt.Sprintf("product_id=$%d", len(args)+1))
		args = append(args, req.Body.ProductId)
	}
	if req.Body.Quantity != 0 {
		cons = append(cons, fmt.Sprintf("quantity=$%d", len(args)+1))
		args = append(args, req.Body.Quantity)
	}
	// conditions to calculate new sum
	var new_price float64
	if req.Body.ProductId != old_prod_id && req.Body.Quantity != int32(old_quantity) && req.Body.ProductId != "" {

		query1 := `select price from products where id = $1`
		// var new_price float64
		err = r.db.QueryRow(query1, req.Body.ProductId).Scan(&new_price)
		if err != nil {
			return nil, errors.New("error fetching product price:" + err.Error())
		}
		new_price = new_price * float64(req.Body.Quantity)

	} else if req.Body.ProductId != old_prod_id && req.Body.Quantity == int32(old_quantity) {
		query1 := `select price from products where id = $1`
		// var new_price float64
		err = r.db.QueryRow(query1, req.Body.ProductId).Scan(&new_price)
		if err != nil {
			return nil, errors.New("error fetching product price:" + err.Error())
		}
		new_price = new_price * float64(old_quantity)

	} else if (req.Body.ProductId == old_prod_id || req.Body.ProductId == "") && req.Body.Quantity != int32(old_quantity) {
		query1 := `select price from products where id = $1`
		// var new_price float64
		err = r.db.QueryRow(query1, old_prod_id).Scan(&new_price)
		if err != nil {
			return nil, errors.New("error fetching product price:" + err.Error())
		}
		new_price = new_price * float64(req.Body.Quantity)
	}
	if new_price != 0 {
		cons = append(cons, fmt.Sprintf("price=$%d", len(args)+1))
		args = append(args, new_price)
	}

	// Ensure there's at least one field to update
	if len(cons) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query2 += strings.Join(cons, ", ")
	query2 += fmt.Sprintf(" WHERE deleted_at = 0 and id=$%d", len(args)+1)
	args = append(args, req.ItemId)

	// Execute the query
	_, err = r.db.Exec(query2, args...)
	if err != nil {
		return nil, err
	}
	// update total of orders
	if new_price > 0 {
		diff := new_price - old_price
		query3 := `UPDATE orders SET total = total + $1 WHERE id = $2`
		_, err = r.db.Exec(query3, diff, order_id)
		if err != nil {
			return nil, errors.New("error updating order total: " + err.Error())
		}
	}

	return &pb.Void{}, nil
}
func (r *OrderItemRepo) DeleteItem(id *pb.ById) (*pb.Void, error) {
	query := `UPDATE order_items SET deleted_at = EXTRACT(EPOCH FROM now()) 
						WHERE id = $1`
	result, err := r.db.Exec(query, id.Id)
	if err != nil {
		return nil, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, errors.New("item not found")
	}
	return &pb.Void{}, nil
}
func (r *OrderItemRepo) ListItems(filter *pb.ItemFilter) (*pb.ItemList, error) {
	query := `SELECT oi.id, 
					oi.order_id, 
					oi.product_id, 
					oi.quantity,
					oi.price,
					oi.created_at
				FROM order_items oi
				JOIN orders o ON oi.order_id = o.id
				WHERE oi.deleted_at = 0`

	// Adding filters if present
	var params []interface{}
	var conditions []string

	if filter.UserId != "" && filter.UserId != "undefined" {
		conditions = append(conditions, "o.user_id = $"+strconv.Itoa(len(params)+1))
		params = append(params, filter.UserId)
	}
	if filter.OrderId != "" {
		conditions = append(conditions, "oi.order_id = $"+strconv.Itoa(len(params)+1))
		params = append(params, filter.OrderId)
	}
	if filter.ProductId != "" {
		conditions = append(conditions, "oi.product_id = $"+strconv.Itoa(len(params)+1))
		params = append(params, filter.ProductId)
	}
	if filter.MinPrice > 0 {
		conditions = append(conditions, "oi.price >= $"+strconv.Itoa(len(params)+1))
		params = append(params, filter.MinPrice)
	}
	if filter.MaxPrice > 0 {
		conditions = append(conditions, "oi.price <= $"+strconv.Itoa(len(params)+1))
		params = append(params, filter.MaxPrice)
	}
	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY oi.created_at DESC"
	if filter.Limit > 0 {
		query += " LIMIT $" + strconv.Itoa(len(params)+1)
		params = append(params, filter.Limit)
	}
	if filter.Offset > 0 {
		query += " OFFSET $" + strconv.Itoa(len(params)+1)
		params = append(params, filter.Offset)
	}

	rows, err := r.db.Query(query, params...)
	if err != nil {
		return nil, fmt.Errorf("error listing items: %v", err)
	}
	defer rows.Close()

	var items []*pb.ItemGet
	for rows.Next() {
		var item pb.ItemGet
		var prod_id string
		err := rows.Scan(&item.Id,
			&item.OrderId,
			&prod_id,
			&item.Quantity,
			&item.Price,
			&item.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning item: %v", err)
		}

		// Get product details
		query1 := `SELECT id,
						name,
						description,
						price,
						image,
						created_at
					FROM products 
					WHERE id = $1 AND deleted_at = 0`
		var product pb.ProductGet
		err = r.db.QueryRow(query1, prod_id).Scan(&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Image,
			&product.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error getting product: %v", err)
		}
		item.Body = &product

		items = append(items, &item)
	}

	// Get total count for pagination
	count := len(items)

	return &pb.ItemList{
		Items:  items,
		Count:  int32(count),
		Limit:  filter.Limit,
		Offset: filter.Offset,
	}, nil
}
