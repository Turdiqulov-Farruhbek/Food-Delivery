package handlers

import (
	"net/http"
	"strconv"

	pb "delivery/gateway/genproto"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// @Summary Create a new kitchen
// @Description Create a new kitchen for a manager
// @Tags kitchens, Admin
// @Accept json
// @Produce json
// @Security  		BearerAuth
// @Param manager_id path string true "Manager ID"
// @Param body body pb.KitchenCreate true "Kitchen Create"
// @Success 200 {string} string "Kitchen Created successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /kitchen/{manager_id}/create [post]
func (h *Handler) CreateKitchen(c *gin.Context) {
	var body pb.KitchenCreate
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	manager_id := c.Param("manager_id")
	req := pb.KitchenCreateReq{
		ManagerId: manager_id,
		Body:      &body,
	}

	//write to kafka
	input, err := protojson.Marshal(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.Producer.ProduceMessages("kitchen-create", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "kitchen created successfully"})
}

// @Summary Update a kitchen
// @Description Update a kitchen by ID
// @Tags kitchens, Admin, Manager
// @Accept json
// @Produce json
// @Security  		BearerAuth
// @Param manager_id path string true "Manager ID"
// @Param kitchen_id path string true "Kitchen ID"
// @Param body body pb.KitchenCreate true "Kitchen Create"
// @Success 200 {string} string "Kitchen Updated successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /kitchen/update/{manager_id}/{kitchen_id} [put]
func (h *Handler) UpdateKitchen(c *gin.Context) {
	var body pb.KitchenCreate
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	manger_id := c.Param("manger_id")
	kitchen_id := c.Param("kitchen_id")
	req := pb.KitchenCreateReq{
		ManagerId: manger_id,
		Body:      &body,
		KitchenId: kitchen_id,
	}

	//write to kafka
	input, err := protojson.Marshal(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.Producer.ProduceMessages("kitchen-update", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "kitchen updated successfully"})
}

// @Summary Delete a kitchen
// @Description Delete a kitchen by ID
// @Tags kitchens, Admin
// @Accept json
// @Produce json
// @Security  		BearerAuth
// @Param kitchen_id path string true "Kitchen ID"
// @Success 200 {string} string "Kitchen deleted successfully"
// @Failure 500 {string} string "Internal Server Error"
// @Router /kitchen/delete/{kitchen_id} [delete]
func (h *Handler) DeleteKitchen(c *gin.Context) {
	id := c.Param("kitchen_id")
	req := pb.ById{Id: id}
	_, err := h.Clients.KitchenC.DeleteKitchen(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "kitchen deleted successfully"})
}

// @Summary List kitchens
// @Description List kitchens with filters
// @Tags kitchens, Admin
// @Accept json
// @Produce json
// @Security  		BearerAuth
// @Param manager_id query string false "Kitchen manager id"
// @Param name query string false "Kitchen name"
// @Param phone query string false "Kitchen phone number"
// @Param description query string false "Kitchen Description"
// @Param address query string false "Kitchen Address"
// @Param limit query string false "Limit"
// @Param offset query string false "Ofsset"
// @Success 200 {object} pb.KitchenList "List of kitchens"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /kitchen/all-kitchens [get]
func (h *Handler) ListKitchens(c *gin.Context) {
	manager_id := c.Query("manager_id")
	name := c.Query("name")
	phone := c.Query("phone")
	description := c.Query("description")
	address := c.Query("address")
	limit_str := c.Query("limit")
	offset_str := c.Query("offset")

	var limit, offset int
	if limit_str != "" && limit_str != "string" {
		limit, _ = strconv.Atoi(limit_str)
	}
	if offset_str != "" && offset_str != "string" {
		offset, _ = strconv.Atoi(offset_str)
	}

	lim := pb.Filter{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	req := pb.KitchenFilter{
		ManagerId:   manager_id,
		Name:        name,
		PhoneNumber: phone,
		Description: description,
		Address:     address,
		Filter:      &lim,
	}

	res, err := h.Clients.KitchenC.ListKitchens(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Get a kitchen by ID
// @Description Get a kitchen by ID
// @Tags kitchens
// @Accept json
// @Produce json
// @Security  		BearerAuth
// @Param kitchen_id path string true "Kitchen ID"
// @Success 200 {object} pb.KitchenGet "Kitchen details"
// @Failure 500 {string} string "Internal Server Error"
// @Router /kitchen/get/{kitchen_id} [get]
func (h *Handler) GetKitchenById(c *gin.Context) {
	id := c.Param("kitchen_id")
	req := pb.ById{Id: id}
	res, err := h.Clients.KitchenC.GetKitchen(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
