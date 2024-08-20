package postgres

import (
	"database/sql"
	// "errors"
	"fmt"
	// "image/draw"
	"strings"

	pb "delivery/product_service/genproto"

	"github.com/google/uuid"
)

type ProductRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) CreateProduct(req *pb.ProductCreateReq) (*pb.Void, error) {
	query := `insert into products(
						id,
						name,
                        description,
                        price,
						image,
						kitchen_id)
			values($1,$2,$3,$4,$5,$6)`
	_, err := r.db.Exec(query,
		uuid.NewString(),
		req.Body.Name,
		req.Body.Description,
		req.Body.Price,
		req.Body.Image,
		req.KitchenId)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (r *ProductRepo) GetProduct(id *pb.ById) (*pb.ProductGet, error) {
	query := `select id, 
					name, 
					description, 
					price, 
					image,
					created_at,
					kitchen_id
				from products where id = $1 and deleted_at = 0`
	row := r.db.QueryRow(query, id.Id)

	var product pb.ProductGet
	err := row.Scan(&product.Id,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Image,
		&product.KitchenId)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return &product, nil
}
func (r *ProductRepo) UpdateProduct(req *pb.ProductCreateReq) (*pb.Void, error) {
	query := "UPDATE products SET "
	var cons []string
	var args []interface{}

	// Dynamically build the query
	if req.Body.Name != "" && req.Body.Name != "string" {
		cons = append(cons, fmt.Sprintf("name=$%d", len(args)+1))
		args = append(args, req.Body.Name)
	}
	if req.Body.Description != "" && req.Body.Description != "string" {
		cons = append(cons, fmt.Sprintf("description=$%d", len(args)+1))
		args = append(args, req.Body.Description)
	}
	if req.Body.Price != 0.0 {
		cons = append(cons, fmt.Sprintf("price=$%d", len(args)+1))
		args = append(args, req.Body.Price)
	}
	if req.Body.Image != "" && req.Body.Image != "string" {
		cons = append(cons, fmt.Sprintf("image=$%d", len(args)+1))
		args = append(args, req.Body.Image)
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

	return &pb.Void{}, nil
}
func (r *ProductRepo) DeleteProduct(id *pb.ById) (*pb.Void, error) {
	query := `UPDATE products SET deleted_at = EXTRACT(EPOCH FROM now())
				WHERE id = $1 and deleted_at = 0`
	_, err := r.db.Exec(query, id.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (r *ProductRepo) ListProducts(req *pb.ProductFilter) (*pb.ProductList, error) {
	query := `SELECT id, 
                    name, 
                    description, 
                    price, 
                    image,
                    created_at ,
					kitchen_id
                FROM products 
                WHERE deleted_at = 0`
	var cons []string
	var args []interface{}

	// Dynamically build the query
	if req.Name != "" && req.Name != "string" {
		cons = append(cons, fmt.Sprintf("name ILIKE $%d", len(args)+1))
		args = append(args, "%"+req.Name+"%")
	}
	if req.MinPrice != 0.0 {
		cons = append(cons, fmt.Sprintf("price >= $%d", len(args)+1))
		args = append(args, req.MinPrice)
	}
	if req.MaxPrice != 0.0 {
		cons = append(cons, fmt.Sprintf("price <= $%d", len(args)+1))
		args = append(args, req.MaxPrice)
	}
	if req.Description != "" && req.Description != "string" {
		cons = append(cons, fmt.Sprintf("description ILIKE $%d", len(args)+1))
		args = append(args, "%"+req.Description+"%")
	}
	if req.KitchenId != "" && req.KitchenId != "string" {
		cons = append(cons, fmt.Sprintf("kitchen_id = $%d", len(args)+1))
		args = append(args, "%"+req.KitchenId+"%")
	}

	// Append conditions to query if any exist
	if len(cons) > 0 {
		query += " AND " + strings.Join(cons, " AND ")
	}

	if req.Filter.Limit != 0 {
		query += fmt.Sprintf("LIMIT $%d ", len(args)+1)
		args = append(args, req.Filter.Limit)
	}
	if req.Filter.Offset != 0 {
		query += fmt.Sprintf("OFFSET $%d ", len(args)+1)
		args = append(args, req.Filter.Offset)
	}

	// Execute the query
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	// Prepare the response
	var productList pb.ProductList
	for rows.Next() {
		var product pb.ProductGet
		if err := rows.Scan(&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Image,
			&product.CreatedAt,
			&product.KitchenId); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		productList.Products = append(productList.Products, &product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %w", err)
	}

	return &productList, nil
}
