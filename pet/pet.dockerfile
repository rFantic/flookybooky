FROM golang
# RUN CGO_ENABLED=0 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest
WORKDIR /go/src/pet
COPY go.mod go.sum ./
RUN go mod download
COPY pet pet
COPY pet/.env .env
RUN go build -o /go/bin/app pet/cmd/main.go
# CMD ["/go/bin/dlv", "--listen=:4000", "--headless=true", "exec","/go/bin/app"]
CMD ["app"]