package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nihilKnight/pbl-platform-backend-gin/controllers"
)

func UserRoutes(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("/add", controllers.CreateUser)
		// user.DELETE("/delete/:id", controllers.DeleteUser)
		// user.PUT("/update/:id", controllers.UpdateUser)
		// user.GET("/all", controllers.GetAllUsers)
		// user.GET("/:id", controllers.GetUserByID)
	}
}