FROM golang
WORKDIR /go/src/graphql
RUN go install github.com/go-delve/delve/cmd/dlv@v1.20.0
COPY go.mod go.sum ./
RUN go mod download
COPY services/graphql services/graphql
COPY services/graphql/.env .env
COPY internal internal
COPY middleware middleware
COPY pb pb
RUN go build -gcflags="all=-N -l" -o /go/bin/app services/graphql/cmd/main.go
# CMD ["app"]
CMD [ "/go/bin/dlv", "--listen=:4000", "--continue", "--headless=true", "--log=true", "--accept-multiclient", "--api-version=2", "exec", "/go/bin/app" ]
