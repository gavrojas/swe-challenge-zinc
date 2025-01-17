package utils

import (
	"fmt"
	"index_data_zinc/zinc"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Verifica si el archivo es un correo electrónico válido (por ejemplo, con extensión ".")
func IsEmailFile(path string) bool {
	path = filepath.ToSlash(path) /* convertir \ y / */
	fileName := filepath.Base(path)

	// Condición simple: archivos que terminan con un punto
	isEmail := strings.HasSuffix(fileName, ".") && fileName != "." && fileName != ".." /*<- Omitir directorios actuales y directorio atrás, archivos terminados con . */
	return isEmail
}

func ParseEmailFile(filePath string) (*zinc.EmailDocument, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %v", filePath, err)
	}
	defer file.Close()
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", filePath, err)
	}

	email := &zinc.EmailDocument{}
	content := string(fileBytes)
	lines := strings.Split(content, "\n")

	var bodyLines []string
	var inBody bool

	for _, line := range lines {
		if line == "" {
			inBody = true
			continue
		}

		if !inBody {
			parts := strings.SplitN(line, ": ", 2)
			if len(parts) == 2 {
				switch parts[0] {
				case "Message-ID":
					email.MessageID = parts[1]
				case "Date":
					email.Date = parts[1]
				case "From":
					email.From = parts[1]
				case "To":
					email.To = parts[1]
				case "Subject":
					email.Subject = parts[1]
				case "Mime-Version":
					email.MimeVersion = parts[1]
				case "Content-Type":
					email.ContentType = parts[1]
				case "Content-Transfer-Encoding":
					email.ContentTransferEncoding = parts[1]
				case "X-From":
					email.XFrom = parts[1]
				case "X-To":
					email.XTo = parts[1]
				case "X-cc":
					email.XCc = parts[1]
				case "X-bcc":
					email.XBcc = parts[1]
				case "X-Folder":
					email.XFolder = parts[1]
				case "X-Origin":
					email.XOrigin = parts[1]
				case "X-FileName":
					email.XFileName = parts[1]
				}
			}
		} else {
			bodyLines = append(bodyLines, line)
		}
	}

	email.Body = strings.Join(bodyLines, "\n")
	return email, nil
}
