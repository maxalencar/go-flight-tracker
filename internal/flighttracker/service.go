package flighttracker

import (
	"errors"

	"flight-tracker/pkg/model"
)

var errorFlightRecordsMustBeProvided = errors.New("a list of flight records must be provided")

type service struct{}

func newService() *service {
	return &service{}
}

// findFlightPath finds the source and destination based on the flight records (assuming it's an one way flight)
func (s *service) findFlightPath(records []model.Flight) ([]string, error) {
	if len(records) == 0 {
		return nil, errorFlightRecordsMustBeProvided
	}

	if len(records) == 1 {
		return []string{records[0].Source, records[0].Destination}, nil
	}

	// create sources and destinations maps
	var sourceMap = make(map[string]int, 0)
	var destinationMap = make(map[string]int, 0)

	// populate maps
	for _, record := range records {
		sourceMap[record.Source]++
		destinationMap[record.Destination]++
	}

	// find source and destination
	var source, destination string

	for _, record := range records {
		if _, ok := destinationMap[record.Source]; !ok {
			source = record.Source
		}

		if _, ok := sourceMap[record.Destination]; !ok {
			destination = record.Destination
		}

		if source != "" && destination != "" {
			break
		}
	}

	return []string{source, destination}, nil
}
