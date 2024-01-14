# syntax=docker/dockerfile:1

# Build Stage
# ---------------------------------------------------------
ARG GO_VERSION=1.21
FROM golang:${GO_VERSION} As build

# Set workdir as /src
WORKDIR /src

# Copy Source code into container (excluding file types in dockerignore)
COPY . .

# Mount a cache for faster, subsequent builds
# Mount binds user the host's go.sum & go.mod files container
# Binds do not copy over the files to the container.
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/server/ ./cmd/server/

#Trimming stage
# ------------------------------------------------------
FROM alpine:latest AS trim

# Copy the executable from the build stage
COPY --from=build src/bin/server ./bin/

EXPOSE 8080
CMD ["/gossip-with-go-backend"]