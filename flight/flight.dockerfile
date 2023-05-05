FROM golang
WORKDIR /go/src/flight
COPY go.mod go.sum ./
RUN go mod download
COPY flight flight
COPY flight/.env .env
RUN go build -o /go/bin/app flight/cmd/main.go
CMD ["app"]