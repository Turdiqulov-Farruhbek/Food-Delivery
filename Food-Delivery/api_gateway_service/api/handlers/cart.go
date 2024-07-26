package handlers

import (
	"context"
	"net/http"
	"gateway/genproto/product"


	"github.com/gin-gonic/gin"
)

// @Summary Create a new cart
// @Description Add a new cart to the system
// @Tags Carts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param cart body product.CreateCartRequest true "Cart Data"
// @Success 201 {object} product.CreateCartResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/carts/create [post]
func (h *Handler) CreateCart(c *gin.Context) {
	var req product.CreateCartRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Cart_P.CreateCart(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// @Summary Get cart by ID
// @Description Retrieve details of a cart by its ID
// @Tags Carts
// @Produce json
// @Security BearerAuth
// @Param cartId path string true "Cart ID"
// @Success 200 {object} product.GetCartResponse
// @Failure 400 {object} string "Invalid Cart ID"
// @Failure 404 {object} string "Cart Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/carts/{cartId} [get]
func (h *Handler) GetCart(c *gin.Context) {
	id := c.Param("cartId")
	req := &product.CartRequest{CartId: id}
	resp, err := h.Cart_P.GetCart(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Update cart by ID
// @Description Update details of an existing cart by its ID
// @Tags Carts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param cartId path string true "Cart ID"
// @Param cart body product.UpdateCartRequest true "Cart Data"
// @Success 200 {object} product.UpdateCartResponse
// @Failure 400 {object} string "Invalid Cart ID"
// @Failure 404 {object} string "Cart Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/carts/{cartId} [put]
func (h *Handler) UpdateCart(c *gin.Context) {
	id := c.Param("cartId")
	var req product.UpdateCartRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.CartId = id
	resp, err := h.Cart_P.UpdateCart(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Delete cart by ID
// @Description Remove a cart from the system by its ID
// @Tags Carts
// @Produce json
// @Security BearerAuth
// @Param cartId path string true "Cart ID"
// @Success 200 {object} string "Cart deleted successfully"
// @Failure 400 {object} string "Invalid Cart ID"
// @Failure 404 {object} string "Cart Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/carts/{cartId} [delete]
func (h *Handler) DeleteCart(c *gin.Context) {
	id := c.Param("cartId")
	req := &product.CartRequest{CartId: id}
	_, err := h.Cart_P.DeleteCart(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cart deleted successfully"})
}


// @Summary Get all carts
// @Description Retrieve a list of all carts
// @Tags Carts
// @Produce json
// @Security BearerAuth
// @Success 200 {object} product.GetAllCartsResponse
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/carts [get]
func (h *Handler) GetAllCarts(c *gin.Context) {
	req := &product.Empty{}
	resp, err := h.Cart_P.ListCarts(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}