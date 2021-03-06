package routes

import (
	"TeachAssistApi/app/controllers"
	"github.com/gin-gonic/gin"
)

func AddUserRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	users.POST("/login", controllers.LoginUser())
	users.GET("/renew_session", controllers.RenewUserSession())
	users.DELETE("/remove", controllers.RemoveUser())
}
