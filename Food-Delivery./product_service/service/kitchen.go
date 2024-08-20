package service

import (
	"context"

	pb "delivery/product_service/genproto"
	"delivery/product_service/storage"
)

type KitchenService struct {
	stg storage.StorageI
	pb.UnimplementedKitchenServiceServer
}

func NewKitchenService(stg storage.StorageI) *KitchenService {
	return &KitchenService{stg: stg}
}

func (k *KitchenService) CreateKitchen(ctx context.Context, req *pb.KitchenCreateReq) (*pb.Void, error) {
	_, err := k.stg.Kitchen().CreateKitchen(req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (k *KitchenService) GetKitchen(ctx context.Context, req *pb.ById) (*pb.KitchenGet, error) {
	return k.stg.Kitchen().GetKitchen(req)
}

func (k *KitchenService) UpdateKitchen(ctx context.Context, req *pb.KitchenCreateReq) (*pb.Void, error) {
	_, err := k.stg.Kitchen().UpdateKitchen(req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (k *KitchenService) DeleteKitchen(ctx context.Context, req *pb.ById) (*pb.Void, error) {
	_, err := k.stg.Kitchen().DeleteKitchen(req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (k *KitchenService) ListKitchens(ctx context.Context, req *pb.KitchenFilter) (*pb.KitchenList, error) {
	return k.stg.Kitchen().ListKitchens(req)
}
