package handlers

import (
	"context"
	"net/http"
	"gateway/genproto/courier"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new courier location
// @Description Add a new courier location to the system
// @Tags Courier
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param location body courier.CreateCourierLocationRequest true "Courier Location Data"
// @Success 201 {object} courier.CourierLocationResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/courier-locations/create [post]
func (h *Handler) CreateCourierLocation(c *gin.Context) {
	var req courier.CreateCourierLocationRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.CourierLocation_C.CreateCourierLocation(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// @Summary Get courier location by ID
// @Description Retrieve details of a courier location by its ID
// @Tags Courier
// @Produce json
// @Security BearerAuth
// @Param locationId path string true "Courier Location ID"
// @Success 200 {object} courier.CourierLocationResponse
// @Failure 400 {object} string "Invalid Location ID"
// @Failure 404 {object} string "Location Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/courier-locations/{locationId} [get]
func (h *Handler) GetCourierLocation(c *gin.Context) {
	id := c.Param("locationId")
	req := &courier.GetCourierLocationRequest{Id: id}
	resp, err := h.CourierLocation_C.GetCourierLocation(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Update courier location by ID
// @Description Update details of an existing courier location by its ID
// @Tags Courier
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param locationId path string true "Courier Location ID"
// @Param location body courier.UpdateCourierLocationRequest true "Courier Location Data"
// @Success 200 {object} courier.CourierLocationResponse
// @Failure 400 {object} string "Invalid Location ID"
// @Failure 404 {object} string "Location Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/courier-locations/{locationId} [put]
func (h *Handler) UpdateCourierLocation(c *gin.Context) {
	id := c.Param("locationId")
	var req courier.UpdateCourierLocationRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Id = id
	resp, err := h.CourierLocation_C.UpdateCourierLocation(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Delete courier location by ID
// @Description Remove a courier location from the system by its ID
// @Tags Courier
// @Produce json
// @Security BearerAuth
// @Param locationId path string true "Courier Location ID"
// @Success 200 {object} courier.DeleteCourierLocationResponse
// @Failure 400 {object} string "Invalid Location ID"
// @Failure 404 {object} string "Location Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/courier-locations/{locationId} [delete]
func (h *Handler) DeleteCourierLocation(c *gin.Context) {
	id := c.Param("locationId")
	req := &courier.DeleteCourierLocationRequest{Id: id}
	resp, err := h.CourierLocation_C.DeleteCourierLocation(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get all courier locations
// @Description Retrieve a list of all courier locations in the system
// @Tags Courier
// @Produce json
// @Security BearerAuth
// @Success 200 {object} courier.GetAllCourierResponse
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/courier-locations [get]
func (h *Handler) GetAllCourier(c *gin.Context) {
	resp, err := h.CourierLocation_C.GetAllCourierLocations(context.Background(), &courier.GetAllCourierLocationsRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
