# syntax=docker/dockerfile:1

# Build Stage
# ---------------------------------------------------------
ARG GO_VERSION=1.21
FROM golang:${GO_VERSION} As build
WORKDIR /src
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/server/main.go
EXPOSE 8080
CMD ["/gossip-with-go-backend"]