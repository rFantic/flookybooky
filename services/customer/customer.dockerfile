FROM golang
WORKDIR /go/src/customer
RUN go install github.com/go-delve/delve/cmd/dlv@v1.20.0
COPY go.mod go.sum ./
RUN go mod download
COPY services/customer services/customer
COPY services/customer/.env .env
COPY internal internal
COPY pb pb
RUN go build -gcflags="all=-N -l" -o /go/bin/app services/customer/cmd/main.go
CMD ["app"]
# CMD [ "/go/bin/dlv", "--continue", "--listen=:4000", "--headless=true", "--log=true", "--accept-multiclient", "--api-version=2", "exec", "/go/bin/app" ]
