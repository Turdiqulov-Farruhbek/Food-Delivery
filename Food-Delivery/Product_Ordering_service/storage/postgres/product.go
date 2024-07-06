package postgres

import (
	"context"
	"errors"
	product "product_ordering/genproto"
	"github.com/jackc/pgx/v5"
)

type Product struct {
	Db *pgx.Conn
}

func NewProduct(db *pgx.Conn) *Product {
	return &Product{Db: db}
}

func (p *Product) Create(ctx context.Context, req *product.CreateProductRequest) (*product.ProductResponse, error) {
	query := `INSERT INTO products (name, description, price, image_url) VALUES ($1, $2, $3, $4) RETURNING product_id`
	var productID string
	err := p.Db.QueryRow(ctx, query, req.Name, req.Description, req.Price, req.ImageUrl).Scan(&productID)
	if err != nil {
		return nil, err
	}
	return &product.ProductResponse{
		Success: true,
		Message: "Product created successfully",
		Product: &product.Product{
			ProductId:   productID,
			Name:        req.Name,
			Description: req.Description,
			Price:       req.Price,
			ImageUrl:    req.ImageUrl,
		},
	}, nil
}

func (p *Product) Get(ctx context.Context, req *product.ProductRequest) (*product.ProductResponse, error) {
	query := `SELECT product_id, name, description, price, image_url FROM products WHERE product_id=$1 AND deleted_at=0`
	var prod product.Product
	err := p.Db.QueryRow(ctx, query, req.ProductId).Scan(&prod.ProductId, &prod.Name, &prod.Description, &prod.Price, &prod.ImageUrl)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product.ProductResponse{
		Success: true,
		Message: "Product retrieved successfully",
		Product: &prod,
	}, nil
}

func (p *Product) Update(ctx context.Context, req *product.UpdateProductRequest) (*product.ProductResponse, error) {
	query := `UPDATE products SET name=$1, description=$2, price=$3, image_url=$4 WHERE product_id=$5 AND deleted_at=0`
	_, err := p.Db.Exec(ctx, query, req.Name, req.Description, req.Price, req.ImageUrl, req.ProductId)
	if err != nil {
		return nil, err
	}
	return &product.ProductResponse{
		Success: true,
		Message: "Product updated successfully",
		Product: &product.Product{
			ProductId:   req.ProductId,
			Name:        req.Name,
			Description: req.Description,
			Price:       req.Price,
			ImageUrl:    req.ImageUrl,
		},
	}, nil
}

func (p *Product) Delete(ctx context.Context, req *product.ProductRequest) (*product.ProductResponse, error) {
	query := `UPDATE products SET deleted_at=EXTRACT(EPOCH FROM NOW())::BIGINT  WHERE deleted_at=0`
	result, err := p.Db.Exec(ctx, query, req.ProductId)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected() == 0 {
		return nil, errors.New("product not found or already deleted")
	}
	return &product.ProductResponse{
		Success: true,
		Message: "Product deleted successfully",
	}, nil
}

func (p *Product) List(ctx context.Context, req *product.Empty) (*product.ProductListResponse, error) {
	query := `SELECT product_id, name, description, price, image_url FROM products WHERE deleted_at=0`
	rows, err := p.Db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*product.Product
	for rows.Next() {
		var prod product.Product
		err := rows.Scan(&prod.ProductId, &prod.Name, &prod.Description, &prod.Price, &prod.ImageUrl)
		if err != nil {
			return nil, err
		}
		products = append(products, &prod)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return &product.ProductListResponse{
		Products: products,
	}, nil
}
