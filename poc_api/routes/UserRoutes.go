package routes

import (
	"github.com/gin-gonic/gin"
	"poc_api/poc_api/controllers"
)

func RegisterUserRoutes(r *gin.Engine, controller *controllers.UserController) {
	user := r.Group("/User")
	{
		user.POST("/register_user", controller.Register)
		user.GET("/find_all_users", controller.FindAll)
	}
}
