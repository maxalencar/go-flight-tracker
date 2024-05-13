package flighttracker

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func (e *endpoint) init() *http.ServeMux {
	// Initialize HTTP request multiplexer
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("POST /find", e.find)

	return mux
}

func (e endpoint) find(w http.ResponseWriter, r *http.Request) {
	flights := make([]model.Flight, 0)
	if err := json.NewDecoder(r.Body).Decode(&flights); err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	result, err := e.service.findFlightPath(flights)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to find flight path: %v", err), http.StatusInternalServerError)
		return
	}

	// Encode object as JSON and write to response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}
}
