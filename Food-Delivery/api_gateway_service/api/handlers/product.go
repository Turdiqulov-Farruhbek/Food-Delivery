package handlers

import (
	"context"
	"net/http"
	"gateway/genproto/product"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new product
// @Description Add a new product to the system
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param product body product.CreateProductRequest true "Product Data"
// @Success 201 {object} product.ProductResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/products/create [post]
func (h *Handler) CreateProduct(c *gin.Context) {
	var req product.CreateProductRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Product_P.CreateProduct(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// @Summary Get product by ID
// @Description Retrieve details of a product by its ID
// @Tags Products
// @Produce json
// @Security BearerAuth
// @Param productId path string true "Product ID"
// @Success 200 {object} product.ProductResponse
// @Failure 400 {object} string "Invalid Product ID"
// @Failure 404 {object} string "Product Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/products/{productId} [get]
func (h *Handler) GetProduct(c *gin.Context) {
	id := c.Param("productId")
	req := &product.ProductRequest{ProductId: id}
	resp, err := h.Product_P.GetProduct(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Update product by ID
// @Description Update details of an existing product by its ID
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param productId path string true "Product ID"
// @Param product body product.UpdateProductRequest true "Product Data"
// @Success 200 {object} product.ProductResponse
// @Failure 400 {object} string "Invalid Product ID"
// @Failure 404 {object} string "Product Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/products/{productId} [put]
func (h *Handler) UpdateProduct(c *gin.Context) {
	id := c.Param("productId")
	var req product.UpdateProductRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.ProductId = id
	resp, err := h.Product_P.UpdateProduct(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Delete product by ID
// @Description Remove a product from the system by its ID
// @Tags Products
// @Produce json
// @Security BearerAuth
// @Param productId path string true "Product ID"
// @Success 200 {object} string "Product deleted successfully"
// @Failure 400 {object} string "Invalid Product ID"
// @Failure 404 {object} string "Product Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/products/{productId} [delete]
func (h *Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("productId")
	req := &product.ProductRequest{ProductId: id}
	_, err := h.Product_P.DeleteProduct(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// @Summary List all products
// @Description Retrieve a list of all products in the system
// @Tags Products
// @Produce json
// @Security BearerAuth
// @Success 200 {object} product.ProductListResponse
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/products [get]
func (h *Handler) GetProducts(c *gin.Context) {
	resp, err := h.Product_P.ListProducts(context.Background(), &product.EmptyProduct{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
