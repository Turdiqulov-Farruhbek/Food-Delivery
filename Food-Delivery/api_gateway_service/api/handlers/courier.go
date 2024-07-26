package handlers

import (
	"context"
	"net/http"

	"gateway/genproto/courier"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new courier
// @Description Add a new courier to the system
// @Tags Couriers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param courier body courier.CreateCourierRequest true "Courier Data"
// @Success 201 {object} courier.CreateCourierResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/couriers/create [post]
func (h *Handler) CreateCourier(c *gin.Context) {
	var req courier.CreateCourierRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Courier_C.CreateCourier(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// @Summary Get courier by ID
// @Description Retrieve details of a courier by their ID
// @Tags Couriers
// @Produce json
// @Security BearerAuth
// @Param courierId path string true "Courier ID"
// @Success 200 {object} courier.GetCourierResponse
// @Failure 400 {object} string "Invalid Courier ID"
// @Failure 404 {object} string "Courier Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/couriers/{courierId} [get]
func (h *Handler) GetCourier(c *gin.Context) {
	id := c.Param("courierId")
	req := &courier.CourierRequest{CourierId: id}
	resp, err := h.Courier_C.GetCourier(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Update courier by ID
// @Description Update details of an existing courier by their ID
// @Tags Couriers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param courierId path string true "Courier ID"
// @Param courier body courier.UpdateCourierRequest true "Courier Data"
// @Success 200 {object} courier.UpdateCourierResponse
// @Failure 400 {object} string "Invalid Courier ID"
// @Failure 404 {object} string "Courier Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/couriers/{courierId} [put]
func (h *Handler) UpdateCourier(c *gin.Context) {
	id := c.Param("courierId")
	var req courier.UpdateCourierRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.CourierId = id
	resp, err := h.Courier_C.UpdateCourier(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Delete courier by ID
// @Description Remove a courier from the system by their ID
// @Tags Couriers
// @Produce json
// @Security BearerAuth
// @Param courierId path string true "Courier ID"
// @Success 200 {object} string "Courier deleted successfully"
// @Failure 400 {object} string "Invalid Courier ID"
// @Failure 404 {object} string "Courier Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/couriers/{courierId} [delete]
func (h *Handler) DeleteCourier(c *gin.Context) {
	id := c.Param("courierId")
	req := &courier.CourierRequest{CourierId: id}
	_, err := h.Courier_C.DeleteCourier(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Courier deleted successfully"})
}

