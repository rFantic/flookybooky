FROM golang
WORKDIR /go/src/flight
RUN go install github.com/go-delve/delve/cmd/dlv@v1.20.0
COPY go.mod go.sum ./
RUN go mod download
COPY services/flight services/flight
COPY services/flight/.env .env
COPY pb pb
RUN go build -o /go/bin/app services/flight/cmd/main.go
CMD ["app"]
