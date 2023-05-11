FROM golang
WORKDIR /go/src/graphql
RUN go install github.com/go-delve/delve/cmd/dlv@latest
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
COPY services/graphql/.env .env
RUN go build -gcflags="all=-N -l" -o /go/bin/app services/graphql/cmd/main.go
# CMD ["app"]
CMD [ "/go/bin/dlv", "--listen=:4000", "--continue", "--headless=true", "--log=true", "--accept-multiclient", "--api-version=2", "exec", "/go/bin/app" ]
