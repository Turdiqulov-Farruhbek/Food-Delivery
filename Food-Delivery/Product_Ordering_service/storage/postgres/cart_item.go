package postgres

import (
	"context"
	"errors"

	cartItem "product_ordering/genproto"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type CartItem struct {
	Db *pgx.Conn
}

func NewCartItem(db *pgx.Conn) *CartItem {
	return &CartItem{Db: db}
}

func (c *CartItem) Create(ctx context.Context, req *cartItem.CreateCartItemRequest) (*cartItem.CartItemResponse, error) {
    cartItemID := uuid.NewString()

    query := `INSERT INTO cartitems (cart_item_id, cart_id, product_id, quantity, options) VALUES ($1, $2, $3, $4, $5)`
    _, err := c.Db.Exec(ctx, query, cartItemID, req.CartId, req.ProductId, req.Quantity, req.Options)
    if err != nil {
        return nil, err
    }

    return &cartItem.CartItemResponse{
        Success: true,
        Message: "Cart item created successfully",
        CartItem: &cartItem.CartItem{
            CartItemId: cartItemID,
            CartId:     req.CartId,
            ProductId:  req.ProductId,
            Quantity:   req.Quantity,
            Options:    req.Options,
        },
    }, nil
}

func (c *CartItem) Get(ctx context.Context, req *cartItem.CartItemRequest) (*cartItem.CartItemResponse, error) {
	query := `SELECT cart_item_id, cart_id, product_id, quantity, options FROM cartitems WHERE cart_item_id=$1 AND deleted_at=0`
	var item cartItem.CartItem
	err := c.Db.QueryRow(ctx, query, req.CartItemId).Scan(&item.CartItemId, &item.CartId, &item.ProductId, &item.Quantity, &item.Options)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("cart item not found")
		}
		return nil, err
	}
	return &cartItem.CartItemResponse{
		Success:  true,
		Message:  "Cart item retrieved successfully",
		CartItem: &item,
	}, nil
}


func (c *CartItem) Update(ctx context.Context, req *cartItem.UpdateCartItemRequest) (*cartItem.CartItemResponse, error) {
    // Mavjud qiymatlarni olish
    getQuery := `SELECT cart_id, product_id, quantity, options FROM cartitems WHERE cart_item_id=$1 AND deleted_at=0`
    var current cartItem.CartItem
    err := c.Db.QueryRow(ctx, getQuery, req.CartItemId).Scan(&current.CartId, &current.ProductId, &current.Quantity, &current.Options)
    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, errors.New("cart item not found")
        }
        return nil, err
    }

    // Cart ID ni tekshirish
    var cartID string
    if req.CartId != "" {
        err := c.Db.QueryRow(ctx, "SELECT cart_id FROM carts WHERE cart_id=$1", req.CartId).Scan(&cartID)
        if err != nil {
            return nil, errors.New("cart not found")
        }
    }

    // Agar yangi qiymatlar berilmagan bo'lsa, eski qiymatlarni saqlab qolish
    cartId := req.CartId
    if cartId == "" {
        cartId = current.CartId
    }

    // Unikal cheklovni tekshirish
    checkDuplicateQuery := `SELECT cart_item_id FROM cartitems WHERE cart_id=$1 AND product_id=$2 AND deleted_at=0`
    var duplicateCartItemId string
    err = c.Db.QueryRow(ctx, checkDuplicateQuery, cartId, req.ProductId).Scan(&duplicateCartItemId)
    if err == nil && duplicateCartItemId != req.CartItemId {
        return nil, errors.New("duplicate cart item for the given cart_id and product_id")
    } else if err != pgx.ErrNoRows {
        return nil, err
    }

    // Yangilash
    updateQuery := `UPDATE cartitems SET cart_id=$1, product_id=$2, quantity=$3, options=$4, updated_at=NOW(), deleted_at=0 WHERE cart_item_id=$5`
    _, err = c.Db.Exec(ctx, updateQuery, cartId, req.ProductId, req.Quantity, req.Options, req.CartItemId)
    if err != nil {
        return nil, err
    }

    return &cartItem.CartItemResponse{
        Success: true,
        Message: "Cart item updated successfully",
        CartItem: &cartItem.CartItem{
            CartItemId: req.CartItemId,
            CartId:     cartId,
            ProductId:  req.ProductId,
            Quantity:   req.Quantity,
            Options:    req.Options,
        },
    }, nil
}



func (c *CartItem) Delete(ctx context.Context, req *cartItem.CartItemRequest) (*cartItem.CartItemResponse, error) {
	query := `UPDATE cartitems SET deleted_at=EXTRACT(EPOCH FROM NOW())::BIGINT WHERE cart_item_id=$1 AND deleted_at=0`
	result, err := c.Db.Exec(ctx, query, req.CartItemId)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected() == 0 {
		return nil, errors.New("cart item not found or already deleted")
	}
	return &cartItem.CartItemResponse{
		Success: true,
		Message: "Cart item deleted successfully",
	}, nil
}

func (c *CartItem) List(ctx context.Context, req *cartItem.Empty) (*cartItem.CartItemListResponse, error) {
	query := `SELECT cart_item_id, cart_id, product_id, quantity, options FROM cartitems WHERE deleted_at=0`
	rows, err := c.Db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartItems []*cartItem.CartItem
	for rows.Next() {
		var item cartItem.CartItem
		err := rows.Scan(&item.CartItemId, &item.CartId, &item.ProductId, &item.Quantity, &item.Options)
		if err != nil {
			return nil, err
		}
		cartItems = append(cartItems, &item)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return &cartItem.CartItemListResponse{
		CartItems: cartItems,
	}, nil
}
