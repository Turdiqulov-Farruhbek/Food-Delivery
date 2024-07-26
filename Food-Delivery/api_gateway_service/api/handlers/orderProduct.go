package handlers

import (
	"context"
	"net/http"
	"gateway/genproto/product"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new order
// @Description Add a new order to the system
// @Tags Orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param order body product.CreateOrderProductRequest true "Order Data"
// @Success 201 {object} product.OrderProductResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/orders [post]
func (h *Handler) CreateOrderProduct(c *gin.Context) {
	var req product.CreateOrderProductRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Order_P.CreateOrder(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// @Summary Get order by ID
// @Description Retrieve details of an order by its ID
// @Tags Orders
// @Produce json
// @Security BearerAuth
// @Param orderId path string true "Order ID"
// @Success 200 {object} product.OrderProductResponse
// @Failure 400 {object} string "Invalid Order ID"
// @Failure 404 {object} string "Order Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/orders/{orderId} [get]
func (h *Handler) GetOrderProduct(c *gin.Context) {
	id := c.Param("orderId")
	req := &product.OrderProductRequest{OrderId: id}
	resp, err := h.Order_P.GetOrder(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Update an order
// @Description Modify details of an existing order
// @Tags Orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param order body product.UpdateOrderProductRequest true "Order Update Data"
// @Success 200 {object} product.OrderProductResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 404 {object} string "Order Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/orders [put]
func (h *Handler) UpdateOrderProduct(c *gin.Context) {
	var req product.UpdateOrderProductRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Order_P.UpdateOrder(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Delete an order
// @Description Remove an order from the system by its ID
// @Tags Orders
// @Produce json
// @Security BearerAuth
// @Param orderId path string true "Order ID"
// @Success 204 {object} string "No Content"
// @Failure 400 {object} string "Invalid Order ID"
// @Failure 404 {object} string "Order Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/orders/{orderId} [delete]
func (h *Handler) DeleteOrderProduct(c *gin.Context) {
	id := c.Param("orderId")
	req := &product.OrderProductRequest{OrderId: id}
	resp, err := h.Order_P.DeleteOrder(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if resp.Success {
		c.Status(http.StatusNoContent)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": resp.Message})
	}
}

// @Summary List all orders
// @Description Retrieve a list of all orders with optional filters
// @Tags Orders
// @Produce json
// @Security BearerAuth
// @Param userId query string false "User ID Filter"
// @Param status query string false "Order Status Filter"
// @Success 200 {object} product.OrderProductListResponse
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/orders [get]
func (h *Handler) ListOrdersProduct(c *gin.Context) {
	req := &product.OrderProductListRequest{
		UserId: c.Query("userId"),
		Status: product.CardStatus(product.CardStatus_value[c.Query("status")]),
	}
	resp, err := h.Order_P.ListOrders(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
