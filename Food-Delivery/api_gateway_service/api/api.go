package api

import (
	handler "gateway/api/handlers"

	_ "gateway/docs"
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
func NewGin(cr, pr, ntfn *grpc.ClientConn) *gin.Engine {
	h := handler.NewHandler(cr, pr, ntfn)

	router := gin.Default()
	
	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust for your specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	
	user := router.Group("api/users")
	{																																																																																																																																																																																																			
		user.POST("/register", h.UserRegister)
		user.GET("/:userId", h.GetUser)
		user.GET("/", h.GetAllUsers)
		user.PUT("/:userId", h.UpdateUser)
		user.DELETE("/:userId", h.DeleteUser)
	}


	courier := router.Group("api/couriers")
	{
		courier.POST("/create", h.CreateCourier)
		courier.GET("/:courierId", h.GetCourier)
		courier.PUT("/:courierId", h.UpdateCourier)
		courier.DELETE("/:courierId", h.DeleteCourier)
	}

	courierOrder := router.Group("api/courierOrders")
	{
		courierOrder.POST("/create", h.CreateCourierOrder)
		courierOrder.GET("/:orderId", h.GetCourierOrder)
		courierOrder.PUT("/:orderId", h.UpdateCourierOrder)
		courierOrder.DELETE("/:orderId", h.DeleteCourierOrder)
		courierOrder.GET("/", h.ListCourierOrders)
	}


	order := router.Group("api/orders")
	{
		order.POST("/create", h.CreateOrder)
		order.GET("/:orderId", h.GetOrder)
		order.GET("/", h.GetOrders)
		order.PUT("/:orderId", h.UpdateOrder)
		order.DELETE("/:orderId", h.DeleteOrder)

		order.GET("/recommended-orders", h.GetRecommendedOrders)
		order.PUT("/orders/:id/status", h.UpdateOrderStatus)
		order.GET("/orders/history", h.GetCourierOrderHistory)
	}


	orderProduct := router.Group("api/ordersProduct")
	{
		orderProduct.POST("/create", h.CreateOrderProduct)
		orderProduct.GET("/:orderId", h.GetOrderProduct)
		orderProduct.PUT("/:orderId", h.UpdateOrderProduct)
		orderProduct.DELETE("/:orderId", h.DeleteOrderProduct)
		orderProduct.GET("/", h.ListOrdersProduct)
	}

	task := router.Group("api/tasks")
	{
		task.POST("/create", h.CreateTask)
		task.GET("/:taskId", h.GetTask)
		task.GET("/", h.GetTasks)
		task.PUT("/:taskId", h.UpdateTask)
		task.DELETE("/:taskId", h.DeleteTask)
	}

	alert := router.Group("api/alerts")
	{
		alert.POST("/admin", h.CreateAdminAlert)
		alert.GET("/admin/:alertId", h.GetAdminAlert)
		alert.PUT("/admin/:alertId", h.UpdateAdminAlert)
		alert.DELETE("/admin/:alertId", h.DeleteAdminAlert)
	}

	notification := router.Group("api/notifications")
	{
		notification.POST("/courier", h.CreateCourierNotification)
		notification.POST("/user", h.CreateUserNotification)
		notification.GET("/courier/:notificationId", h.GetCourierNotification)
		notification.GET("/user/:notificationId", h.GetUserNotification)
	}

	cart := router.Group("api/carts")
	{
		cart.POST("/create", h.CreateCart)
		cart.GET("/:cartId", h.GetCart)
		cart.GET("/", h.GetAllCarts)
		cart.PUT("/:cartId", h.UpdateCart)
		cart.DELETE("/:cartId", h.DeleteCart)
	}

	product := router.Group("api/products")
	{
		product.POST("/create", h.CreateProduct)
		product.GET("/:productId", h.GetProduct)
		product.GET("/", h.GetProducts)
		product.PUT("/:productId", h.UpdateProduct)
		product.DELETE("/:productId", h.DeleteProduct)
	}


	return router
}
