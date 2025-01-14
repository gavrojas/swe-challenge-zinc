package main

import (
	"fmt"
	"log"
	"mail_indexer_zinc/auth"
	"mail_indexer_zinc/emails"
	"mail_indexer_zinc/shared"
	"mail_indexer_zinc/users"
	"mail_indexer_zinc/zinc"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Crear el índice de usuarios si no existe
	if err := zinc.CreateIndex("users"); err != nil {
		log.Fatalf("Error creating index: %s", err)
		return
	}

	// Instanciar el router con chi
	r := chi.NewRouter()

	r.Use(shared.Cors)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// Rutas
	users.AddUserRoutes(r)
	auth.AddAuthRoutes(r)
	emails.AddEmailsRoutes(r)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server running"))
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Puerto en el que escucha el servidor
	port := ":443"
	certFile := "/etc/ssl/certs/fullchain.pem"
	keyFile := "/etc/ssl/private/privkey.pem"
	fmt.Printf("Server running on %s\n", port)

	if err := http.ListenAndServeTLS(port, certFile, keyFile, r); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}

	http.ListenAndServe(":443", r)
}
