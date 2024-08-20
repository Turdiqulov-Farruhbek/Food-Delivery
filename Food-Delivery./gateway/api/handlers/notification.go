package handlers

import (
	"net/http"

	_ "delivery/gateway/docs"
	pb "delivery/gateway/genproto"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// @Summary Create a Notification
// @Description Create a new notification and produce a message to Kafka
// @Tags Notifications
// @Accept  json
// @Produce  json
// @Security  		BearerAuth
// @Param notification body pb.NotificationCreate true "Notification Data"
// @Success 200 {object} string "Notification created successfully"
// @Failure 400 {object} string "Invalid request data"
// @Failure 500 {object} string "Internal server error"
// @Router /notification/create [post]
func (h *Handler) CreateNotification(c *gin.Context) {
	var req pb.NotificationCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input, err := protojson.Marshal(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.Producer.ProduceMessages("notification-create", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "notification Created successfully"})
}

// @Summary Delete a Notification
// @Description Delete an existing notification by its ID
// @Tags Notifications
// @Produce  json
// @Accepts json
// @Security  		BearerAuth
// @Param notification_id path string true "Notification ID"
// @Success 200 {object} string "Notification deleted successfully"
// @Failure 400 {object} string "Invalid notification ID"
// @Failure 500 {object} string "Internal server error"
// @Router /notification/delete/{notification_id} [delete]
func (h *Handler) DeleteNotification(c *gin.Context) {
	id := c.Param("notification_id")
	req := pb.ById{Id: id}
	_, err := h.Clients.NotificationC.DeleteNotification(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "notification deleted successfully"})
}

// @Summary Update a Notification
// @Description Update an existing notification by its ID
// @Tags Notifications
// @Accept  json
// @Produce  json
// @Security  		BearerAuth
// @Param notification_id path string true "Notification ID"
// @Param notification body pb.NotificationUpt true "Notification Data"
// @Success 200 {object} string "Notification updated successfully"
// @Failure 400 {object} string "Invalid request data"
// @Failure 500 {object} string "Internal server error"
// @Router /notification/update/{notification_id} [put]
func (h *Handler) UpdateNotification(c *gin.Context) {
	var body pb.NotificationUpt
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("notification_id")
	req := pb.NotificationUpdate{NotificationId: id, Body: &body}

	_, err := h.Clients.NotificationC.UpdateNotification(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "notification updated successfully"})
}

// @Summary Get Notifications
// @Description Retrieve notifications based on provided filters
// @Tags Notifications
// @Accept  json
// @Produce  json
// @Security  		BearerAuth
// @Param filter body pb.NotifFilter true "Filter Data"
// @Success 200 {object} pb.NotificationList "List of notifications"
// @Failure 400 {object} string "Invalid filter parameters"
// @Failure 500 {object} string "Internal server error"
// @Router /notification/all-notifications [get]
func (h *Handler) GetNotifications(c *gin.Context) {
	var req pb.NotifFilter
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.Clients.NotificationC.GetNotifications(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Get Notification by ID
// @Description Retrieve a notification by its ID
// @Tags Notifications
// @Produce  json
// @Accept  json
// @Security  		BearerAuth
// @Param notification_id path string true "Notification ID"
// @Success 200 {object} pb.NotificationGet "Notification details"
// @Failure 400 {object} string "Invalid notification ID"
// @Failure 500 {object} string "Internal server error"
// @Router /notification/get/{notification_id} [get]
func (h *Handler) GetNotificationById(c *gin.Context) {
	id := c.Param("notification_id")
	req := pb.ById{Id: id}
	res, err := h.Clients.NotificationC.GetNotification(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Notify All Users
// @Description Send a notification to all users
// @Tags Notifications
// @Accept  json
// @Produce  json
// @Security  		BearerAuth
// @Param notification body pb.NotificationMessage true "Notification Message"
// @Success 200 {object} string "Notification sent to all users"
// @Failure 400 {object} string "Invalid request data"
// @Failure 500 {object} string "Internal server error"
// @Router /notification/notify-all [post]
func (h *Handler) NotifyAll(c *gin.Context) {
	var body pb.NotificationMessage
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input, err := protojson.Marshal(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.Producer.ProduceMessages("notify-all", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "all group notified"})

}
