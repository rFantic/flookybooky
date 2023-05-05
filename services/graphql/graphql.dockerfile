FROM golang
WORKDIR /go/src/graphql
COPY go.mod go.sum ./
RUN go mod download
COPY services/graphql services/graphql
COPY services/graphql/.env .env
RUN go build -o /go/bin/app services/graphql/cmd/main.go
CMD ["app"]