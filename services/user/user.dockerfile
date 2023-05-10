FROM golang
RUN go install github.com/go-delve/delve/cmd/dlv@latest
WORKDIR /go/src/user
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
COPY services/user/.env .env
RUN go build -o /go/bin/app services/user/cmd/main.go
CMD ["app"]
# CMD [ "/go/bin/dlv", "--listen=:4000", "--headless=true", "--log=true", "--accept-multiclient", "--api-version=2", "exec", "/go/bin/app" ]