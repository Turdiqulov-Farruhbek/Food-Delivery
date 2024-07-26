package handlers

import (
	"context"
	"net/http"
	"gateway/genproto/notification"

	"github.com/gin-gonic/gin"
)



// @Summary Create a new user notification
// @Description Add a new user notification to the system
// @Tags Notifications
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param notification body notification.CreateUserNotificationRequest true "Notification Data"
// @Success 201 {object} notification.CreateUserNotificationResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/notifications/user [post]
func (h *Handler) CreateUserNotification(c *gin.Context) {
	var req notification.CreateUserNotificationRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.UserNtfn_N.CreateUserNotification(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}


// @Summary Get user notification by ID
// @Description Retrieve details of a user notification by its ID
// @Tags Notifications
// @Produce json
// @Security BearerAuth
// @Param notificationId path string true "Notification ID"
// @Success 200 {object} notification.GetUserNotificationResponse
// @Failure 400 {object} string "Invalid Notification ID"
// @Failure 404 {object} string "Notification Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/notifications/user/{notificationId} [get]
func (h *Handler) GetUserNotification(c *gin.Context) {
	id := c.Param("notificationId")
	req := &notification.UserNotificationRequest{NotificationId: id}
	resp, err := h.UserNtfn_N.GetUserNotification(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}


// @Summary List all user notifications
// @Description Retrieve a list of all user notifications
// @Tags Notifications
// @Produce json
// @Security BearerAuth
// @Success 200 {object} notification.UserNotificationListResponse
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/notifications/user [get]
func (h *Handler) ListUserNotifications(c *gin.Context) {
	resp, err := h.UserNtfn_N.ListUserNotifications(context.Background(), &notification.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}


// @Summary Delete a user notification
// @Description Remove a user notification from the system by its ID
// @Tags Notifications
// @Produce json
// @Security BearerAuth
// @Param notificationId path string true "Notification ID"
// @Success 204 {object} string "No Content"
// @Failure 400 {object} string "Invalid Notification ID"
// @Failure 404 {object} string "Notification Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/notifications/user/{notificationId} [delete]
func (h *Handler) DeleteUserNotification(c *gin.Context) {
	id := c.Param("notificationId")
	req := &notification.UserNotificationRequest{NotificationId: id}
	resp, err := h.UserNtfn_N.DeleteUserNotification(context.Background(), req)
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
