package service

import (
	"context"

	pb "delivery/product_service/genproto"
	"delivery/product_service/storage"
)

type LocationService struct {
	stg storage.StorageI
	pb.UnimplementedLocationServiceServer
}

func NewLocationService(stg storage.StorageI) *LocationService {
	return &LocationService{stg: stg}
}

func (s *LocationService) CreateLocation(ctx context.Context, req *pb.LocationCreateReq) (*pb.Void, error) {
	_, err := s.stg.Location().CreateLocation(req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (s *LocationService) GetLocation(ctx context.Context, req *pb.ById) (*pb.Location, error) {
	return s.stg.Location().GetLocation(req)
}

func (s *LocationService) UpdateLocation(ctx context.Context, req *pb.LocationCreateReq) (*pb.Void, error) {
	_, err := s.stg.Location().UpdateLocation(req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
