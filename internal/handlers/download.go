package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/SyedAmirAli/secure-zip-vault/internal/config"
	"github.com/SyedAmirAli/secure-zip-vault/internal/services"
)

// DownloadProjectBackup handles the request to download the project backup
func DownloadProjectBackup(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a backup of the project and database
		backupPath, err := services.CreateBackup(cfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create backup"})
			return
		}

		// Set the appropriate headers for file download
		filename := filepath.Base(backupPath)
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Header("Content-Type", "application/zip")
		c.File(backupPath)

		// The file will be deleted after serving in a production environment
		// This would be handled by a cleanup routine
	}
}

// GetBackupStatus returns the current status of the backup process
func GetBackupStatus(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// In a real implementation, you would track the backup status
		// For now, we'll just return a simple status
		c.JSON(http.StatusOK, gin.H{
			"status": "in_progress",
			"progress": 75,
			"message": "Creating ZIP archive...",
		})
	}
}
