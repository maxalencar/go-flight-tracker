package flighttracker

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"flight-tracker/pkg/model"
)

func TestFindFlightPath(t *testing.T) {
	service := newService()

	testCases := []struct {
		name     string
		given    []model.Flight
		expected any
	}{
		{
			name:     "TestMissingFlightRecords",
			given:    []model.Flight{},
			expected: errorFlightRecordsMustBeProvided,
		},
		{
			name:     "TestDirectFlight",
			given:    []model.Flight{{Source: "SFO", Destination: "EWR"}},
			expected: []string{"SFO", "EWR"},
		},
		{
			name: "TestMultiFlights",
			given: []model.Flight{
				{Source: "ATL", Destination: "EWR"},
				{Source: "SFO", Destination: "ATL"},
			},
			expected: []string{"SFO", "EWR"},
		},
		{
			name: "TestMultiFlights2",
			given: []model.Flight{
				{Source: "IND", Destination: "EWR"},
				{Source: "SFO", Destination: "ATL"},
				{Source: "GSO", Destination: "IND"},
				{Source: "ATL", Destination: "GSO"},
			},
			expected: []string{"SFO", "EWR"},
		},
		{
			name: "TestMultiFlights3",
			given: []model.Flight{
				{Source: "GSO", Destination: "IND"},
				{Source: "ATL", Destination: "GSO"},
				{Source: "DUB", Destination: "ATL"},
				{Source: "ATL", Destination: "GRU"},
				{Source: "IND", Destination: "EWR"},
				{Source: "GRU", Destination: "BSB"},
				{Source: "JFK", Destination: "DUB"},
				{Source: "BSB", Destination: "GRU"},
				{Source: "GRU", Destination: "JFK"},
				{Source: "SFO", Destination: "ATL"},
			},
			expected: []string{"SFO", "EWR"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := service.findFlightPath(tc.given)

			if _, ok := tc.expected.(error); ok {
				assert.Nil(t, result)
				assert.Equal(t, tc.expected, err)
				return
			}

			assert.Nil(t, err)
			assert.Equal(t, tc.expected, result)
		})

	}
}
