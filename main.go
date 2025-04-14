package main

import (
	"github.com/nihilKnight/pbl-platform-backend-gin/routes"

	"github.com/nihilKnight/pbl-platform-backend-gin/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	
	r := gin.Default()
	
	api := r.Group("/api/v2")
	{
		routes.UserRoutes(api)
		// routes.CourseRoutes(api)
		// routes.UserCourseRoutes(api)
	}
	
	r.Run(":8080")
}