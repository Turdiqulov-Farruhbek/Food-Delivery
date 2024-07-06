package service

import (
	"context"
	"product_ordering/genproto"
	"product_ordering/storage"
)

type CartService struct {
	store storage.StorageInterface
	genproto.UnimplementedCartServiceServer
}

func NewCartService(store storage.StorageInterface) *CartService {
	return &CartService{store: store}
}

func (s *CartService) CreateCart(ctx context.Context, req *genproto.CreateCartRequest) (*genproto.CartResponse, error) {
	return s.store.Cart().Create(ctx, req)
}

func (s *CartService) GetCart(ctx context.Context, req *genproto.CartRequest) (*genproto.CartResponse, error) {
	return s.store.Cart().Get(ctx, req)
}

func (s *CartService) UpdateCart(ctx context.Context, req *genproto.UpdateCartRequest) (*genproto.CartResponse, error) {
	return s.store.Cart().Update(ctx, req)
}

func (s *CartService) DeleteCart(ctx context.Context, req *genproto.CartRequest) (*genproto.CartResponse, error) {
	return s.store.Cart().Delete(ctx, req)
}

func (s *CartService) ListCarts(ctx context.Context, req *genproto.Empty) (*genproto.CartListResponse, error) {
	return s.store.Cart().List(ctx, req)
}
