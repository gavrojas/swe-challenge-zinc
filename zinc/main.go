package main

import (
	"fmt"
	"index_data_zinc/config"
	"index_data_zinc/mails"
	"index_data_zinc/zinc"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Verificar que se haya pasado un argumento
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <data_directory>\n", os.Args[0])
	}

	// Obtener el directorio de datos desde los argumentos
	dataDir := os.Args[1]
	mailDir := filepath.Join(dataDir, "/maildir")
	fmt.Printf("mailDir: %s\n", mailDir)

	// Crear índice si no existe
	err := zinc.CreateZincIndex(config.ZincIndex)
	if err != nil {
		log.Fatalf("Failed to create index: %v", err)
	}

	// Verificar si el directorio existe
	if _, err := os.Stat(mailDir); os.IsNotExist(err) {
		log.Fatalf("Directory does not exist: %v", err)
	}

	// Llamar a la función para procesar correos electrónicos
	if err := mails.ProcessEmails(mailDir); err != nil {
		log.Fatalf("Error processing emails: %v", err)
	}
}
