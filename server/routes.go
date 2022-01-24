package server

import (
	"docker-example/handler"
	"github.com/go-chi/chi/v5"

	"net/http"
)

type Server struct {
	chi.Router
}

func SetupRoutes() *Server {
	router := chi.NewRouter()
	router.Route("/api", func(api chi.Router) {
		api.Get("/welcome", handler.Welcome)
	})
	return &Server{router}
}

func (svc *Server) Run(port string) error {
	return http.ListenAndServe(port, svc)
}
