package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RightHandSide/CVWO-App/internal/handlers/content"
	"github.com/RightHandSide/CVWO-App/internal/handlers/users"
	"github.com/go-chi/chi/v5"
)

func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		// USER
		r.Get("/users", func(w http.ResponseWriter, req *http.Request) {
			response, _ := users.HandleUser(w, req)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})

		r.Post("/register", func(w http.ResponseWriter, req *http.Request) {
			response, _ := users.HandleRegister(w, req)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})

		// GET CONTENT
		r.Get("/threads", func(w http.ResponseWriter, req *http.Request) {
			response, _ := content.HandleThread(w, req)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})

		// r.Get("/thread-{id}", func(w http.ResponseWriter, req *http.Request) {
		// 	response, _ = content.ListPost(w, req)

		// 	w.Header().Set("Content-Type", "application/json")
		// 	json.NewEncoder(w).Encode(response)
		// })

		// r.Get("/post-{id}", func(w http.ResponseWriter, req *http.Request) {
		// 	response, _ = content.ListComment(w, req)

		// 	w.Header().Set("Content-Type", "application/json")
		// 	json.NewEncoder(w).Encode(response)
		// })
	}
}

// Different Routes
