package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// CreateZipArchive creates a ZIP archive containing the specified paths
func CreateZipArchive(zipPath string, sourcePaths []string) error {
	// Create the ZIP file
	zipFile, err := os.Create(zipPath)
	if err != nil {
		return fmt.Errorf("failed to create ZIP file: %w", err)
	}
	defer zipFile.Close()

	// Create a new ZIP writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Add each source path to the ZIP
	for _, sourcePath := range sourcePaths {
		// Walk through all files in the source path
		err := filepath.Walk(sourcePath, func(filePath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Skip directories themselves (we'll create them implicitly)
			if info.IsDir() {
				return nil
			}

			// Create a ZIP header
			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return fmt.Errorf("failed to create ZIP header: %w", err)
			}

			// Set the name to be relative to the source path
			relPath, err := filepath.Rel(filepath.Dir(sourcePath), filePath)
			if err != nil {
				return fmt.Errorf("failed to get relative path: %w", err)
			}
			header.Name = filepath.ToSlash(relPath)

			// Use deflate compression
			header.Method = zip.Deflate

			// Create the file in the ZIP
			writer, err := zipWriter.CreateHeader(header)
			if err != nil {
				return fmt.Errorf("failed to create file in ZIP: %w", err)
			}

			// Open the source file
			file, err := os.Open(filePath)
			if err != nil {
				return fmt.Errorf("failed to open source file: %w", err)
			}
			defer file.Close()

			// Copy the file contents to the ZIP
			_, err = io.Copy(writer, file)
			if err != nil {
				return fmt.Errorf("failed to copy file to ZIP: %w", err)
			}

			return nil
		})

		if err != nil {
			return fmt.Errorf("failed to add %s to ZIP: %w", sourcePath, err)
		}
	}

	return nil
}
