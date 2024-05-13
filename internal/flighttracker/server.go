package flighttracker

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init() {
	service := newService()
	endpoint := newEndpoint(service)

	s.mux = endpoint.init()
}

func (s *Server) Start(port string) error {
	// Start HTTP server
	addr := fmt.Sprintf(":%s", port)
	log.Printf("server listening on port %s...\n", port)

	server := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: 10 * time.Second,
		Handler:           s.mux,
	}

	return server.ListenAndServe()
}
