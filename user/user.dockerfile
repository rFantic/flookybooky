FROM golang
WORKDIR /go/src/user
COPY go.mod go.sum ./
COPY user user
COPY user/.env .env
RUN go build -o /go/bin/app user/cmd/main.go
CMD ["app"]