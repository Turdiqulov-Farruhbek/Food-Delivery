package handlers

import (
	"context"
	"fmt"
	"gateway/genproto/user"
	"net/http"

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

// @Summary Get All Users
// @Description Retrieve a list of all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} user.GetAllUsersResponse
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/users [get]
func (h *Handler) GetAllUsers(c *gin.Context) {
	resp, err := h.User.GetAllUsers(context.Background(), &user.GetAllUsersRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get User By ID
// @Description Retrieve a specific user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} user.UserResponse
// @Failure 404 {object} string "User Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/users/{user_id} [get]
func (h *Handler) GetUser(c *gin.Context) {
	userID := c.Param("user_id")
	fmt.Println(userID)
	resp, err := h.User.GetUser(context.Background(), &user.UserRequest{UserId: userID})
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Update User
// @Description Update user details
// @Tags Users
// @Accept json
// @Produce json
// @Param user body user.UpdateUserRequest true "Update User Data"
// @Success 200 {object} user.UserResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 404 {object} string "User Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/users/{user_id} [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	var req user.UpdateUserRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := c.Param("user_id")
	req.UserId = userID
	resp, err := h.User.UpdateUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Delete User
// @Description Delete a user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} user.UserResponse
// @Failure 404 {object} string "User Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/users/{user_id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	userID := c.Param("user_id")
	resp, err := h.User.DeleteUser(context.Background(), &user.UserRequest{UserId: userID})
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
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

// @Summary Get All Couriers
// @Description Retrieve a list of all couriers
// @Tags Couriers
// @Accept json
// @Produce json
// @Success 200 {object} user.GetAllCouriersResponse
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/couriers [get]
func (h *Handler) GetAllCouriers(c *gin.Context) {
	resp, err := h.User.GetAllCouriers(context.Background(), &user.GetAllCouriersRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// // @Summary Get Courier By ID
// // @Description Retrieve a specific courier by ID
// // @Tags Couriers
// // @Accept json
// // @Produce json
// // @Param courier_id path string true "Courier ID"
// // @Success 200 {object} user.CourierResponse
// // @Failure 404 {object} string "Courier Not Found"
// // @Failure 500 {object} string "Internal Server Error"
// // @Router /api/couriers/{courier_id} [get]
// func (h *Handler) GetCourier(c *gin.Context) {
// 	courierID := c.Param("courier_id")
// 	resp, err := h.User.GetCourier(context.Background(), &user.CourierRequest{CourierId: courierID})
// 	if err != nil {
// 		if err.Error() == "courier not found" {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "Courier not found"})
// 		} else {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		}
// 		return
// 	}
// 	c.JSON(http.StatusOK, resp)
// }

// // @Summary Update Courier
// // @Description Update courier details
// // @Tags Couriers
// // @Accept json
// // @Produce json
// // @Param courier body user.UpdateCourierRequest true "Update Courier Data"
// // @Success 200 {object} user.CourierResponse
// // @Failure 400 {object} string "Invalid Data"
// // @Failure 404 {object} string "Courier Not Found"
// // @Failure 500 {object} string "Internal Server Error"
// // @Router /api/couriers/{courier_id} [put]
// func (h *Handler) UpdateCourier(c *gin.Context) {
// 	var req user.UpdateCourierRequest
// 	if err := c.BindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	courierID := c.Param("courier_id")
// 	req.CourierId = courierID
// 	resp, err := h.User.UpdateCourier(context.Background(), &req)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, resp)
// }

// // @Summary Delete Courier
// // @Description Delete a courier by ID
// // @Tags Couriers
// // @Accept json
// // @Produce json
// // @Param courier_id path string true "Courier ID"
// // @Success 200 {object} user.CourierResponse
// // @Failure 404 {object} string "Courier Not Found"
// // @Failure 500 {object} string "Internal Server Error"
// // @Router /api/couriers/{courier_id} [delete]
// func (h *Handler) DeleteCourier(c *gin.Context) {
// 	courierID := c.Param("courier_id")
// 	resp, err := h.User.DeleteCourier(context.Background(), &user.CourierRequest{CourierId: courierID})
// 	if err != nil {
// 		if err.Error() == "courier not found" {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "Courier not found"})
// 		} else {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		}
// 		return
// 	}
// 	c.JSON(http.StatusOK, resp)
// }
