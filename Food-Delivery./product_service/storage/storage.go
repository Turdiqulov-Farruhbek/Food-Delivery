package storage

import pb "delivery/product_service/genproto"

type StorageI interface {
	Product() ProductI
	Order() OrderI
	OrderItem() OrderItemI
	Kitchen() KitchenI
	Location() LocationI
	Courier() CourierI
}

type CourierI interface {
	AcceptOrder(req *pb.OrderAcept) (*pb.Void, error)
}

type ProductI interface {
	CreateProduct(req *pb.ProductCreateReq) (*pb.Void, error)
	GetProduct(id *pb.ById) (*pb.ProductGet, error)
	UpdateProduct(req *pb.ProductCreateReq) (*pb.Void, error)
	DeleteProduct(id *pb.ById) (*pb.Void, error)
	ListProducts(req *pb.ProductFilter) (*pb.ProductList, error)
}

type OrderI interface {
	CreateOrder(req *pb.OrderCreateReq) (*pb.Void, error)
	GetOrder(id *pb.ById) (*pb.OrderGet, error)
	UpdateOrder(req *pb.OrderUpdate) (*pb.Void, error)
	DeleteOrder(id *pb.ById) (*pb.Void, error)
	ListOrders(filter *pb.OrderFilter) (*pb.OrderList, error)
	AssignOrder(req *pb.OrderAssign) (*pb.Void, error)
}

type OrderItemI interface {
	AddItem(req *pb.ItemCreateReq) (*pb.Void, error)
	GetItem(id *pb.ById) (*pb.ItemGet, error)
	UpdateItem(req *pb.ItemCreateReq) (*pb.Void, error)
	DeleteItem(id *pb.ById) (*pb.Void, error)
	ListItems(filter *pb.ItemFilter) (*pb.ItemList, error)
}

type KitchenI interface {
	CreateKitchen(req *pb.KitchenCreateReq) (*pb.Void, error)
	UpdateKitchen(req *pb.KitchenCreateReq) (*pb.Void, error)
	DeleteKitchen(id *pb.ById) (*pb.Void, error)
	ListKitchens(req *pb.KitchenFilter) (*pb.KitchenList, error)
	GetKitchen(id *pb.ById) (*pb.KitchenGet, error)
}

type LocationI interface {
	CreateLocation(req *pb.LocationCreateReq) (*pb.Void, error)
	GetLocation(id *pb.ById) (*pb.Location, error)
	UpdateLocation(req *pb.LocationCreateReq) (*pb.Void, error)
}
