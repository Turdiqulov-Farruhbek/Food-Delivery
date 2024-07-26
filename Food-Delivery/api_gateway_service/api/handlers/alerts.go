package handlers

import (
	"context"
	"net/http"
	"gateway/genproto/notification"


	"github.com/gin-gonic/gin"
)

// @Summary Create a new admin alert
// @Description Add a new admin alert to the system
// @Tags Alerts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param alert body notification.CreateAdminAlertRequest true "Alert Data"
// @Success 201 {object} notification.CreateAdminAlertResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/alerts/admin [post]
func (h *Handler) CreateAdminAlert(c *gin.Context) {
	var req notification.CreateAdminAlertRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.AdminAlert_N.CreateAdminAlert(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// @Summary Get admin alert by ID
// @Description Retrieve details of an admin alert by its ID
// @Tags Alerts
// @Produce json
// @Security BearerAuth
// @Param alertId path string true "Alert ID"
// @Success 200 {object} notification.GetAdminAlertResponse
// @Failure 400 {object} string "Invalid Alert ID"
// @Failure 404 {object} string "Alert Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/alerts/admin/{alertId} [get]
func (h *Handler) GetAdminAlert(c *gin.Context) {
	id := c.Param("alertId")
	req := &notification.AdminAlertRequest{AlertId: id}
	resp, err := h.AdminAlert_N.GetAdminAlert(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

//@Summary Update admin alert by ID
// @Description Update details of an existing admin alert by its ID
// @Tags Alerts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param alertId path string true "Alert ID"
// @Param alert body notification.UpdateAdminAlertRequest true "Alert Data"
// @Success 200 {object} notification.UpdateAdminAlertResponse
// @Failure 400 {object} string "Invalid Alert ID"
// @Failure 404 {object} string "Alert Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/alerts/admin/{alertId} [put]
func (h *Handler) UpdateAdminAlert(c *gin.Context) {
	id := c.Param("alertId")
	var req notification.UpdateAdminAlertRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.AlertId = id
	resp, err := h.AdminAlert_N.UpdateAdminAlert(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Delete admin alert by ID
// @Description Remove an admin alert from the system by its ID
// @Tags Alerts
// @Produce json
// @Security BearerAuth
// @Param alertId path string true "Alert ID"
// @Success 200 {object} string "Alert deleted successfully"
// @Failure 400 {object} string "Invalid Alert ID"
// @Failure 404 {object} string "Alert Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/alerts/admin/{alertId} [delete]
func (h *Handler) DeleteAdminAlert(c *gin.Context) {
	id := c.Param("alertId")
	req := &notification.AdminAlertRequest{AlertId: id}
	_, err := h.AdminAlert_N.DeleteAdminAlert(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Alert deleted successfully"})
}
