package flighttracker

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo *echo.Echo
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init() {
	service := newService()
	endpoint := newEndpoint(service)

	s.Echo = endpoint.init()
}

func (s *Server) Start(port string) error {
	return s.Echo.Start(fmt.Sprintf(":%s", port))
}
