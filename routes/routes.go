package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kamogelosekhukhune777/ecom-cart/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("user/signup", controllers.SignUp())
	incomingRoutes.POST("users/login", controllers.Login())
	incomingRoutes.POST("admin/addproduct", controllers.ProductViewAdmin())
	incomingRoutes.GET("users/product", controllers.SearchProduct())
	incomingRoutes.GET("users/search", controllers.SearchProductByQuery())
}
