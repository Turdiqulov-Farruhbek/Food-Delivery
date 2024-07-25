package handlers

import (
	c "gateway/genproto/courier"
	n "gateway/genproto/notification"
	p "gateway/genproto/product"
	u "gateway/genproto/user"
)


type Handler struct {
	Courier c.CourierServiceClient
	CourierOrder c.CourierOrderServiceClient
	Order c.OrderServiceClient
	AdminAlert n.AdminAlertServiceClient
	CourierNtfn n.CourierNotificationServiceClient
	UserNtfn n.UserNotificationServiceClient
	Cart p.CartServiceClient
	CartItemS p.CartItemServiceClient
	Order p.OrderServiceClient
	OrderRecommendation p.OrderRecommendationServiceClient
	Product p.ProductServiceClient
}