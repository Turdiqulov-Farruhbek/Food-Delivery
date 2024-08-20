package handlers

import (
	"net/http"

	pb "delivery/gateway/genproto"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// @Summary Create Location
// @Description Create a new location for a courier
// @Tags Location, Courier
// @Accept  json
// @Produce  json
// @Security  		BearerAuth
// @Param courier_id path string true "Courier ID"
// @Param body body pb.LocationCreate true "Location Create Body"
// @Success 200 {string} string "Location created successfully"
// @Failure 400 {object} string "Invalid request parameters"
// @Failure 500 {object} string "Internal server error"
// @Router /location/{courier_id}/create [post]
func (h *Handler) CreateLocation(c *gin.Context) {
	var body pb.LocationCreate
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	courier_id := c.Param("courier_id")
	req := pb.LocationCreateReq{
		CourierId: courier_id,
		Body:      &body,
	}

	// write to kafka
	input, err := protojson.Marshal(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.Producer.ProduceMessages("location-create", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "location created successfully"})
}

// @Summary Get Courier Location
// @Description Get the current location of a courier
// @Tags Location
// @Accept  json
// @Produce  json
// @Security  		BearerAuth
// @Param courier_id path string true "Courier ID"
// @Success 200 {object} pb.LocationCreate "Courier location"
// @Failure 500 {object} string "Internal server error"
// @Router /location/get/{courier_id} [get]
func (h *Handler) GetCourierLocation(c *gin.Context) {
	courier_id := c.Param("courier_id")
	req := pb.ById{Id: courier_id}
	res, err := h.Clients.LocationC.GetLocation(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Update Courier Location
// @Description Update the location of a courier
// @Tags Location
// @Accept  json
// @Produce  json
// @Security  		BearerAuth
// @Param courier_id path string true "Courier ID"
// @Param body body pb.LocationCreate true "Location Update Body"
// @Success 200 {string} string "Location updated successfully"
// @Failure 400 {object} string "Invalid request parameters"
// @Failure 500 {object} string "Internal server error"
// @Router /location/update/{courier_id} [put]
func (h *Handler) UpdateCourierLocation(c *gin.Context) {
	var body pb.LocationCreate
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	courier_id := c.Param("courier_id")
	req := pb.LocationCreateReq{
		CourierId: courier_id,
		Body:      &body,
	}
	input, err := protojson.Marshal(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.Producer.ProduceMessages("location-update", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "location updated sucesfully"})
}
