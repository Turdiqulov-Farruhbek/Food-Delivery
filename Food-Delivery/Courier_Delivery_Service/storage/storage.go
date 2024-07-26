package storage

import(
	"context"
	courierOrder "courier_delivery/genproto"

)


type StorageCourierInterface interface {
	CourierOrder() CourierOrderInterface
	Courier() CourierInterface
	Order() OrderInterface
	Task() TaskInterface
	CourierLocation() CourierLocationInterface
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
	ListOrders (ctx context.Context, req *courierOrder.GetRecommendedOrdersRequest) (*courierOrder.OrderListResponse, error)
}

type TaskInterface interface {
	CreateTask (ctx context.Context, req *courierOrder.CreateTaskRequest) (*courierOrder.TaskResponse, error)
	GetTask (ctx context.Context, req *courierOrder.GetTaskRequest) (*courierOrder.TaskResponse, error)
	UpdateTask (ctx context.Context, req *courierOrder.UpdateTaskRequest) (*courierOrder.TaskResponse, error)
	DeleteTask (ctx context.Context, req *courierOrder.DeleteTaskRequest) (*courierOrder.DeleteTaskResponse, error)
	GetAllTasks (ctx context.Context, req *courierOrder.GetAllTasksRequest) (*courierOrder.GetAllTasksResponse, error)
}

type CourierLocationInterface interface {
	CreateCourierLocation (ctx context.Context, req *courierOrder.CreateCourierLocationRequest) (*courierOrder.CourierLocationResponse, error)
    GetCourierLocation (ctx context.Context, req *courierOrder.GetCourierLocationRequest) (*courierOrder.CourierLocationResponse, error)
    UpdateCourierLocation (ctx context.Context, req *courierOrder.UpdateCourierLocationRequest) (*courierOrder.CourierLocationResponse, error)
	DeleteCourierLocation (ctx context.Context, req *courierOrder.DeleteCourierLocationRequest) (*courierOrder.DeleteCourierLocationResponse, error)
	GetAllCourierLocations (ctx context.Context, req *courierOrder.GetAllCourierLocationsRequest) (*courierOrder.GetAllCourierLocationsResponse, error)
}