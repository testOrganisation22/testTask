FROM golang:1.21.5
WORKDIR /app
COPY . .
RUN go mod vendor

go test ./...

WORKDIR ./cmd
RUN go build -o main .
CMD ["./main"]