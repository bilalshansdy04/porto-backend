package api

import (
	"backend-porto/models"
	"backend-porto/routes"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var app *gin.Engine
var initErr error

func init() {
	// Let's capture the env vars to see if they are empty
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	
	if host == "" || port == "" {
		// We won't crash, we'll store the error to display it in the browser
		initErr = os.ErrNotExist // Just a marker
	} else {
		// Connect to Database normally
		models.ConnectDatabase()
	}

	// Use release mode for Vercel
	gin.SetMode(gin.ReleaseMode)
	app = gin.Default()

	// CORS configuration for the frontend
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"}, 
		AllowOriginFunc:  func(origin string) bool { return true },
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
	if initErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Environment variables DB_HOST or DB_PORT are empty. Vercel is not injecting them.", "host": "` + os.Getenv("DB_HOST") + `", "port": "` + os.Getenv("DB_PORT") + `"}`))
		return
	}
	app.ServeHTTP(w, r)
}
