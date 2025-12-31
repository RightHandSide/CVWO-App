package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RightHandSide/CVWO-App/internal/api"
	"github.com/RightHandSide/CVWO-App/internal/handlers/content"
	"github.com/RightHandSide/CVWO-App/internal/handlers/users"
	"github.com/go-chi/chi/v5"
)

func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		// USER
		r.Get("/users", helper(users.HandleUser))
		r.Post("/register", helper(users.HandleRegister))
		r.Post("/login", helper(users.HandleLogin))

		// CONTENT
		r.Get("/threads", helper(content.HandleThread))
		r.Post("/createthreads", helper(content.CreateThread))

		r.Get("/thread/{id}", helper(content.HandlePost))
		r.Post("/createposts/{id}", helper(content.CreatePost))

		r.Get("/post/{id}", helper(content.HandleComment))
		r.Post("/createcomments/{id}", helper(content.CreateComment))
	}
}

type HandleFunc func(w http.ResponseWriter, req *http.Request) (*api.Response, error)

func helper(h HandleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		response, err := h(w, req)
		if err != nil {
			fmt.Println(`ERROR: ?`, err)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// Different Routes
