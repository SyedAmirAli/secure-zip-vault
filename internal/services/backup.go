package services

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/SyedAmirAli/secure-zip-vault/internal/config"
	"github.com/SyedAmirAli/secure-zip-vault/internal/utils"
)

// CreateBackup creates a backup of the project files and database
func CreateBackup(cfg *config.Config) (string, error) {
	// Create temp directory if it doesn't exist
	if err := os.MkdirAll(cfg.BackupTempDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create temp directory: %w", err)
	}

	// Generate a timestamp for the backup filename
	timestamp := time.Now().Format("20060102-150405")
	backupFilename := fmt.Sprintf("project-backup-%s.zip", timestamp)
	backupPath := filepath.Join(cfg.BackupTempDir, backupFilename)

	// Create a temporary directory for the database dump
	dbDumpDir := filepath.Join(cfg.BackupTempDir, "db-dump-"+timestamp)
	if err := os.MkdirAll(dbDumpDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create database dump directory: %w", err)
	}
	defer os.RemoveAll(dbDumpDir) // Clean up the dump directory when done

	// Dump the MySQL database
	dbDumpFile := filepath.Join(dbDumpDir, "database.sql")
	if err := dumpDatabase(cfg, dbDumpFile); err != nil {
		return "", fmt.Errorf("failed to dump database: %w", err)
	}

	// Create a ZIP file containing the project files and database dump
	if err := utils.CreateZipArchive(backupPath, []string{cfg.ProjectPath, dbDumpDir}); err != nil {
		return "", fmt.Errorf("failed to create ZIP archive: %w", err)
	}

	// Encrypt the ZIP file (in a real implementation)
	// For now, we'll skip this step

	return backupPath, nil
}

// dumpDatabase dumps the MySQL database to a file
func dumpDatabase(cfg *config.Config, outputFile string) error {
	cmd := exec.Command(
		"mysqldump",
		"--user="+cfg.DatabaseUser,
		"--password="+cfg.DatabasePass,
		cfg.DatabaseName,
	)

	// Open the output file
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Set the output to the file
	cmd.Stdout = file

	// Run the command
	return cmd.Run()
}

// ScheduleDailyBackup sets up a daily backup to Google Drive
// This would typically be called when the application starts
func ScheduleDailyBackup(cfg *config.Config) {
	// In a real implementation, you would set up a cron job or use a scheduler
	// For this example, we'll just log that it would be scheduled
	fmt.Println("Daily backup to Google Drive would be scheduled here")
}
