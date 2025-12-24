package router

import (
	"github.com/RightHandSide/CVWO-App/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func Setup() chi.Router {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}))

	setUpRoutes(r)
	return r
}

func setUpRoutes(r chi.Router) {
	r.Group(routes.GetRoutes())
}
