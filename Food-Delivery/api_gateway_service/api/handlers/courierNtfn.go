package handlers

import (
	"context"
	"gateway/genproto/notification"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new courier notification
// @Description Add a new courier notification to the system
// @Tags Notifications
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param notification body notification.CreateCourierNotificationRequest true "Notification Data"
// @Success 201 {object} notification.CreateCourierNotificationResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/notifications/courier [post]
func (h *Handler) CreateCourierNotification(c *gin.Context) {
	var req notification.CreateCourierNotificationRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.CourierNtfn_N.CreateCourierNotification(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}



// @Summary Get courier notification by ID
// @Description Retrieve details of a courier notification by its ID
// @Tags Notifications
// @Produce json
// @Security BearerAuth
// @Param notificationId path string true "Notification ID"
// @Success 200 {object} notification.GetCourierNotificationResponse
// @Failure 400 {object} string "Invalid Notification ID"
// @Failure 404 {object} string "Notification Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/notifications/courier/{notificationId} [get]
func (h *Handler) GetCourierNotification(c *gin.Context) {
	id := c.Param("notificationId")
	req := &notification.CourierNotificationRequest{NotificationId: id}
	resp, err := h.CourierNtfn_N.GetCourierNotification(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}


// @Summary List all courier notifications
// @Description Retrieve a list of all courier notifications
// @Tags Notifications
// @Produce json
// @Security BearerAuth
// @Success 200 {object} notification.CourierNotificationListResponse
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/notifications/courier [get]
func (h *Handler) ListCourierNotifications(c *gin.Context) {
	resp, err := h.CourierNtfn_N.ListCourierNotifications(context.Background(), &notification.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}


// @Summary Delete a courier notification
// @Description Remove a courier notification from the system by its ID
// @Tags Notifications
// @Produce json
// @Security BearerAuth
// @Param notificationId path string true "Notification ID"
// @Success 204 {object} string "No Content"
// @Failure 400 {object} string "Invalid Notification ID"
// @Failure 404 {object} string "Notification Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/notifications/courier/{notificationId} [delete]
func (h *Handler) DeleteCourierNotification(c *gin.Context) {
	id := c.Param("notificationId")
	req := &notification.CourierNotificationRequest{NotificationId: id}
	resp, err := h.CourierNtfn_N.DeleteCourierNotification(context.Background(), req)
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
