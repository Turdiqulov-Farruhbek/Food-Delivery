package handlers

import (
	"context"
	"net/http"
	"gateway/genproto/product"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new order recommendation
// @Description Add a new order recommendation to the system
// @Tags Recommendations
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param recommendation body product.CreateOrderRecommendationRequest true "Order Recommendation Data"
// @Success 201 {object} product.OrderRecommendationResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/recommendations/order [post]
func (h *Handler) CreateOrderRecommendation(c *gin.Context) {
	var req product.CreateOrderRecommendationRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.OrderRecommendation_P.CreateOrderRecommendation(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// @Summary Get order recommendation by ID
// @Description Retrieve details of an order recommendation by its ID
// @Tags Recommendations
// @Produce json
// @Security BearerAuth
// @Param recommendationId path string true "Recommendation ID"
// @Success 200 {object} product.OrderRecommendationResponse
// @Failure 400 {object} string "Invalid Recommendation ID"
// @Failure 404 {object} string "Recommendation Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/recommendations/order/{recommendationId} [get]
func (h *Handler) GetOrderRecommendation(c *gin.Context) {
	id := c.Param("recommendationId")
	req := &product.OrderRecommendationRequest{RecommendationId: id}
	resp, err := h.OrderRecommendation_P.GetOrderRecommendation(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Update an order recommendation
// @Description Modify details of an existing order recommendation
// @Tags Recommendations
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param recommendation body product.UpdateOrderRecommendationRequest true "Order Recommendation Update Data"
// @Success 200 {object} product.OrderRecommendationResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 404 {object} string "Recommendation Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/recommendations/order [put]
func (h *Handler) UpdateOrderRecommendation(c *gin.Context) {
	var req product.UpdateOrderRecommendationRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.OrderRecommendation_P.UpdateOrderRecommendation(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Delete an order recommendation
// @Description Remove an order recommendation from the system by its ID
// @Tags Recommendations
// @Produce json
// @Security BearerAuth
// @Param recommendationId path string true "Recommendation ID"
// @Success 204 {object} string "No Content"
// @Failure 400 {object} string "Invalid Recommendation ID"
// @Failure 404 {object} string "Recommendation Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/recommendations/order/{recommendationId} [delete]
func (h *Handler) DeleteOrderRecommendation(c *gin.Context) {
	id := c.Param("recommendationId")
	req := &product.OrderRecommendationRequest{RecommendationId: id}
	resp, err := h.OrderRecommendation_P.DeleteOrderRecommendation(context.Background(), req)
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

// @Summary List all order recommendations
// @Description Retrieve a list of all order recommendations
// @Tags Recommendations
// @Produce json
// @Security BearerAuth
// @Success 200 {object} product.OrderRecommendationListResponse
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/recommendations/order [get]
func (h *Handler) ListOrderRecommendations(c *gin.Context) {
	resp, err := h.OrderRecommendation_P.ListOrderRecommendations(context.Background(), &product.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
