package emails

import (
	"mail_indexer_zinc/shared"

	"github.com/go-chi/chi/v5"
)

func AddEmailsRoutes(r chi.Router) {
	r.Route("/emails", func(r chi.Router) {
		// Aplicar middleware de autenticación al grupo
		r.Use(shared.AuthenticateSession)

		r.Post("/ByField", SearchMailsByField)
	})
}
