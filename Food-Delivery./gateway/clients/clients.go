package clients

import (
	"delivery/gateway/config"
	pb "delivery/gateway/genproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	AuthC         pb.AuthServiceClient
	ProductC      pb.ProductServiceClient
	ItemC         pb.ItemServiceClient
	OrderC        pb.OrderServiceClient
	KitchenC      pb.KitchenServiceClient
	LocationC     pb.LocationServiceClient
	NotificationC pb.NotificationServiceClient
	CourierC      pb.CourierServiceClient
}

func NewClients(cfg *config.Config) (*Clients, error) {
	auth_path := cfg.AuthHost + ":" + cfg.AuthPort
	product_path := cfg.ProductHost + ":" + cfg.ProductPort
	notification_path := cfg.NotificationHost + ":" + cfg.NotificationPort

	conn_auth, err := grpc.NewClient(auth_path, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	conn_product, err := grpc.NewClient(product_path, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	conn_notification, err := grpc.NewClient(notification_path, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	authC := pb.NewAuthServiceClient(conn_auth)
	productC := pb.NewProductServiceClient(conn_product)
	notificationC := pb.NewNotificationServiceClient(conn_notification)
	orderC := pb.NewOrderServiceClient(conn_product)
	kitchenC := pb.NewKitchenServiceClient(conn_product)
	locationC := pb.NewLocationServiceClient(conn_product)
	itemC := pb.NewItemServiceClient(conn_product)
	courierC := pb.NewCourierServiceClient(conn_product)

	return &Clients{
		AuthC:         authC,
		ProductC:      productC,
		NotificationC: notificationC,
		OrderC:        orderC,
		KitchenC:      kitchenC,
		LocationC:     locationC,
		ItemC:         itemC,
		CourierC:      courierC,
	}, nil
}
