package storage

import (
	"context"
	cartItem "product_ordering/genproto/product"
)

type StorageInterface interface {
	CartItem() CartItemInterface
	Cart() CartInterface
	OrderRecomend() OrderRecomendInterface
	Order() OrderInterface
	Product() ProductInterface
}

type CartItemInterface interface {
	Create(ctx context.Context, req *cartItem.CreateCartItemRequest) (*cartItem.CartItemResponse, error)
	Get(ctx context.Context, req *cartItem.CartItemRequest) (*cartItem.CartItemResponse, error)
	Update(ctx context.Context, req *cartItem.UpdateCartItemRequest) (*cartItem.CartItemResponse, error)
	Delete(ctx context.Context, req *cartItem.CartItemRequest) (*cartItem.CartItemResponse, error)
	List(ctx context.Context, req *cartItem.CartItemRequest) (*cartItem.CartItemListResponse, error)
}

type CartInterface interface {
	Create(ctx context.Context, req *cartItem.CreateCartRequest) (*cartItem.CartResponse, error)
	Get(ctx context.Context, req *cartItem.CartRequest) (*cartItem.CartResponse, error)
	Update(ctx context.Context, req *cartItem.UpdateCartRequest) (*cartItem.CartResponse, error)
	Delete(ctx context.Context, req *cartItem.CartRequest) (*cartItem.CartResponse, error)
	List(ctx context.Context, req *cartItem.Empty) (*cartItem.CartListResponse, error)
}

type OrderRecomendInterface interface {
	Create(ctx context.Context, req *cartItem.CreateOrderRecommendationRequest) (*cartItem.OrderRecommendationResponse, error)
	Get(ctx context.Context, req *cartItem.OrderRecommendationRequest) (*cartItem.OrderRecommendationResponse, error)
	Update(ctx context.Context, req *cartItem.UpdateOrderRecommendationRequest) (*cartItem.OrderRecommendationResponse, error)
	Delete(ctx context.Context, req *cartItem.OrderRecommendationRequest) (*cartItem.OrderRecommendationResponse, error)
	List(ctx context.Context, req *cartItem.Empty) (*cartItem.OrderRecommendationListResponse, error)
}

type OrderInterface interface {
	Create(ctx context.Context, req *cartItem.CreateOrderProductRequest) (*cartItem.OrderProductResponse, error)
	Get(ctx context.Context, req *cartItem.OrderProductRequest) (*cartItem.OrderProductResponse, error)
	Update(ctx context.Context, req *cartItem.UpdateOrderProductRequest) (*cartItem.OrderProductResponse, error)
	Delete(ctx context.Context, req *cartItem.OrderProductRequest) (*cartItem.OrderProductResponse, error)
	List(ctx context.Context, req *cartItem.OrderProductListRequest) (*cartItem.OrderProductListResponse, error)
}

type ProductInterface interface {
	Create(ctx context.Context, req *cartItem.CreateProductRequest) (*cartItem.ProductResponse, error)
	Get(ctx context.Context, req *cartItem.ProductRequest) (*cartItem.ProductResponse, error)
	Update(ctx context.Context, req *cartItem.UpdateProductRequest) (*cartItem.ProductResponse, error)
	Delete(ctx context.Context, req *cartItem.ProductRequest) (*cartItem.ProductResponse, error)
	List(ctx context.Context, req *cartItem.Empty) (*cartItem.ProductListResponse, error)
}
