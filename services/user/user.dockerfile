FROM golang
WORKDIR /go/src/user
COPY go.mod go.sum ./
RUN go mod download
COPY services/user services/user
COPY services/user/.env .env
RUN go build -o /go/bin/app services/user/cmd/main.go
CMD ["app"]