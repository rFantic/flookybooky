FROM golang
RUN go install github.com/go-delve/delve/cmd/dlv@latest
WORKDIR /go/src/user
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
COPY services/user/.env .env
RUN go build -o /go/bin/app services/user/cmd/main.go
<<<<<<< HEAD
CMD ["app"]
# CMD [ "/go/bin/dlv", "--listen=:4000", "--headless=true", "--log=true", "--accept-multiclient", "--api-version=2", "exec", "/go/bin/app" ]
=======
# CMD ["app"]
CMD [ "/go/bin/dlv", "--listen=:4000", "--continue", "--headless=true", "--log=true", "--accept-multiclient", "--api-version=2", "exec", "/go/bin/app" ]
>>>>>>> 3861f27 (Precommit fix.)
