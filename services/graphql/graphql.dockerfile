FROM golang
WORKDIR /go/src/graphql
RUN go install github.com/go-delve/delve/cmd/dlv@latest
COPY go.mod go.sum ./
RUN go mod download
COPY services/graphql services/graphql
COPY internal internal
COPY middleware middleware
COPY services/graphql/.env .env
RUN go build -o /go/bin/app services/graphql/cmd/main.go
CMD ["app"]
# CMD [ "/go/bin/dlv", "--listen=:4000", "--headless=true", "--log=true", "--accept-multiclient", "--api-version=2", "exec", "/go/bin/app" ]
