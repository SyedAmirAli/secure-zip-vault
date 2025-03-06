package services

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"

	"github.com/SyedAmirAli/secure-zip-vault/internal/config"
)

// DriveService handles interactions with Google Drive
type DriveService struct {
	service *drive.Service
	config  *config.Config
}

// NewDriveService creates a new Google Drive service
func NewDriveService(cfg *config.Config) (*DriveService, error) {
	ctx := context.Background()

	// Read the credentials file
	credBytes, err := os.ReadFile(cfg.GoogleCredentialsFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read Google credentials file: %w", err)
	}

	// Configure the Google Drive client
	config, err := google.JWTConfigFromJSON(credBytes, drive.DriveFileScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse Google credentials: %w", err)
	}

	// Create the Drive client
	client := config.Client(ctx)
	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("unable to create Drive service: %w", err)
	}

	return &DriveService{
		service: srv,
		config:  cfg,
	}, nil
}

// UploadBackup uploads a backup file to Google Drive
func (d *DriveService) UploadBackup(filePath string) (string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("unable to open file for upload: %w", err)
	}
	defer file.Close()

	// Create a new file on Google Drive
	driveFile := &drive.File{
		Name:     fmt.Sprintf("backup-%s.zip", time.Now().Format("2006-01-02-150405")),
		MimeType: "application/zip",
	}

	// Set parent folder if specified
	if d.config.GoogleDriveFolderID != "" {
		driveFile.Parents = []string{d.config.GoogleDriveFolderID}
	}

	// Upload the file
	res, err := d.service.Files.Create(driveFile).Media(file).Do()
	if err != nil {
		return "", fmt.Errorf("unable to upload file to Google Drive: %w", err)
	}

	return res.Id, nil
}

// ListBackups returns a list of backup files in the configured Google Drive folder
func (d *DriveService) ListBackups() ([]*drive.File, error) {
	var query string
	if d.config.GoogleDriveFolderID != "" {
		query = fmt.Sprintf("'%s' in parents and trashed = false", d.config.GoogleDriveFolderID)
	} else {
		query = "trashed = false"
	}

	// List files in the folder
	res, err := d.service.Files.List().
		Q(query).
		Fields("files(id, name, createdTime, size)").
		OrderBy("createdTime desc").
		Do()
	if err != nil {
		return nil, fmt.Errorf("unable to list files from Google Drive: %w", err)
	}

	return res.Files, nil
}

// DownloadBackup downloads a file from Google Drive by ID
func (d *DriveService) DownloadBackup(fileID, destPath string) error {
	// Get the file
	res, err := d.service.Files.Get(fileID).Download()
	if err != nil {
		return fmt.Errorf("unable to download file from Google Drive: %w", err)
	}
	defer res.Body.Close()

	// Create the destination file
	destFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("unable to create destination file: %w", err)
	}
	defer destFile.Close()

	// Copy the content
	_, err = io.Copy(destFile, res.Body)
	if err != nil {
		return fmt.Errorf("unable to save downloaded file: %w", err)
	}

	return nil
}

// DeleteBackup deletes a file from Google Drive by ID
func (d *DriveService) DeleteBackup(fileID string) error {
	err := d.service.Files.Delete(fileID).Do()
	if err != nil {
		return fmt.Errorf("unable to delete file from Google Drive: %w", err)
	}
	return nil
}

// ScheduleGDriveBackup sets up automatic daily backups to Google Drive
func ScheduleGDriveBackup(cfg *config.Config) {
	// This function would typically set up a goroutine with a ticker
	// to perform backups at regular intervals
	go func() {
		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()

		for range ticker.C {
			// Create a backup
			backupPath, err := CreateBackup(cfg)
			if err != nil {
				fmt.Printf("Error creating backup: %v\n", err)
				continue
			}

			// Upload to Google Drive
			driveService, err := NewDriveService(cfg)
			if err != nil {
				fmt.Printf("Error initializing Google Drive service: %v\n", err)
				continue
			}

			fileID, err := driveService.UploadBackup(backupPath)
			if err != nil {
				fmt.Printf("Error uploading backup to Google Drive: %v\n", err)
				continue
			}

			fmt.Printf("Backup successfully uploaded to Google Drive with ID: %s\n", fileID)

			// Clean up the local backup file
			os.Remove(backupPath)
		}
	}()
}
