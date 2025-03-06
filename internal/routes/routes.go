package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/SyedAmirAli/secure-zip-vault/internal/config"
	"github.com/SyedAmirAli/secure-zip-vault/internal/handlers"
)

// SetupRouter configures the Gin router with all routes
func SetupRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	// Serve static files for the frontend
	r.Static("/static", "./web/static")
	r.LoadHTMLGlob("web/templates/*")

	// Home page
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "SecureZipVault",
		})
	})

	// API routes
	api := r.Group("/api")
	{
		// Authentication
		api.POST("/auth/login", handlers.Login(cfg))

		// Protected routes
		protected := api.Group("/")
		protected.Use(handlers.AuthMiddleware(cfg))
		{
			// Download
			protected.GET("/download", handlers.DownloadProjectBackup(cfg))
			protected.GET("/backup/status", handlers.GetBackupStatus(cfg))
		}
	}

	return r
}
