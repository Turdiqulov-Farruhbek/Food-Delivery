package handlers

import (
	"net/http"
	"strconv"

	pb "delivery/gateway/genproto"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// @Summary List Orders
// @Description Retrieve orders based on provided filters
// @Tags Orders, Admin
// @Accept  json
// @Produce  json
// @Security  		BearerAuth
// @Param user_id query string false "User ID"
// @Param status query string false "Status"
// @Param address query string false "Address"
// @Param courier_id query string false "Courier ID"
// @Param min_total query number false "Minimum Total"
// @Param max_total query number false "Maximum Total"
// @Param time_from query string false "Time From"
// @Param time_to query string false "Time To"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.OrderList "List of orders"
// @Failure 400 {object} string "Invalid filter parameters"
// @Failure 500 {object} string "Internal server error"
// @Router /order/all-orders [get]
func (h *Handler) ListOrders(c *gin.Context) {
	var req pb.OrderFilter
	req.Filter = &pb.Filter{}

	req.UserId = c.Query("user_id")
	req.Status = c.Query("status")
	req.Address = c.Query("address")
	req.CourierId = c.Query("courier_id")
	req.TimeFrom = c.Query("time_from")
	req.TimeTo = c.Query("time_to")

	minTotalStr := c.Query("min_total")
	maxTotalStr := c.Query("max_total")
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	var minTotal, maxTotal, limit, offset int
	var err error
	if minTotalStr != "" && minTotalStr != "string" {
		minTotal, err = strconv.Atoi(minTotalStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid min_total parameter"})
			return
		}
	}
	if maxTotalStr != "" && maxTotalStr != "string" {
		maxTotal, err = strconv.Atoi(maxTotalStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid max_total parameter"})
			return
		}
	}
	if limitStr != "" && limitStr != "string" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
			return
		}
	}
	if offsetStr != "" && offsetStr != "string" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset parameter"})
			return
		}
	}
	req.MinTotal = float32(minTotal)
	req.MaxTotal = float32(maxTotal)
	req.Filter.Limit = int32(limit)
	req.Filter.Offset = int32(offset)

	res, err := h.Clients.OrderC.ListOrders(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Assign an Order
// @Description Assign an order to a specific entity
// @Tags Orders, Admin
// @Accept  json
// @Produce  json
// @Security  		BearerAuth
// @Param order body pb.OrderAssign true "Order Assignment Data"
// @Success 200 {object} string "Order assigned successfully"
// @Failure 400 {object} string "Invalid request data"
// @Failure 500 {object} string "Internal server error"
// @Router /order/assign [post]
func (h *Handler) AssignOrder(c *gin.Context) {
	var body pb.OrderAssign
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.Clients.OrderC.AssignOrder(c, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "order assigned successfully"})
}

// @Summary Create a New Order
// @Description Create a new order and produce a message to Kafka
// @Tags Orders, User
// @Accept  json
// @Produce  json
// @Security  		BearerAuth
// @Param user_id path string true "User ID"
// @Param order body pb.OrderCreate true "Order Data"
// @Success 200 {object} string "Order created successfully"
// @Failure 400 {object} string "Invalid request data"
// @Failure 500 {object} string "Internal server error"
// @Router /order/create/{user_id} [post]
func (h *Handler) CreateOrder(c *gin.Context) {
	var body pb.OrderCreate
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user_id := c.Param("user_id")
	req := pb.OrderCreateReq{
		UserId: user_id,
		Body:   &body,
	}

	input, err := protojson.Marshal(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.Producer.ProduceMessages("order-create", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "order created sucesfully"})

}

// @Summary Get Order by ID
// @Description Retrieve an order by its ID
// @Tags Orders, User
// @Produce  json
// @Accepts json
// @Security  		BearerAuth
// @Param order_id path string true "Order ID"
// @Success 200 {object} pb.OrderGet "Order details"
// @Failure 400 {object} string "Invalid order ID"
// @Failure 500 {object} string "Internal server error"
// @Router /order/get/{order_id} [get]
func (h *Handler) GetOrder(c *gin.Context) {
	order_id := c.Param("order_id")
	req := pb.ById{Id: order_id}
	res, err := h.Clients.OrderC.GetOrder(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Update an Order
// @Description Update an existing order by its ID and produce a message to Kafka
// @Tags Orders, User
// @Accept  json
// @Produce  json
// @Security  		BearerAuth
// @Param order_id path string true "Order ID"
// @Param order body pb.OrderUpt true "Order Data"
// @Success 200 {object} string "Order updated successfully"
// @Failure 400 {object} string "Invalid request data"
// @Failure 500 {object} string "Internal server error"
// @Router /order/update/{order_id} [put]
func (h *Handler) UpdateOrder(c *gin.Context) {
	var body pb.OrderUpt
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order_id := c.Param("order_id")
	req := pb.OrderUpdate{
		Id:   order_id,
		Body: &body,
	}

	input, err := protojson.Marshal(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.Producer.ProduceMessages("order-update", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "order updated"})
}

// @Summary Delete an Order
// @Description Delete an existing order by its ID
// @Tags Orders, User
// @Produce  json
// @Accepts json
// @Security  		BearerAuth
// @Param order_id path string true "Order ID"
// @Success 200 {object} string "Order deleted successfully"
// @Failure 400 {object} string "Invalid order ID"
// @Failure 500 {object} string "Internal server error"
// @Router /order/delete/{order_id} [delete]
func (h *Handler) DeleteOrder(c *gin.Context) {
	order_id := c.Param("order_id")
	req := pb.ById{Id: order_id}
	_, err := h.Clients.OrderC.DeleteOrder(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "order deleted successfully"})
}

// @Summary Pick Up an Order
// @Description Mark an order as picked up
// @Tags Orders, Courier
// @Produce  json
// @Accepts json
// @Security  		BearerAuth
// @Param order_id path string true "Order ID"
// @Success 200 {object} string "Order picked up successfully"
// @Failure 400 {object} string "Invalid order ID"
// @Failure 500 {object} string "Internal server error"
// @Router /order/pickup/{order_id} [post]
func (h *Handler) PickUpOrder(c *gin.Context) {
	order_id := c.Param("order_id")
	req := pb.ById{Id: order_id}
	_, err := h.Clients.OrderC.PickUpOrder(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "order picked up successfully"})
}

// @Summary Checks an Order
// @Description Checks out the order
// @Tags Orders
// @Accept  json
// @Produce  json
// @Security  		BearerAuth
// @Param order_id path string true "Order ID"
// @Success 200 {object} pb.OrderReview "order review"
// @Failure 400 {object} string "Invalid request data"
// @Failure 500 {object} string "Internal server error"
// @Router /order/{order_id}/checkout [get]
func (h *Handler) CheckOrder(c *gin.Context) {
	id := c.Param("order_id")
	req := pb.ById{
		Id: id,
	}
	res, err := h.Clients.OrderC.ReviewOrder(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
