package mails

import (
	"fmt"
	"index_data_zinc/config"
	"index_data_zinc/utils"
	"index_data_zinc/zinc"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func ProcessEmails(mailDir string) error {
	var wg sync.WaitGroup
	totalEmails := 0
	var batch []zinc.EmailDocument

	files := make(chan []zinc.EmailDocument, 1000)

	// Lanzar trabajadores
	for i := 0; i < config.NumWorkers; i++ {
		wg.Add(1)
		go ProcessFiles(i, files, &wg)
	}

	// Explorar el directorio y pasar los archivos al canal
	go func() {
		err := filepath.Walk(mailDir, func(path string, info os.FileInfo, err error) error {
			path = filepath.ToSlash(path)
			if err != nil {
				// Log the specific error but continue walking
				log.Printf("Error accessing path %q: %v\n", path, err)
				return err
			}

			if info.IsDir() || !utils.IsEmailFile(path) {
				return nil
			}

			email, parseErr := utils.ParseEmailFile(path)

			if parseErr != nil {
				logFile, _ := os.OpenFile("failed_files.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				defer logFile.Close()
				log.SetOutput(logFile)
				log.Printf("Error parsing email file %s: %v", path, parseErr)
				return nil
			}

			// Extract user and folder information
			relPath, _ := filepath.Rel(mailDir, path)
			pathParts := strings.Split(relPath, string(filepath.Separator))

			if len(pathParts) >= 2 {
				email.User = pathParts[0]
				email.FolderPath = strings.Join(pathParts[1:len(pathParts)-1], "/")
			}

			batch = append(batch, *email)
			totalEmails++

			// When batch reaches specified size, send to Zinc
			if len(batch) > 0 {
				files <- batch
				batch = nil // Reiniciar el lote
			}
			return nil
		})
		if err != nil {
			log.Fatalf("Failed to walk the directory: %v", err)
		}

		close(files)
	}()

	wg.Wait()

	fmt.Printf("Total emails processed: %d\n", totalEmails)
	return nil
}

func ProcessFiles(id int, files <-chan []zinc.EmailDocument, wg *sync.WaitGroup) {
	defer wg.Done()
	for file := range files {
		log.Printf("Worker %d processing batch of size %d\n", id, len(file))
		if err := zinc.SendBulkToZinc(file); err != nil {
			log.Printf("Worker %d encountered error: %v\n", id, err)
		}
	}
}

// ProcessEmailsWithoutRoutines processes emails sequentially without using goroutines
func ProcessEmailsWithoutRoutines(mailDir string) error {
	totalEmails := 0
	var batch []zinc.EmailDocument

	err := filepath.Walk(mailDir, func(path string, info os.FileInfo, err error) error {
		path = filepath.ToSlash(path)
		if err != nil {
			log.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}

		if info.IsDir() || !utils.IsEmailFile(path) {
			return nil
		}

		email, parseErr := utils.ParseEmailFile(path)

		if parseErr != nil {
			logFile, _ := os.OpenFile("failed_files.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			defer logFile.Close()
			log.SetOutput(logFile)
			log.Printf("Error parsing email file %s: %v", path, parseErr)
			return nil
		}

		// Extract user and folder information
		relPath, _ := filepath.Rel(mailDir, path)
		pathParts := strings.Split(relPath, string(filepath.Separator))

		if len(pathParts) >= 2 {
			email.User = pathParts[0]
			email.FolderPath = strings.Join(pathParts[1:len(pathParts)-1], "/")
		}

		batch = append(batch, *email)
		totalEmails++

		// When batch reaches specified size, send to Zinc
		if len(batch) >= config.BatchSize {
			if err := zinc.SendBulkToZinc(batch); err != nil {
				log.Printf("Error sending batch: %v", err)
			}
			batch = nil // Reset batch
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Failed to walk the directory: %v", err)
	}

	// Send any remaining emails in the batch
	if len(batch) > 0 {
		if err := zinc.SendBulkToZinc(batch); err != nil {
			log.Printf("Error sending final batch: %v", err)
		}
	}

	fmt.Printf("Indexing complete! Total emails processed: %d\n", totalEmails)
	return nil
}
