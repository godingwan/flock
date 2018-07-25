package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jsungholee/flock"
)

// Server has fields useful for starting up a web server
type Server struct {
	addr    string
	port    uint
	handler http.Handler
}

// NewServer returns a new Server struct given an address and a port to listen on
func NewServer(addr string, port uint, handler http.Handler) Server {
	return Server{
		addr:    addr,
		port:    port,
		handler: handler,
	}
}

// Start fires up the server to start responding to requests.
func (s Server) Start() error {
	return http.ListenAndServe(fmt.Sprintf("%s:%d", s.addr, s.port), s.handler)
}

// BuildRouter creates all of the routes for the api and attaches handlers to them
func BuildRouter(svc flock.Service) *chi.Mux {
	r := chi.NewMux()

	r.Get("/health", handleHealth)
	r.Method(http.MethodGet, "/tweets", HandleSearch(svc))

	return r
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("I'm tweeting"))
}
