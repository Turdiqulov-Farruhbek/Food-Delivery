package handlers

import (
	"context"
	"net/http"
	"gateway/genproto/courier"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new courier order
// @Description Add a new courier order to the system
// @Tags CourierOrders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param order body courier.CreateCourierOrderRequest true "Courier Order Data"
// @Success 201 {object} courier.CourierOrderResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/courier-orders [post]
func (h *Handler) CreateCourierOrder(c *gin.Context) {
	var req courier.CreateCourierOrderRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.CourierOrder_C.CreateCourierOrder(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// @Summary Get courier order by ID
// @Description Retrieve details of a courier order by its ID
// @Tags CourierOrders
// @Produce json
// @Security BearerAuth
// @Param orderId path string true "Courier Order ID"
// @Success 200 {object} courier.CourierOrderResponse
// @Failure 400 {object} string "Invalid Order ID"
// @Failure 404 {object} string "Order Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/courier-orders/{orderId} [get]
func (h *Handler) GetCourierOrder(c *gin.Context) {
	id := c.Param("orderId")
	req := &courier.CourierOrderRequest{CourierOrderId: id}
	resp, err := h.CourierOrder_C.GetCourierOrder(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Update a courier order
// @Description Modify details of an existing courier order
// @Tags CourierOrders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param order body courier.UpdateCourierOrderRequest true "Courier Order Update Data"
// @Success 200 {object} courier.CourierOrderResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 404 {object} string "Order Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/courier-orders [put]
func (h *Handler) UpdateCourierOrder(c *gin.Context) {
	var req courier.UpdateCourierOrderRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.CourierOrderId = c.Param("orderId")
	resp, err := h.CourierOrder_C.UpdateCourierOrder(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Delete a courier order
// @Description Remove a courier order from the system by its ID
// @Tags CourierOrders
// @Produce json
// @Security BearerAuth
// @Param orderId path string true "Courier Order ID"
// @Success 204 {object} string "No Content"
// @Failure 400 {object} string "Invalid Order ID"
// @Failure 404 {object} string "Order Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/courier-orders/{orderId} [delete]
func (h *Handler) DeleteCourierOrder(c *gin.Context) {
	id := c.Param("orderId")
	req := &courier.CourierOrderRequest{CourierOrderId: id}
	resp, err := h.CourierOrder_C.DeleteCourierOrder(context.Background(), req)
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

// @Summary List all courier orders
// @Description Retrieve a list of all courier orders
// @Tags CourierOrders
// @Produce json
// @Security BearerAuth
// @Success 200 {object} courier.CourierOrderListResponse
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/courier-orders [get]
func (h *Handler) ListCourierOrders(c *gin.Context) {
	req := &courier.Empty{}
	resp, err := h.CourierOrder_C.ListCourierOrders(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
