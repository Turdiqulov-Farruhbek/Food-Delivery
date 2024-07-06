package service

import (
	"context"
	product "product_ordering/genproto"
	"product_ordering/storage"
)

// ProductService RPC xizmatini taqdim etadi
type ProductService struct {
	store storage.StorageInterface
	product.UnimplementedProductServiceServer
}

// NewProductService yangi ProductService ni yaratadi
func NewProductService(store storage.StorageInterface) *ProductService {
	return &ProductService{store: store}
}

// CreateProduct RPC chaqiruvini bajaradi va mahsulot yaratadi
func (s *ProductService) CreateProduct(ctx context.Context, req *product.CreateProductRequest) (*product.ProductResponse, error) {
	return s.store.Product().Create(ctx, req)
}

// GetProduct RPC chaqiruvini bajaradi va mahsulot ma'lumotlarini qaytaradi
func (s *ProductService) GetProduct(ctx context.Context, req *product.ProductRequest) (*product.ProductResponse, error) {
	return s.store.Product().Get(ctx, req)
}

// UpdateProduct RPC chaqiruvini bajaradi va mahsulotni yangilaydi
func (s *ProductService) UpdateProduct(ctx context.Context, req *product.UpdateProductRequest) (*product.ProductResponse, error) {
	return s.store.Product().Update(ctx, req)
}

// DeleteProduct RPC chaqiruvini bajaradi va mahsulotni o'chiradi
func (s *ProductService) DeleteProduct(ctx context.Context, req *product.ProductRequest) (*product.ProductResponse, error) {
	return s.store.Product().Delete(ctx, req)
}

// ListProducts RPC chaqiruvini bajaradi va mahsulotlar ro'yxatini qaytaradi
func (s *ProductService) ListProducts(ctx context.Context, req *product.Empty) (*product.ProductListResponse, error) {
	return s.store.Product().List(ctx, req)
}
