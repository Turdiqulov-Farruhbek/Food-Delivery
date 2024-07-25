package postgres

import (
	"context"
	"errors"

	ntf "Notification/genproto"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type AdminAlert struct {
	Db *pgx.Conn
}

func NewAdminAlirt(db *pgx.Conn) *AdminAlert {
	return &AdminAlert{Db: db}
}

func (a *AdminAlert) CreateAdminAlert(ctx context.Context, req *ntf.CreateAdminAlertRequest) (*ntf.AdminAlertResponse, error) {
	alertID := uuid.New().String()

	query := `
		INSERT INTO adminalerts (alert_id, admin_id, type, message)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING alert_id, admin_id, type, message, created_at
	`

	row := a.Db.QueryRow(ctx, query, alertID, req.AdminId, req.Type, req.Message)
	var alert ntf.AdminAlert
	err := row.Scan(&alert.AlertId, &alert.AdminId, &alert.Type, &alert.Message, &alert.CreatedAt, &alert.IsRead)
	if err != nil {
		return nil, err
	}

	return &ntf.AdminAlertResponse{
		Success: true,
		Message: "AdminAlert create successfully",
	}, nil
}

func (a *AdminAlert) GetAdminAlert(ctx context.Context, req *ntf.AdminAlertRequest) (*ntf.AdminAlertResponse, error) {
	query := `
		SELECT alert_id, admin_id, type, message, created_at, is_read
		FROM adminalerts
		WHERE alert_id = $1  AND deleted_at = 0
	`

	row := a.Db.QueryRow(ctx, query, req.AlertId)
	var alert ntf.AdminAlert
	err := row.Scan(&alert.AlertId, &alert.AdminId, &alert.Type, &alert.Message, &alert.CreatedAt, &alert.IsRead)
	if err != nil {
		return nil, err
	}

	return &ntf.AdminAlertResponse{
		Success: true,
		Message: "AdminAlert get successfully",
	}, nil
}

func (a *AdminAlert) DeleteAdminAlert(ctx context.Context, req *ntf.AdminAlertRequest) (*ntf.AdminAlertResponse, error) {
	query := `UPDATE adminalerts SET deleted_at=EXTRACT(EPOCH FROM NOW())::BIGINT  WHERE alert_id=$1 AND deleted_at=0`
	result, err := a.Db.Exec(ctx, query, req.AlertId)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected() == 0 {
		return nil, errors.New("order adminalerts not found or already deleted")
	}
	return &ntf.AdminAlertResponse{
		Success: true,
		Message: "AdminAlert deleted successfully",
	}, nil
}

func (a *AdminAlert) ListAdminAlerts(ctx context.Context, req *ntf.Empty) (*ntf.AdminAlertListResponse, error) {
	query := `
		SELECT alert_id, admin_id, type, message, created_at, is_read
		FROM adminalerts WHERE deleted_at = 0
	`

	rows, err := a.Db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alerts []*ntf.AdminAlert
	for rows.Next() {
		var alert ntf.AdminAlert
		err := rows.Scan(&alert.AlertId, &alert.AdminId, &alert.Type, &alert.Message, &alert.CreatedAt, &alert.IsRead)
		if err != nil {
			return nil, err
		}
		alerts = append(alerts, &alert)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &ntf.AdminAlertListResponse{
		Alerts: alerts,
	}, nil
}
