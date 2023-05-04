FROM golang
WORKDIR /go/src/graphql
COPY go.mod go.sum ./
RUN go mod download
COPY graphql graphql
COPY graphql/.env .env
RUN go build -o /go/bin/app graphql/cmd/main.go
CMD ["app"]