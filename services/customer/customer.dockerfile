FROM golang
WORKDIR /go/src/customer
COPY go.mod go.sum ./
RUN go mod download
COPY services/customer services/customer
COPY internal internal
COPY services/customer/.env .env
# RUN go build -o /go/bin/app services/customer/cmd/main.go
# CMD ["app"]
