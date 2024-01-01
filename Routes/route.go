package Routes

import (
	"github.com/arun-kushwaha04/ecommerce-backend/Controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/user/signup", Controllers.SignUp())
	incomingRoutes.POST("/user/login", Controllers.Login())
	incomingRoutes.POST("/admin/addProduct", Controllers.AddProduct())
	incomingRoutes.GET("/user/viewProduct", Controllers.ViewProduct())
	incomingRoutes.GET("/user/search", Controllers.SearchProduct())
}
