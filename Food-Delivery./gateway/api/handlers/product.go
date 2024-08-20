package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"mime/multipart"
	"path/filepath"

	pb "delivery/gateway/genproto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"google.golang.org/protobuf/encoding/protojson"
)

// @Summary Create a New Product
// @Description Create a new product and produce a message to Kafka
// @Tags Products, Kitchen Manager
// @Accept  json
// @Produce  json
// @Security  		BearerAuth
// @Param kitchen_id path string true "Kitchen ID"
// @Param product body pb.ProductCreate true "Product Data"
// @Success 200 {object} string "Product created successfully"
// @Failure 400 {object} string "Invalid request data"
// @Failure 500 {object} string "Internal server error"
// @Router /product/create/{kitchen_id} [post]
func (h *Handler) CreateProduct(c *gin.Context) {
	var body pb.ProductCreate
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	kitchen_id := c.Param("kitchen_id")
	req := pb.ProductCreateReq{
		KitchenId: kitchen_id,
		Body:      &body,
	}

	input, err := protojson.Marshal(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.Producer.ProduceMessages("product-create", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "prodyuct created successfully"})
}

// @Summary Get Product by ID
// @Description Retrieve a product by its ID
// @Tags Products, User
// @Produce  json
// @Accepts json
// @Security  		BearerAuth
// @Param product_id path string true "Product ID"
// @Success 200 {object} pb.ProductGet "Product details"
// @Failure 400 {object} string "Invalid product ID"
// @Failure 500 {object} string "Internal server error"
// @Router /product/get/{product_id} [get]
func (h *Handler) GetProduct(c *gin.Context) {
	id := c.Param("product_id")
	req := pb.ById{Id: id}
	res, err := h.Clients.ProductC.GetProduct(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Update a Product
// @Description Update an existing product by its ID and produce a message to Kafka
// @Tags Products, Kitchen Manager
// @Accept  json
// @Produce  json
// @Security  		BearerAuth
// @Param product_id path string true "Product ID"
// @Param product body pb.ProductCreate true "Product Data"
// @Success 200 {object} string "Product updated successfully"
// @Failure 400 {object} string "Invalid request data"
// @Failure 500 {object} string "Internal server error"
// @Router /product/update/{product_id} [put]
func (h *Handler) UpdateProduct(c *gin.Context) {
	var body pb.ProductCreate
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("product_id")
	req := pb.ProductCreateReq{Id: id, Body: &body}

	input, err := protojson.Marshal(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.Producer.ProduceMessages("product-update", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "product has been updaated successfully"})
}

// @Summary Delete a Product
// @Description Delete an existing product by its ID
// @Tags Products, Kitchen Manager
// @Produce  json
// @Accept json
// @Security  		BearerAuth
// @Param product_id path string true "Product ID"
// @Success 200 {object} string "Product deleted successfully"
// @Failure 400 {object} string "Invalid product ID"
// @Failure 500 {object} string "Internal server error"
// @Router /product/delete/{product_id} [delete]
func (h *Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("product_id")
	req := pb.ById{Id: id}
	_, err := h.Clients.ProductC.DeleteProduct(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
}

// @Summary List Products
// @Description Retrieve products based on provided filters
// @Tags Products, User
// @Accept  json
// @Produce  json
// @Security  		BearerAuth
// @Param Name query string false "Product Name"
// @Param MinPrice query number false "Minimum Price"
// @Param MaxPrice query number false "Maximum Price"
// @Param Description query string false "Product Description"
// @Param KitchenId query string false "Kitchen ID"
// @Param Limit query int32 false "Limit"
// @Param Offset query int32 false "Offset"
// @Success 200 {object} pb.ProductList "List of products"
// @Failure 400 {object} string "Invalid filter parameters"
// @Failure 500 {object} string "Internal server error"
// @Router /product/all-product [get]
func (h *Handler) ListProducts(c *gin.Context) {
	var req pb.ProductFilter
	req.Filter = &pb.Filter{}

	req.Name = c.Query("Name")

	if minPrice := c.Query("MinPrice"); minPrice != "" {
		if v, err := strconv.ParseFloat(minPrice, 32); err == nil {
			req.MinPrice = float32(v)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid MinPrice"})
			return
		}
	}

	if maxPrice := c.Query("MaxPrice"); maxPrice != "" {
		if v, err := strconv.ParseFloat(maxPrice, 32); err == nil {
			req.MaxPrice = float32(v)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid MaxPrice"})
			return
		}
	}

	req.Description = c.Query("Description")
	req.KitchenId = c.Query("KitchenId")

	if limit := c.Query("Limit"); limit != "" {
		if v, err := strconv.Atoi(limit); err == nil {
			req.Filter.Limit = int32(v)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Limit"})
			return
		}
	}

	if offset := c.Query("Offset"); offset != "" {
		if v, err := strconv.Atoi(offset); err == nil {
			req.Filter.Offset = int32(v)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Offset"})
			return
		}
	}

	res, err := h.Clients.ProductC.ListProducts(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

type File struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

// File upload
// @Security    BearerAuth
// @Summary File upload
// @Description File upload
// @Tags Products
// @Accept json
// @Produce json
// @Security  		BearerAuth
// @Param file formData file true "File"
// @Param product_id path string true "Product ID"
// @Router /product/{product_id}/upload_photo [post]
// @Success 200 {object} string
func (h *Handler) UploadProductPhoto(c *gin.Context) {
	var file File

	err := c.ShouldBind(&file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if filepath.Ext(file.File.Filename) != ".png" && filepath.Ext(file.File.Filename) != ".jpg" && filepath.Ext(file.File.Filename) != ".jpeg" {
		fmt.Println(filepath.Ext(file.File.Filename) != ".png" || filepath.Ext(file.File.Filename) != ".jpg")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Couln't find matching file format",
		})
		return
	}

	ext := filepath.Ext(file.File.Filename)

	uploadDir := "./minio/media"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err := os.Mkdir(uploadDir, os.ModePerm)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Couln't find matching file format",
			})
			log.Println("Error creating media directory", err.Error())
			return
		}
	}

	id := uuid.New().String()

	newFilename := id + ext
	uploadPath := filepath.Join(uploadDir, newFilename)

	if err := c.SaveUploadedFile(file.File, uploadPath); err != nil {
		c.JSON(500, gin.H{
			"error": "Couln't find matching file format",
		})
		log.Println(err)
		return
	}

	objectName := newFilename
	_, err = h.minio.FPutObject(context.Background(), "product-pictures", objectName, uploadPath, minio.PutObjectOptions{
		ContentType: "image/png",
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": err.Error(),
		})
		log.Println(err)
		return
	}

	minioURL := fmt.Sprintf("http://localhost:9000/product-pictures/%s", objectName)

	req := pb.ProductCreateReq{
		Id: c.Param("product_id"),
		Body: &pb.ProductCreate{
			Image: minioURL,
		},
	}
	input, err := protojson.Marshal(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.Producer.ProduceMessages("product-update", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"url":     minioURL,
	})
}
