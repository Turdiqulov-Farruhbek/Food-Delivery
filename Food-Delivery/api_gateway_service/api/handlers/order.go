package handlers

import (
	"context"
	"net/http"
	"strconv"
	"gateway/genproto/courier"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new order
// @Description Add a new order to the system
// @Tags Orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param order body courier.CreateOrderRequest true "Order Data"
// @Success 201 {object} courier.CreateOrderResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/orders/create [post]
func (h *Handler) CreateOrder(c *gin.Context) {
	var req courier.CreateOrderRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Order_C.CreateOrder(context.Background(), &req)
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
// @Success 200 {object} courier.GetOrderResponse
// @Failure 400 {object} string "Invalid Order ID"
// @Failure 404 {object} string "Order Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/orders/{orderId} [get]
func (h *Handler) GetOrder(c *gin.Context) {
	id := c.Param("orderId")
	req := &courier.OrderRequest{OrderId: id}
	resp, err := h.Order_C.GetOrder(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get all orders
// @Description Retrieve a list of all orders
// @Tags Orders
// @Produce json
// @Security BearerAuth
// @Success 200 {object} courier.GetOrdersResponse
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/orders [get]
func (h *Handler) GetOrders(c *gin.Context) {
	req := &courier.GetRecommendedOrdersRequest{}
	resp, err := h.Order_C.ListOrders(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Update order by ID
// @Description Update details of an existing order by its ID
// @Tags Orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param orderId path string true "Order ID"
// @Param order body courier.UpdateOrderRequest true "Order Data"
// @Success 200 {object} courier.UpdateOrderResponse
// @Failure 400 {object} string "Invalid Order ID"
// @Failure 404 {object} string "Order Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/orders/{orderId} [put]
func (h *Handler) UpdateOrder(c *gin.Context) {
	id := c.Param("orderId")
	var req courier.UpdateOrderRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.OrderId = id
	resp, err := h.Order_C.UpdateOrder(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Delete order by ID
// @Description Remove an order from the system by its ID
// @Tags Orders
// @Produce json
// @Security BearerAuth
// @Param orderId path string true "Order ID"
// @Success 200 {object} string "Order deleted successfully"
// @Failure 400 {object} string "Invalid Order ID"
// @Failure 404 {object} string "Order Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/orders/{orderId} [delete]
func (h *Handler) DeleteOrder(c *gin.Context) {
	id := c.Param("orderId")
	req := &courier.OrderRequest{OrderId: id}
	_, err := h.Order_C.DeleteOrder(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

// @Summary Get recommended orders for a courier
// @Description Retrieve a list of recommended orders for a courier
// @Tags Orders
// @Produce json
// @Security BearerAuth
// @Param courierId query string true "Courier ID"
// @Success 200 {object} courier.GetRecommendedOrdersResponse
// @Failure 400 {object} string "Invalid Courier ID"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/recommended-orders [get]
func (h *Handler) GetRecommendedOrders(c *gin.Context) {
	req := &courier.GetRecommendedOrdersRequest{}
	resp, err := h.Order_C.GetRecommendedOrders(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Update order status
// @Description Update the status of an existing order
// @Tags Orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Order ID"
// @Param status body courier.UpdateOrderStatusRequest true "Order Status"
// @Success 200 {object} courier.UpdateOrderStatusResponse
// @Failure 400 {object} string "Invalid Order ID"
// @Failure 404 {object} string "Order Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/orders/{id}/status [put]
func (h *Handler) UpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")
	var req courier.UpdateOrderStatusRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.OrderId = id
	resp, err := h.Order_C.UpdateOrderStatus(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get courier order history
// @Description Retrieve the order history for a specific courier
// @Tags Orders
// @Produce json
// @Security BearerAuth
// @Param courierId query string true "Courier ID"
// @Success 200 {object} courier.GetCourierOrderHistoryResponse
// @Failure 400 {object} string "Invalid Courier ID"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/orders/history [get]
func (h *Handler) GetCourierOrderHistory(c *gin.Context) {
	courierId := c.Query("courierId")
	id, err := strconv.Atoi(courierId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Courier ID"})
		return
	}
	req := &courier.GetCourierOrderHistoryRequest{CourierId: int32(id)}
	resp, err := h.Order_C.GetCourierOrderHistory(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
