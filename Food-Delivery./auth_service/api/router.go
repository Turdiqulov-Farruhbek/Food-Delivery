package api

import (
	"delivery/auth_service/api/handler"
	"delivery/auth_service/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "delivery/auth_service/docs"
)

// @Title Online Forum API Documentation
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// TODO kamroq middleware ishlating. Iloji boricha umumiy middleware ishlatishga harakat qiling
func NewGin(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	user := r.Group("/user")
	user.Use(middleware.Middleware())
	{
		user.POST("/register", h.RegisterUser)
		user.POST("/login", h.LoginUser)
		user.PUT("/profile", h.UpdateProfile)
		user.PUT("/password", h.ChangePassword)
		user.POST("/forgot_password", h.ForgotPassword)
		user.POST("/reset_password", h.ResetPassword)
		user.GET("/profile", h.GetUserProfile)
		user.POST("/confirm", h.ConfirmEmail)
		user.POST("/resend", h.ResendCode)
	}

	return r
}
