# Flight Tracker

Flight Tracker microservice used to query customer's flight path.

## Getting Set Up

Before running the application, you will need to ensure that you have a few requirements installed;
You will need Go.

### Go
### Docker (optional)

[Go](https://golang.org/) is an open source programming language that makes it easy to build simple, reliable, and efficient software.

[Docker](https://docker.com/) is used to build and sharing containerized applications.

## Project Structure

Following [Standard Go Project Layout](https://github.com/golang-standards/project-layout), apologies in advance if it is too much for the purpose of the assignment.

### `/cmd`

Main application for this project.

### `/internal`

Internal application logic

### `/pkg`

Model that is okay to be shared with external applications.

## Running the server

    go run cmd/flighttracker/main.go

### Running using docker-compose

The docker-compose file that resides under the root folder. To run it just run the following command under the root folder: `docker-compose up --build`. (update the port if you are not running in default port 8080)

### Test

[Table-driven tests using subtests](https://blog.golang.org/subtests) were used as the approach to reduce the amount of repetitive code compared to repeating the same code for each test and makes it straightforward to add more test cases.

[Testify](https://github.com/stretchr/testify) the [assert](https://github.com/stretchr/testify#assert-package) package assert provides a set of comprehensive testing tools for use with the normal Go testing system.

## Running the tests

    go test ./... -v

### Endpoints

#### POST /find

Find Customer's Flight Path

Accepts JSON input in the format:

`[{"source": "source", "destination": "destination"}]`

and it returns an array containing the customer's source and destination based on the flight records.

Example:

Request:

    curl -X POST http://localhost:8080/find \
    -H 'Content-Type: application/json' \
    -d '[{"source": "IND","destination": "EWR"},{"source": "SFO","destination": "ATL"},{"source": "GSO","destination": "IND"},{"source": "ATL","destination": "GSO"}]'

Response:
    
    ["SFO", "EWR"]

[@maxalencar](https://github.com/maxalencar)
