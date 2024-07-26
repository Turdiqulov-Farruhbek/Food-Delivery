package api

import (
	handler "gateway/api/handlers"

	// _ "gateway/docs"
	_ "gateway/genproto/courier"
	_ "gateway/genproto/notification"
	_ "gateway/genproto/product"
	_ "gateway/genproto/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

// @title Food Delivery System API
// @version 1.0
// @description API for Food Delivery System resources
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(conn *grpc.ClientConn) *gin.Engine {
	h := handler.NewHandler(conn)

	router := gin.Default()

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust for your specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	courier := router.Group("/couriers")
	{
		courier.POST("/create", h.CreateCourier)
		courier.GET("/:courierId", h.GetCourier)
		courier.PUT("/:courierId", h.UpdateCourier)
		courier.DELETE("/:courierId", h.DeleteCourier)

		router.GET("/recommended-orders", getRecommendedOrders)
		router.PUT("/orders/:id/status", updateOrderStatus)
		router.GET("/orders/history", getCourierOrderHistory)
	}

	order := router.Group("/orders")
	{
		order.POST("/create", h.CreateOrder)
		order.GET("/:orderId", h.GetOrder)
		order.GET("/", h.GetOrders)
		order.PUT("/:orderId", h.UpdateOrder)
		order.DELETE("/:orderId", h.DeleteOrder)
	}

	task := router.Group("/tasks")
	{
		task.POST("/create", h.CreateTask)
		task.GET("/:taskId", h.GetTask)
		task.GET("/", h.GetTasks)
		task.PUT("/:taskId", h.UpdateTask)
		task.DELETE("/:taskId", h.DeleteTask)
	}

	alert := router.Group("/alerts")
	{
		alert.POST("/admin", h.CreateAdminAlert)
		alert.GET("/admin/:alertId", h.GetAdminAlert)
		alert.PUT("/admin/:alertId", h.UpdateAdminAlert)
		alert.DELETE("/admin/:alertId", h.DeleteAdminAlert)
	}

	notification := router.Group("/notifications")
	{
		notification.POST("/courier", h.CreateCourierNotification)
		notification.POST("/user", h.CreateUserNotification)
		notification.GET("/courier/:notificationId", h.GetCourierNotification)
		notification.GET("/user/:notificationId", h.GetUserNotification)
	}

	cart := router.Group("/carts")
	{
		cart.POST("/create", h.CreateCart)
		cart.GET("/:cartId", h.GetCart)
		cart.PUT("/:cartId", h.UpdateCart)
		cart.DELETE("/:cartId", h.DeleteCart)
	}

	product := router.Group("/products")
	{
		product.POST("/create", h.CreateProduct)
		product.GET("/:productId", h.GetProduct)
		product.GET("/", h.GetProducts)
		product.PUT("/:productId", h.UpdateProduct)
		product.DELETE("/:productId", h.DeleteProduct)
	}

	user := router.Group("/users")
	{
		user.POST("/create", h.CreateUser)
		user.GET("/:userId", h.GetUser)
		user.PUT("/:userId", h.UpdateUser)
		user.DELETE("/:userId", h.DeleteUser)
	}

	return router
}
