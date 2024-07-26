package handlers

import (
	"context"
	"net/http"
	"gateway/genproto/user"
	"github.com/gin-gonic/gin"
)

// @Summary User Registration
// @Description Register a new user in the system
// @Tags Users
// @Accept json
// @Produce json
// @Param user body user.UserRegisterRequest true "User Registration Data"
// @Success 201 {object} user.UserRegisterResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/users/register [post]
func (h *Handler) UserRegister(c *gin.Context) {
	var req user.UserRegisterRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.User.UserRegister(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// @Summary User Login
// @Description Authenticate a user and return a JWT token
// @Tags Users
// @Accept json
// @Produce json
// @Param user body user.UserLoginRequest true "User Login Data"
// @Success 200 {object} user.UserLoginResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/users/login [post]
func (h *Handler) UserLogin(c *gin.Context) {
	var req user.UserLoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.User.UserLogin(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !resp.Success {
		c.JSON(http.StatusUnauthorized, gin.H{"message": resp.Message})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Courier Registration
// @Description Register a new courier in the system
// @Tags Couriers
// @Accept json
// @Produce json
// @Param courier body user.CourierRegisterRequest true "Courier Registration Data"
// @Success 201 {object} user.CourierRegisterResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/couriers/register [post]
func (h *Handler) CourierRegister(c *gin.Context) {
	var req user.CourierRegisterRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.User.CourierRegister(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// @Summary Courier Login
// @Description Authenticate a courier and return a JWT token
// @Tags Couriers
// @Accept json
// @Produce json
// @Param courier body user.CourierLoginRequest true "Courier Login Data"
// @Success 200 {object} user.CourierLoginResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/couriers/login [post]
func (h *Handler) CourierLogin(c *gin.Context) {
	var req user.CourierLoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.User.CourierLogin(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !resp.Success {
		c.JSON(http.StatusUnauthorized, gin.H{"message": resp.Message})
		return
	}
	c.JSON(http.StatusOK, resp)
}
