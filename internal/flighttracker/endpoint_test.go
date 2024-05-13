package flighttracker

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"flight-tracker/pkg/model"
)

func TestFind(t *testing.T) {
	service := newService()
	endpoint := newEndpoint(service)
	ei := endpoint.init()

	testCases := []struct {
		name     string
		given    []model.Flight
		expected int
	}{
		{
			name:     "TestMissingFlightRecords",
			given:    []model.Flight{},
			expected: http.StatusInternalServerError,
		},
		{
			name:     "TestSuccess",
			given:    []model.Flight{{Source: "SFO", Destination: "EWR"}},
			expected: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b, err := json.Marshal(tc.given)
			assert.NoError(t, err)

			reqBody := bytes.NewReader(b)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/find", reqBody)
			r.Header.Add("Content-Type", "application/json")

			ei.ServeHTTP(w, r)
			assert.Equal(t, tc.expected, w.Code)
		})
	}
}
