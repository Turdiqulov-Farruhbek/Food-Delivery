package handlers

import (
	"net/http"
	"strconv"

	pb "delivery/gateway/genproto"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// AddItem adds a new item
// @Summary 		Add Item
// @Description 	Add a new item
// @Tags 			Item,User
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		user_id path    string       true   "User ID"
// @Param     		item    body    pb.ItemCreate  true   "Item data"
// @Success 		200  {string}  string "Item added successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/item/{user_id}/create [post]
func (h *Handler) AddItem(c *gin.Context) {
	var body pb.ItemCreate
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user_id := c.Param("user_id")
	req := pb.ItemCreateReq{
		UserId: user_id,
		Body:   &body,
	}

	input, err := protojson.Marshal(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.Producer.ProduceMessages("item-create", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "item has ben created succesfully"})
}

// GetItem retrieves an item by ID
// @Summary 		Get Item
// @Description 	Get an item by ID
// @Tags 			Item,User
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		item_id  path    string      true   "Item ID"
// @Success 		200  {object}  pb.ItemGet "Item data"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/item/get/{item_id} [get]
func (h *Handler) GetItem(c *gin.Context) {
	id := c.Param("item_id")
	req := pb.ById{Id: id}
	res, err := h.Clients.ItemC.GetItem(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateItem updates an existing item
// @Summary 		Update Item
// @Description 	Update an existing item
// @Tags 			Item,User,Admin
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		item_id  path    string       true   "Item ID"
// @Param     		item     body    pb.ItemCreate  true   "Item data"
// @Success 		200  {string}  string "Item updated successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/item/update/{item_id} [put]
func (h *Handler) UpdateItem(c *gin.Context) {
	var body pb.ItemCreate
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("item_id")
	req := pb.ItemCreateReq{ItemId: id, Body: &body}

	input, err := protojson.Marshal(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.Producer.ProduceMessages("item-update", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "item has been updated"})
}

// DeleteItem deletes an item by ID
// @Summary 		Delete Item
// @Description 	Delete an item by ID
// @Tags 			Item,User,Admin
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		item_id  path    string      true   "Item ID"
// @Success 		200  {string}  string "Item deleted successfully"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/item/delete/{item_id} [delete]
func (h *Handler) DeleteItem(c *gin.Context) {
	id := c.Param("item_id")
	req := pb.ById{Id: id}
	_, err := h.Clients.ItemC.DeleteItem(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "item deleted successfully"})
}

// ListItems lists items based on filters
// @Summary 		List Items
// @Description 	List items based on filters
// @Tags 			Item,User,Admin
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		user_id      path    string   flase   "User ID"
// @Param     		product_id   query   string   false  "Product ID"
// @Param     		order_id     query   string   false  "Order ID"
// @Param     		limit        query   int      false  "Limit"
// @Param     		offset       query   int      false  "Offset"
// @Param     		min_price    query   int      false  "Min Price"
// @Param     		max_price    query   int      false  "Max Price"
// @Success 		200  {object}  pb.ItemList "List of items"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/item/{user_id}/all-items [get]
func (h *Handler) ListItems(c *gin.Context) {
	user_id := c.Param("user_id")
	product_id := c.Query("product_id")
	order_id := c.Query("order_id")
	limit_str := c.Query("limit")
	offset_str := c.Query("offset")
	min_price_str := c.Query("min_price")
	max_price_str := c.Query("max_price")

	var limit, offset, min_price, max_price int
	var err error

	if limit_str != "" && limit_str != "string" {
		limit, err = strconv.Atoi(limit_str)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
			return
		}
	}
	if offset_str != "" && offset_str != "string" {
		offset, err = strconv.Atoi(offset_str)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset parameter"})
			return
		}
	}
	if min_price_str != "" && min_price_str != "string" {
		min_price, err = strconv.Atoi(min_price_str)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid min_price parameter"})
			return
		}
	}
	if max_price_str != "" && max_price_str != "string" {
		max_price, err = strconv.Atoi(max_price_str)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid max_price parameter"})
			return
		}
	}
	req := pb.ItemFilter{
		UserId:    user_id,
		ProductId: product_id,
		OrderId:   order_id,
		Limit:     int32(limit),
		Offset:    int32(offset),
		MinPrice:  float32(min_price),
		MaxPrice:  float32(max_price),
	}
	res, err := h.Clients.ItemC.ListItems(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
