package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	pb "delivery/notification_service/genproto"
	"delivery/notification_service/helper"

	"github.com/google/uuid"
)

type NotificationRepo struct {
	db    *sql.DB
	UserC pb.AuthServiceClient
}

func NewNotificationRepo(db *sql.DB, UserClient pb.AuthServiceClient) *NotificationRepo {
	return &NotificationRepo{db: db, UserC: UserClient}
}
func (r *NotificationRepo) CreateNotification(req *pb.NotificationCreate) (*pb.Void, error) {
	var sender_id string
	if req.SenderId == "" {
		fil := pb.Filter{
			Limit:  1,
			Offset: 0,
		}
		filter := pb.UserFilter{
			Role:   "admin",
			Filter: &fil,
		}
		res, err := r.UserC.GetAllUsers(context.Background(), &filter)
		if err != nil {
			return nil, err
		}
		sender_id = res.Users[0].Id
	} else {
		sender_id = req.SenderId
	}

	tr, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	//geting the email
	user, err := r.UserC.GetUserProfile(context.Background(), &pb.ById{Id: req.RecieverId})
	if err != nil {
		tr.Rollback()
		return nil, err
	}

	//sending email
	from := "shamsioqilov@gmail.com"
	password := "rdpo ehng abtm fuzy"
	err = helper.SendVerificationCode(helper.Params{
		From:     from,
		Password: password,
		To:       user.Email,
		Message:  fmt.Sprintf("Hi %s", user.FullName),
		Code:     req.Message,
	})

	if err != nil {
		tr.Rollback()
		return nil, errors.New("failed to send notification email" + err.Error())
	}
	query := `insert into notifications(id,
										reciever_id,
										sender_id,
										message,
										status)
						values($1,$2,$3,$4,$5)`
	_, err = tr.Exec(query,
		uuid.NewString(),
		req.RecieverId,
		sender_id,
		req.Message,
		"pending")
	if err != nil {
		tr.Rollback()
		return nil, err
	}

	err = tr.Commit()
	if err != nil {
		tr.Rollback()
		return nil, err
	}
	return &pb.Void{}, nil
}
func (r *NotificationRepo) NotifyAll(req *pb.NotificationMessage) (*pb.Void, error) {
	if req.SenderId == "" {
		fil := pb.Filter{
			Limit:  1,
			Offset: 0,
		}
		filter := pb.UserFilter{
			Role:   "admin",
			Filter: &fil,
		}
		res, err := r.UserC.GetAllUsers(context.Background(), &filter)
		if err != nil {
			return nil, err
		}
		req.SenderId = res.Users[0].Id
	}

	fil := pb.Filter{}
	filter := pb.UserFilter{
		Role:   req.TargetGroup,
		Filter: &fil,
	}
	users, err := r.UserC.GetAllUsers(context.Background(), &filter)
	if err != nil {
		return nil, err
	}

	for _, user := range users.Users {
		sender_id := req.SenderId

		tr, err := r.db.Begin()
		if err != nil {
			return nil, err
		}

		//sending email
		from := "shamsioqilov@gmail.com"
		password := "rdpo ehng abtm fuzy"
		err = helper.SendVerificationCode(helper.Params{
			From:     from,
			Password: password,
			To:       user.Email,
			Message:  fmt.Sprintf("Hi %s", user.FullName),
			Code:     req.Message,
		})

		if err != nil {
			tr.Rollback()
			return nil, errors.New("failed to send notification email" + err.Error())
		}
		query := `insert into notifications(id,
											reciever_id,
											sender_id,
											message)
							values($1,$2,$3,$4)`
		_, err = tr.Exec(query,
			uuid.NewString(),
			user.Id,
			sender_id,
			req.Message)
		if err != nil {
			tr.Rollback()
			return nil, err
		}

		err = tr.Commit()
		if err != nil {
			tr.Rollback()
			return nil, err
		}
	}
	// if err != nil {
	// 	return nil, err
	// }
	return &pb.Void{}, nil
}
func (r *NotificationRepo) DeleteNotificcation(id *pb.ById) (*pb.Void, error) {
	query := `update notifications set deleted_at = EXTRACT(EPOCH FROM now()) 
				where id = $1 and deleted_at = 0`
	_, err := r.db.Exec(query, id.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (r *NotificationRepo) UpdateNotificcation(req *pb.NotificationUpdate) (*pb.Void, error) {
	query := "UPDATE notifications SET "
	var cons []string
	var args []interface{}

	// Dynamically build the query
	if req.Body.Message != "" && req.Body.Message != "string" {
		cons = append(cons, fmt.Sprintf("message=$%d", len(args)+1))
		args = append(args, req.Body.Message)
	}
	if req.Body.Status != "" && req.Body.Status != "string" {
		cons = append(cons, fmt.Sprintf("status=$%d", len(args)+1))
		args = append(args, req.Body.Status)
	}

	// Ensure there's at least one field to update
	if len(cons) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query += strings.Join(cons, ", ")
	query += fmt.Sprintf(" WHERE deleted_at = 0 and id=$%d", len(args)+1)
	args = append(args, req.NotificationId)

	// Execute the query
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}
func (r *NotificationRepo) GetNotifications(req *pb.NotifFilter) (*pb.NotificationList, error) {

	query := `SELECT id, 
					reciever_id, 
					message, 
					status, 
					created_at,
					sender_id
		FROM notifications 
		WHERE deleted_at = 0`
	var cons []string
	var args []interface{}

	// Dynamically build the query
	if req.RecieverId != "" && req.RecieverId != "string" {
		cons = append(cons, fmt.Sprintf("reciever_id = $%d", len(args)+1))
		args = append(args, "%"+req.RecieverId+"%")
	}

	if req.Status != "" && req.Status != "string" {
		cons = append(cons, fmt.Sprintf("status = $%d", len(args)+1))
		args = append(args, "%"+req.Status+"%")
	}
	if req.SenderId != "" && req.SenderId != "string" {
		cons = append(cons, fmt.Sprintf("sender_id = $%d", len(args)+1))
		args = append(args, "%"+req.SenderId+"%")
	}

	// Append conditions to query if any exist
	if len(cons) > 0 {
		query += " AND " + strings.Join(cons, " AND ")
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)+1, len(args)+2)
	args = append(args, req.Filter.Limit, req.Filter.Offset)

	// Execute the query
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	// Prepare the response
	var notificationList pb.NotificationList
	for rows.Next() {
		var notif pb.NotificationGet
		if err := rows.Scan(&notif.Id,
			&notif.UserId,
			&notif.Message,
			&notif.Status,
			&notif.CreatedAt,
			&notif.SenderId); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		notificationList.Notifications = append(notificationList.Notifications, &notif)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %w", err)
	}
	for _, n := range notificationList.Notifications {
		query := `update notifications set status = read 
					where deleted_at = 0 and id = $1`
		_, err := r.db.Exec(query, n.Id)
		if err != nil {
			return nil, err
		}
	}

	return &notificationList, nil
}
func (r *NotificationRepo) GetNotification(id *pb.ById) (*pb.NotificationGet, error) {
	query := `select id,
					reciever_id,
					message,
					status,
					created_at,
					sender_id
			from notifications where deleted_at = 0 and id = $1`
	row := r.db.QueryRow(query, id.Id)

	var notif pb.NotificationGet
	err := row.Scan(&notif.Id,
		&notif.UserId,
		&notif.Message,
		&notif.Status,
		&notif.CreatedAt,
		&notif.SenderId)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return &notif, nil
}
