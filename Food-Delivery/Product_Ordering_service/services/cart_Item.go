package service

import (
	"context"

	cartItem "product_ordering/genproto"
	"product_ordering/storage"
)

// CartItemService RPC xizmatini taqdim etadi
type CartItemService struct {
	store storage.StorageInterface
	cartItem.UnimplementedCartItemServiceServer
}

// NewCartItemService yangi CartItemService ni yaratadi
func NewCartItemService(store storage.StorageInterface) *CartItemService {
	return &CartItemService{store: store}
}

// CreateCartItem RPC chaqiruvini bajaradi va savat mahsulotini yaratadi
func (s *CartItemService) CreateCartItem(ctx context.Context, req *cartItem.CreateCartItemRequest) (*cartItem.CartItemResponse, error) {
	return s.store.CartItem().Create(ctx, req)
}

// GetCartItem RPC chaqiruvini bajaradi va savat mahsuloti ma'lumotlarini qaytaradi
func (s *CartItemService) GetCartItem(ctx context.Context, req *cartItem.CartItemRequest) (*cartItem.CartItemResponse, error) {
	return s.store.CartItem().Get(ctx, req)
}

// UpdateCartItem RPC chaqiruvini bajaradi va savat mahsulotini yangilaydi
func (s *CartItemService) UpdateCartItem(ctx context.Context, req *cartItem.UpdateCartItemRequest) (*cartItem.CartItemResponse, error) {
	return s.store.CartItem().Update(ctx, req)
}

// DeleteCartItem RPC chaqiruvini bajaradi va savat mahsulotini o'chiradi
func (s *CartItemService) DeleteCartItem(ctx context.Context, req *cartItem.CartItemRequest) (*cartItem.CartItemResponse, error) {
	return s.store.CartItem().Delete(ctx, req)
}

// ListCartItems RPC chaqiruvini bajaradi va savat mahsulotlari ro'yxatini qaytaradi
func (s *CartItemService) ListCartItems(ctx context.Context, req *cartItem.Empty) (*cartItem.CartItemListResponse, error) {
	return s.store.CartItem().List(ctx, req)
}
