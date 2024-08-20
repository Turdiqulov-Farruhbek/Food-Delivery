package api

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"delivery/gateway/api/handlers"
	middlerware "delivery/gateway/api/middleware"
)

// @title Instant Delivery API Documentation
// @version 1.0
// @description API for Instant Delivery resources
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(h *handlers.Handler) *gin.Engine {
	router := gin.Default()

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust for your specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	enforcer, err := casbin.NewEnforcer("./api/casbin/model.conf", "./api/casbin/policy.csv")
	if err != nil {
		panic(err)
	}
	router.Use(middlerware.NewAuth(enforcer))

	hr := router.Group("/users")
	{
		hr.POST("/courier/create", h.CreateCourier)
		hr.POST("/k-manager/create", h.CreateKitchenManager)
		hr.GET("/all-users", h.GetAllUsers)
		hr.POST("/:courier_id/order/accept", h.AcceptOrder)
	}

	kitchen := router.Group("/kitchen")
	{
		kitchen.POST("/:manager_id/create", h.CreateKitchen)
		kitchen.GET("/all-kitchens", h.ListKitchens)
		kitchen.PUT("/update/:manager_id/:kitchen_id", h.UpdateKitchen)
		kitchen.DELETE("/delete/:kitchen_id", h.DeleteKitchen)
		kitchen.GET("/get/:kitchen_id", h.GetKitchenById)
	}
	notification := router.Group("/notification")
	{
		notification.POST("/create", h.CreateNotification)
		notification.GET("/all-notifications", h.GetNotifications)
		notification.GET("/get/:notification_id", h.GetNotificationById)
		notification.PUT("/update/:notification_id", h.UpdateNotification)
		notification.DELETE("/delete/:notification_id", h.DeleteNotification)
		notification.POST("/notify-all", h.NotifyAll)
	}
	order := router.Group("/order")
	{
		order.POST("/create/:user_id", h.CreateOrder)
		order.GET("/all-orders", h.ListOrders)
		order.GET("/get/:order_id", h.GetOrder)
		order.PUT("/update/:order_id", h.UpdateOrder)
		order.DELETE("/delete/:order_id", h.DeleteOrder)
		order.POST("/assign", h.AssignOrder)
		order.POST("/pickup/:order_id", h.PickUpOrder)
		order.GET("/:order_id/checkout", h.CheckOrder)
	}
	product := router.Group("/product")
	{
		product.POST("/create/:kitchen_id", h.CreateProduct)
		product.GET("/get/:product_id", h.GetProduct)
		product.PUT("/update/:product_id", h.UpdateProduct)
		product.DELETE("/delete/:product_id", h.DeleteProduct)
		product.GET("/all-product", h.ListProducts)
		product.POST("/:product_id/upload_photo", h.UploadProductPhoto)
	}
	item := router.Group("/item")
	{
		item.POST("/:user_id/create", h.AddItem)
		item.GET("/get/:item_id", h.GetItem)
		item.PUT("/update/:item_id", h.UpdateItem)
		item.DELETE("/delete/:item_id", h.DeleteItem)
		item.GET("/:user_id/all-items", h.ListItems)
		item.GET("/filter", h.ListItems)
	}
	location := router.Group("/location")
	{
		location.POST("/:courier_id/create", h.CreateLocation)
		location.GET("/get/:courier_id", h.GetCourierLocation)
		location.PUT("/update/:courier_id", h.UpdateCourierLocation)
	}

	return router
}
