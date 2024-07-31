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

func NewHandler(cr, pr, ntfn *grpc.ClientConn) *Handler {
	return &Handler{
        Courier_C:             c.NewCourierServiceClient(cr),
        CourierOrder_C:        c.NewCourierOrderServiceClient(cr),
        Order_C:               c.NewOrderServiceClient(cr),
        CourierLocation_C:     c.NewCourierLocationServiceClient(cr),
        Task_C:                c.NewTaskServiceClient(cr),
        AdminAlert_N:          n.NewAdminAlertServiceClient(ntfn),
        CourierNtfn_N:         n.NewCourierNotificationServiceClient(ntfn),
        UserNtfn_N:            n.NewUserNotificationServiceClient(ntfn),
        Cart_P:                p.NewCartServiceClient(pr),
        CartItemS_P:           p.NewCartItemServiceClient(pr),
        Order_P:               p.NewOrderProductServiceClient(pr),
        OrderRecommendation_P: p.NewOrderRecommendationServiceClient(pr),
		Product_P:             p.NewProductServiceClient(pr),
        User:                  u.NewAuthServiceClient(cr),
	}
}
