package auth

import (
	"github.com/go-chi/chi/v5"
)

func AddAuthRoutes(r chi.Router) {
	// Define the /auth group
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", Register)
		r.Post("/login", Login)
		r.Post("/logout", Logout)
	})
}
