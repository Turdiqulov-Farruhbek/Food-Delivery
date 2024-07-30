package postgres

import (
	"context"
	"errors"
	gen "product_ordering/genproto/product"

	"github.com/jackc/pgx/v5"
)

type OrderRecommendation struct {
	Db *pgx.Conn
}

func NewOrderRecomend (db *pgx.Conn) *OrderRecommendation {
	return &OrderRecommendation{Db: db}
}

func (o *OrderRecommendation) Create(ctx context.Context, req *gen.CreateOrderRecommendationRequest) (*gen.OrderRecommendationResponse, error) {
	query := `INSERT INTO orderrecommendations (order_id, courier_id, recommended_at, status) VALUES ($1, $2, $3, $4) RETURNING recommendation_id`
	var recommendationID string
	err := o.Db.QueryRow(ctx, query, req.OrderId, req.CourierId, req.RecommendedAt, req.Status).Scan(&recommendationID)
	if err != nil {
		return nil, err
	}
	return &gen.OrderRecommendationResponse{
		Success: true,
		Message: "Order recommendation created successfully",
		OrderRecommendation: &gen.OrderRecommendation{
			RecommendationId: recommendationID,
			OrderId:          req.OrderId,
			CourierId:        req.CourierId,
			RecommendedAt:    req.RecommendedAt,
			Status:           req.Status,
		},
	}, nil
}

func (o *OrderRecommendation) Get(ctx context.Context, req *gen.OrderRecommendationRequest) (*gen.OrderRecommendationResponse, error) {
	query := `SELECT recommendation_id, order_id, courier_id, recommended_at, status FROM orderrecommendations WHERE recommendation_id=$1`
	var recommendation gen.OrderRecommendation
	err := o.Db.QueryRow(ctx, query, req.RecommendationId).Scan(&recommendation.RecommendationId, &recommendation.OrderId, &recommendation.CourierId, &recommendation.RecommendedAt, &recommendation.Status)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("order recommendation not found")
		}
		return nil, err
	}
	return &gen.OrderRecommendationResponse{
		Success:            true,
		Message:            "Order recommendation retrieved successfully",
		OrderRecommendation: &recommendation,
	}, nil
}

func (o *OrderRecommendation) Update(ctx context.Context, req *gen.UpdateOrderRecommendationRequest) (*gen.OrderRecommendationResponse, error) {
	// Mavjud qiymatlarni olish
	getQuery := `SELECT order_id, courier_id, recommended_at, status FROM orderrecommendations WHERE recommendation_id=$1`
	var current gen.OrderRecommendation
	err := o.Db.QueryRow(ctx, getQuery, req.RecommendationId).Scan(&current.OrderId, &current.CourierId, &current.RecommendedAt, &current.Status)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("order recommendation not found")
		}
		return nil, err
	}

	// Agar yangi qiymatlar berilmagan bo'lsa, eski qiymatlarni saqlab qolish
	orderId := req.OrderId
	if orderId == "" {
		orderId = current.OrderId
	}

	courierId := req.CourierId
	if courierId == "" {
		courierId = current.CourierId
	}

	recommendedAt := req.RecommendedAt
	if recommendedAt == "" {
		recommendedAt = current.RecommendedAt
	}

	status := req.Status
	if status == gen.RecommendationStatus_PENDING {
		status = current.Status
	}

	// Yangilash
	updateQuery := `UPDATE orderrecommendations SET order_id=$1, courier_id=$2, recommended_at=$3, status=$4 WHERE recommendation_id=$5`
	_, err = o.Db.Exec(ctx, updateQuery, orderId, courierId, recommendedAt, status, req.RecommendationId)
	if err != nil {
		return nil, err
	}

	return &gen.OrderRecommendationResponse{
		Success: true,
		Message: "Order recommendation updated successfully",
		OrderRecommendation: &gen.OrderRecommendation{
			RecommendationId: req.RecommendationId,
			OrderId:          orderId,
			CourierId:        courierId,
			RecommendedAt:    recommendedAt,
			Status:           status,
		},
	}, nil
}

func (o *OrderRecommendation) Delete(ctx context.Context, req *gen.OrderRecommendationRequest) (*gen.OrderRecommendationResponse, error) {
	query := `UPDATE orderrecommendations SET deleted_at=EXTRACT(EPOCH FROM NOW())::BIGINT  WHERE recommendation_id=$1 AND deleted_at=0`
	result, err := o.Db.Exec(ctx, query, req.RecommendationId)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected() == 0 {
		return nil, errors.New("order recommendation not found or already deleted")
	}
	return &gen.OrderRecommendationResponse{
		Success: true,
		Message: "Order recommendation deleted successfully",
	}, nil
}

func (o *OrderRecommendation) List(ctx context.Context, req *gen.Empty) (*gen.OrderRecommendationListResponse, error) {
	query := `SELECT recommendation_id, order_id, courier_id, recommended_at, status FROM orderrecommendations WHERE deleted_at=0`
	rows, err := o.Db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recommendations []*gen.OrderRecommendation
	for rows.Next() {
		var recommendation gen.OrderRecommendation
		err := rows.Scan(&recommendation.RecommendationId, &recommendation.OrderId, &recommendation.CourierId, &recommendation.RecommendedAt, &recommendation.Status)
		if err != nil {
			return nil, err
		}
		recommendations = append(recommendations, &recommendation)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return &gen.OrderRecommendationListResponse{
		OrderRecommendations: recommendations,
	}, nil
}
