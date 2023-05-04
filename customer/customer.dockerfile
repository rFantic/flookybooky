FROM golang
WORKDIR /go/src/customer
COPY go.mod go.sum ./
RUN go mod download
COPY customer customer
COPY customer/.env .env
RUN go build -o /go/bin/app customer/cmd/main.go
CMD ["app"]