FROM golang
WORKDIR /go/src/flight
COPY go.mod go.sum ./
RUN go mod download
COPY services/flight services/flight
COPY services/flight/.env .env
RUN go build -o /go/bin/app services/flight/cmd/main.go
CMD ["app"]
