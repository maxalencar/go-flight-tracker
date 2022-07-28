FROM golang:1.18-alpine AS builder

WORKDIR /go/src/app
COPY . .

RUN /usr/local/go/bin/go build -o flighttracker ./cmd/flighttracker/

FROM alpine

COPY --from=builder /go/src/app/flighttracker /usr/local/bin/flighttracker
WORKDIR /usr/local/bin

CMD [ "flighttracker" ]