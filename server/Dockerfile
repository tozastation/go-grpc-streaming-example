# Build Container
FROM golang:latest as builder
WORKDIR /go/src/github.com/tozastation/go-grpc-streaming-example
COPY . .
# Set Environment Variable
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
# Build
RUN go build -o app main.go

# Runtime Container
FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/tozastation/go-grpc-streaming-example/app /app
ENV CONNECTION_STRING=Data Source=127.0.0.1:1433;Initial Catalog=Weather;User ID=SA;Password=Test@1234
EXPOSE 3001
ENTRYPOINT ["/app"]