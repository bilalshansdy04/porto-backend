package api

import (
	"backend-porto/models"
	"backend-porto/routes"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func init() {
	// Connect to Database
	models.ConnectDatabase()

	// Use release mode for Vercel
	gin.SetMode(gin.ReleaseMode)
	app = gin.Default()

	// CORS configuration for the frontend
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"}, // Common React dev server ports
		AllowOriginFunc:  func(origin string) bool { return true },                   // Allow all origins for simplicity in Vercel
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	app.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong from vercel",
		})
	})

	// Setup API routes
	routes.SetupRoutes(app)
}

// Handler is the entrypoint for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
