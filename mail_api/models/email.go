package models

import (
	"time"
)

// basic struct for an email
type Email struct {
	MessageID               string    `json:"message_id"`                // ID único del mensaje
	Date                    time.Time `json:"date"`                      // Fecha de envío del correo
	From                    string    `json:"from"`                      // Dirección de quien envió el correo
	To                      string    `json:"to"`                        // Dirección de destino del correo
	Cc                      string    `json:"cc"`                        // Dirección de copia oculta
	Bcc                     string    `json:"bcc"`                       // Dirección de copia oculta
	Subject                 string    `json:"subject"`                   // Asunto del correo
	MimeVersion             string    `json:"mime_version"`              // Versión MIME
	ContentType             string    `json:"content_type"`              // Tipo de contenido
	ContentTransferEncoding string    `json:"content_transfer_encoding"` // Codificación del contenido
	Folder                  string    `json:"folder"`                    // Carpeta donde está almacenado el correo
	Origin                  string    `json:"origin"`                    // Origen del mensaje
	FileName                string    `json:"file_name"`                 // Nombre del archivo
	Body                    string    `json:"body"`                      // Cuerpo del correo (contenido)
}
