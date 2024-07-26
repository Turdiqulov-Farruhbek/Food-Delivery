package handlers

import (
	"context"
	"net/http"
	"gateway/genproto/product"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new cart item
// @Description Add a new item to a cart
// @Tags CartItems
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param cart_item body product.CreateCartItemRequest true "Cart Item Data"
// @Success 201 {object} product.CartItemResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/cart-items/create [post]
func (h *Handler) CreateCartItem(c *gin.Context) {
	var req product.CreateCartItemRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.CartItemS_P.CreateCartItem(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// @Summary Get cart item by ID
// @Description Retrieve details of a cart item by its ID
// @Tags CartItems
// @Produce json
// @Security BearerAuth
// @Param cartItemId path string true "Cart Item ID"
// @Success 200 {object} product.CartItemResponse
// @Failure 400 {object} string "Invalid Cart Item ID"
// @Failure 404 {object} string "Cart Item Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/cart-items/{cartItemId} [get]
func (h *Handler) GetCartItem(c *gin.Context) {
	id := c.Param("cartItemId")
	req := &product.CartItemRequest{CartItemId: id}
	resp, err := h.CartItemS_P.GetCartItem(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Update cart item by ID
// @Description Update details of an existing cart item by its ID
// @Tags CartItems
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param cartItemId path string true "Cart Item ID"
// @Param cart_item body product.UpdateCartItemRequest true "Cart Item Data"
// @Success 200 {object} product.CartItemResponse
// @Failure 400 {object} string "Invalid Cart Item ID"
// @Failure 404 {object} string "Cart Item Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/cart-items/{cartItemId} [put]
func (h *Handler) UpdateCartItem(c *gin.Context) {
	id := c.Param("cartItemId")
	var req product.UpdateCartItemRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.CartItemId = id
	resp, err := h.CartItemS_P.UpdateCartItem(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Delete cart item by ID
// @Description Remove a cart item from the system by its ID
// @Tags CartItems
// @Produce json
// @Security BearerAuth
// @Param cartItemId path string true "Cart Item ID"
// @Success 200 {object} string "Cart Item deleted successfully"
// @Failure 400 {object} string "Invalid Cart Item ID"
// @Failure 404 {object} string "Cart Item Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/cart-items/{cartItemId} [delete]
func (h *Handler) DeleteCartItem(c *gin.Context) {
	id := c.Param("cartItemId")
	req := &product.CartItemRequest{CartItemId: id}
	_, err := h.CartItemS_P.DeleteCartItem(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cart Item deleted successfully"})
}

// @Summary List all cart items
// @Description Retrieve a list of all items in a specific cart
// @Tags CartItems
// @Produce json
// @Security BearerAuth
// @Param cartId query string true "Cart ID"
// @Success 200 {object} product.CartItemListResponse
// @Failure 400 {object} string "Invalid Cart ID"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/cart-items [get]
func (h *Handler) ListCartItems(c *gin.Context) {
	cartId := c.Query("cartId")
	req := &product.CartItemRequest{CartItemId: cartId}
	resp, err := h.CartItemS_P.ListCartItems(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
