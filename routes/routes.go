package routes

import (
	"backend-porto/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Serve static files for uploaded images
	r.Static("/uploads", "./public/uploads")

	api := r.Group("/api")
	{
		// Dashboard & Profile
		api.GET("/dashboard/stats", controllers.GetDashboardStats)
		api.POST("/profile", controllers.CreateProfile)
		api.PUT("/profile", controllers.UpdateProfile)

		// Projects
		api.GET("/projects", controllers.GetProjects)
		api.GET("/projects/:id", controllers.GetProject)
		api.POST("/projects", controllers.CreateProject)
		api.PUT("/projects/:id", controllers.UpdateProject)
		api.DELETE("/projects/:id", controllers.DeleteProject)

		// Professional Journey / Experiences
		api.GET("/experiences", controllers.GetExperiences)
		api.POST("/experiences", controllers.CreateExperience)
		api.PUT("/experiences/:id", controllers.UpdateExperience)
		api.DELETE("/experiences/:id", controllers.DeleteExperience)

		// Skills
		api.GET("/skills", controllers.GetSkills)
		api.POST("/skills", controllers.CreateSkill)
		api.PUT("/skills/:id", controllers.UpdateSkill)
		api.DELETE("/skills/:id", controllers.DeleteSkill)
	}
}
