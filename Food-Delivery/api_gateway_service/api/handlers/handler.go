package handlers

import (
	c "gateway/genproto/courier"
	n "gateway/genproto/notification"
	p "gateway/genproto/product"
	u "gateway/genproto/user"

	"google.golang.org/grpc"
)

type Handler struct {
	Courier_C             c.CourierServiceClient
	CourierOrder_C        c.CourierOrderServiceClient
	Order_C               c.OrderServiceClient
	CourierLocation_C     c.CourierLocationServiceClient
	Task_C                c.TaskServiceClient
	AdminAlert_N          n.AdminAlertServiceClient
	CourierNtfn_N         n.CourierNotificationServiceClient
	UserNtfn_N            n.UserNotificationServiceClient
	Cart_P                p.CartServiceClient
	CartItemS_P           p.CartItemServiceClient
	Order_P               p.OrderProductServiceClient
	OrderRecommendation_P p.OrderRecommendationServiceClient
	Product_P             p.ProductServiceClient
	User                  u.AuthServiceClient
}

func NewHandler(conn *grpc.ClientConn) *Handler {
	return &Handler{
        Courier_C:             c.NewCourierServiceClient(conn),
        CourierOrder_C:        c.NewCourierOrderServiceClient(conn),
        Order_C:               c.NewOrderServiceClient(conn),
        CourierLocation_C:     c.NewCourierLocationServiceClient(conn),
        Task_C:                c.NewTaskServiceClient(conn),
        AdminAlert_N:          n.NewAdminAlertServiceClient(conn),
        CourierNtfn_N:         n.NewCourierNotificationServiceClient(conn),
        UserNtfn_N:            n.NewUserNotificationServiceClient(conn),
        Cart_P:                p.NewCartServiceClient(conn),
        CartItemS_P:           p.NewCartItemServiceClient(conn),
        Order_P:               p.NewOrderProductServiceClient(conn),
        OrderRecommendation_P: p.NewOrderRecommendationServiceClient(conn),
		Product_P:             p.NewProductServiceClient(conn),
        User:                  u.NewAuthServiceClient(conn),
	}
}
