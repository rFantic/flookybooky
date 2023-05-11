FROM golang
WORKDIR /go/src/customer
RUN go install github.com/go-delve/delve/cmd/dlv@v1.20.0
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
COPY services/customer/.env .env
RUN go build -gcflags="all=-N -l" -o /go/bin/app services/customer/cmd/main.go
# CMD ["app"]
CMD [ "/go/bin/dlv", "--listen=:4000", "--headless=true", "--log=true", "--accept-multiclient", "--api-version=2", "exec", "/go/bin/app" ]
