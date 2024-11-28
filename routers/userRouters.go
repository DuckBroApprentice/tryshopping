package routers

import (
	"github.com/DuckBroApprentice/Shopping/controllers"
	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users")

	user.GET("/", controllers.FindAllUsers)
	user.POST("/", controllers.PostUser)
	user.DELETE("/:id", controllers.DeleteUser)
	user.PUT("/:id", controllers.PutUser)
}
