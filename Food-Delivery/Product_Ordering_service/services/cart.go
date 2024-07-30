package service

import (
	"context"
	gen "product_ordering/genproto/product"
	"product_ordering/storage"
)

type CartService struct {
	store storage.StorageInterface
	gen.UnimplementedCartServiceServer
}

func NewCartService(store storage.StorageInterface) *CartService {
	return &CartService{store: store}
}

func (s *CartService) CreateCart(ctx context.Context, req *gen.CreateCartRequest) (*gen.CartResponse, error) {
	return s.store.Cart().Create(ctx, req)
}

func (s *CartService) GetCart(ctx context.Context, req *gen.CartRequest) (*gen.CartResponse, error) {
	return s.store.Cart().Get(ctx, req)
}

func (s *CartService) UpdateCart(ctx context.Context, req *gen.UpdateCartRequest) (*gen.CartResponse, error) {
	return s.store.Cart().Update(ctx, req)
}

func (s *CartService) DeleteCart(ctx context.Context, req *gen.CartRequest) (*gen.CartResponse, error) {
	return s.store.Cart().Delete(ctx, req)
}

func (s *CartService) ListCarts(ctx context.Context, req *gen.Empty) (*gen.CartListResponse, error) {
	return s.store.Cart().List(ctx, req)
}
