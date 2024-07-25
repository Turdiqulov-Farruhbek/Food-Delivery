package storage

import(
	"context"
	courierOrder "courier_delivery/genproto"

)


type StorageInterface interface {
	CourierOrder() CourierOrderInterface
	Courier() CourierInterface
	Order() OrderInterface
}



type CourierOrderInterface interface {
	CreateCourierOrder (ctx context.Context, req *courierOrder.CreateCourierOrderRequest) (*courierOrder.CourierOrderResponse, error)
	GetCourierOrder (ctx context.Context, req *courierOrder.CourierOrderRequest) (*courierOrder.CourierOrderResponse, error)
	UpdateCourierOrder (ctx context.Context, req *courierOrder.UpdateCourierOrderRequest) (*courierOrder.CourierOrderResponse, error)
	DeleteCourierOrder (ctx context.Context, req *courierOrder.CourierOrderRequest) (*courierOrder.CourierOrderResponse, error)
	ListCourierOrders (ctx context.Context, req *courierOrder.Empty) (*courierOrder.CourierOrderListResponse, error)
}

type CourierInterface interface {
	CreateCourier (ctx context.Context, req *courierOrder.CreateCourierRequest) (*courierOrder.CourierResponse, error)
	GetCourier (ctx context.Context, req *courierOrder.CourierRequest) (*courierOrder.CourierResponse, error)
	UpdateCourier (ctx context.Context, req *courierOrder.UpdateCourierRequest) (*courierOrder.CourierResponse, error)
	DeleteCourier (ctx context.Context, req *courierOrder.CourierRequest) (*courierOrder.CourierResponse, error)
	ListCouriers (ctx context.Context, req *courierOrder.Empty) (*courierOrder.CourierListResponse, error)
}

type OrderInterface interface {
	CreateOrder (ctx context.Context, req *courierOrder.CreateOrderRequest) (*courierOrder.OrderResponse, error)
	GetOrder (ctx context.Context, req *courierOrder.OrderRequest) (*courierOrder.OrderResponse, error)
	UpdateOrder (ctx context.Context, req *courierOrder.UpdateOrderRequest) (*courierOrder.OrderResponse, error)
	DeleteOrder (ctx context.Context, req *courierOrder.OrderRequest) (*courierOrder.OrderResponse, error)
	ListOrders (ctx context.Context, req *courierOrder.Empty) (*courierOrder.OrderListResponse, error)
}