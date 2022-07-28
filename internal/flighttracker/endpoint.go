package flighttracker

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"flight-tracker/pkg/model"
)

type endpoint struct {
	service *service
}

func newEndpoint(service *service) *endpoint {
	return &endpoint{
		service: service,
	}
}

func (e *endpoint) init() *echo.Echo {
	// Echo instance
	ei := echo.New()

	// Middleware
	ei.Use(middleware.Logger())
	ei.Use(middleware.Recover())

	// Routes
	ei.POST("/find", e.find)

	return ei
}

func (e endpoint) find(c echo.Context) error {
	flights := make([]model.Flight, 0)
	if err := c.Bind(&flights); err != nil {
		return err
	}

	result, err := e.service.findFlightPath(flights)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}
