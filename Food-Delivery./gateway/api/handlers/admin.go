package handlers

import (
	"net/http"
	"strconv"

	pb "delivery/gateway/genproto"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateCourier creates a new courier
// @Summary 		Create Courier
// @Description 	Create a new courier
// @Tags 			Admin
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		courier  body    pb.UserCreateReq  true   "Courier data"
// @Success 		201  {object}  string "Courier Created successfully"
// @Failure 		400  {string}  string     "Invalid request"
// @Failure 		500  {string}  string     "Internal server error"
// @Router 			/users/courier/create [post]
func (h *Handler) CreateCourier(c *gin.Context) {
	var req pb.UserCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input, err := protojson.Marshal(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.Producer.ProduceMessages("courier-create", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "courier Created successfully"})
}

// CreateKitchenManager creates a new kitchen manager
// @Summary 		Create Kitchen Manager
// @Description 	Create a new kitchen manager
// @Tags 			Admin
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		kitchen_manager  body    pb.UserCreateReq  true   "Kitchen Manager data"
// @Success 		201  {object}  string "Kitchen Manager Created successfully"
// @Failure 		400  {string}  string     "Invalid request"
// @Failure 		500  {string}  string     "Internal server error"
// @Router 			/users/k-manager/create [post]
func (h *Handler) CreateKitchenManager(c *gin.Context) {
	var req pb.UserCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input, err := protojson.Marshal(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.Producer.ProduceMessages("kitchen-manager-create", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "kitchen manager Created successfully"})
}

// GetAllUsers gets all users based on filters
// @Summary 		Get All Users
// @Description 	Get all users based on filters
// @Tags 			Admin
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		role  query    string  false   "User role"
// @Param     		is_working  query   bool  false   "User role"
// @Param     		limit  query    int32  false   "User limit"
// @Param     		offset  query   int32  false   "User offset"
// @Success 		200  {object}  pb.UserList "List of users"
// @Failure 		400  {string}  string     "Invalid request"
// @Failure 		500  {string}  string     "Internal server error"
// @Router 			/users/all-users [get]
func (h *Handler) GetAllUsers(c *gin.Context) {
	role := c.Query("role")
	is_working := c.Query("is_working")
	limit_str := c.Query("limit")
	offset_str := c.Query("offset")

	var limit, offset int
	var err error
	if limit_str != "" && limit_str != "string" {
		limit, err = strconv.Atoi(limit_str)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
			return
		}
	}
	if offset_str != "" && offset_str != "string" {
		offset, err = strconv.Atoi(offset_str)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
			return
		}
	}
	filter := &pb.Filter{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	req := pb.UserFilter{
		Role:      role,
		IsWorking: is_working,
		Filter:    filter,
	}
	res, err := h.Clients.AuthC.GetAllUsers(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// AcceptOrder accepts an order
// @Summary 		Accept Order
// @Description 	Accept an order
// @Tags 			Order, Courier
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		courier_id  path    string         true   "Courier ID"
// @Param     		order       body    pb.OrderId     true   "Order ID"
// @Success 		200  {object}  string "Order accepted"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/users/{courier_id}/order/accept/ [post]
func (h *Handler) AcceptOrder(c *gin.Context) {
	var order_id pb.OrderId
	if err := c.ShouldBindJSON(&order_id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	courier_id := c.Param("courier_id")
	req := pb.OrderAcept{
		CourierId: courier_id,
		OrderId:   &order_id,
	}

	input, err := protojson.Marshal(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.Producer.ProduceMessages("order-accept", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//read response from kafka
	c.JSON(200, gin.H{"message": "accept request has been made, you willl recieve email if you got th order"})

}

//kitchen
