package config

import (
	"os"
	
	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	// Server configuration
	ServerPort string

	// Authentication
	JWTSecret     string
	AdminPassword string // Hardcoded password for authentication

	// Backup configuration
	ProjectPath   string
	DatabaseName  string
	DatabaseUser  string
	DatabasePass  string
	BackupTempDir string

	// Google Drive configuration
	GoogleCredentialsFile string
	GoogleDriveFolderID  string
}

// Load returns a Config struct with values from environment variables or defaults
func Load() (*Config, error) {
	// Load .env file if it exists
	godotenv.Load()
	
	return &Config{
		ServerPort:           getEnv("SERVER_PORT", "8080"),
		JWTSecret:            getEnv("JWT_SECRET", "your-super-secret-jwt-key"),
		AdminPassword:        getEnv("ADMIN_PASSWORD", "SecureZipVault2023!"), // Now from env
		ProjectPath:          getEnv("PROJECT_PATH", "/var/www/myproject"),
		DatabaseName:         getEnv("DB_NAME", "myproject"),
		DatabaseUser:         getEnv("DB_USER", "dbuser"),
		DatabasePass:         getEnv("DB_PASS", "dbpassword"),
		BackupTempDir:        getEnv("BACKUP_TEMP_DIR", "/tmp/secure-zip-vault"),
		GoogleCredentialsFile: getEnv("GOOGLE_CREDS", "credentials.json"),
		GoogleDriveFolderID:  getEnv("GDRIVE_FOLDER_ID", ""),
	}, nil
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
