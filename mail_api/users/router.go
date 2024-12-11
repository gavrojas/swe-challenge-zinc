package users

import (
	"mail_indexer_zinc/shared"

	"github.com/go-chi/chi/v5"
)

func AddUserRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		// Aplicar middleware de autenticación al grupo
		r.Use(shared.AuthenticateSession)

		r.Get("/", GetAllUsers)
		r.Post("/", CreateUser)
		r.Get("/{id}", GetUserById)
		r.Delete("/{id}", DeleteUser)
		r.Put("/{id}", UpdateUser)
	})
}
