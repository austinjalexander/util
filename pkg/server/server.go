package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

const (
	timeout = time.Second * 20
)

// Handler represents a HTTP handler for the server router.
type Handler struct {
	Func                          func(http.ResponseWriter, *http.Request)
	Path                          string
	Headers, Methods, QueryParams []string
}

// JSONresponse reprents a JSON response body returned by the server.
type JSONresponse struct {
	Data   interface{} `json:"data,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}

// Server represents modifiable server configuration.
type Server struct {
	Handlers          []Handler
	OnlyJSONresponses bool
	Port              uint16
}

// New returns a new server configuration.
func New() Server {
	return Server{}
}

// Run creates a new routed server and runs it.
func (s Server) Run() {
	r := mux.NewRouter()
	s.configureMiddleware(r)
	s.registerRoutes(r)

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", s.Port),
		Handler:      r,
		IdleTimeout:  timeout,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	srv.Shutdown(ctx)
	log.Println("info: gracefully shutting down server...")
	os.Exit(0)
}

func (s Server) configureMiddleware(r *mux.Router) {
	if s.OnlyJSONresponses {
		r.Use(onlyJSONresponses)
	}
}

func (s Server) registerRoutes(r *mux.Router) {
	for _, h := range s.Handlers {
		r.HandleFunc(h.Path, h.Func).
			Headers(h.Headers...).
			Methods(h.Methods...).
			Queries(h.QueryParams...)
	}
}

// Middleware.
func onlyJSONresponses(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
