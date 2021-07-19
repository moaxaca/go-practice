# Base
FROM --platform=${BUILDPLATFORM} golang:1.16-alpine AS base
RUN apk add build-base

# Codebase
FROM base as code
WORKDIR /app

ARG APPLICATION="address_validation"

COPY $APPLICATION/go.mod .
COPY $APPLICATION/go.sum .
RUN go mod download

COPY $APPLICATION/api api
COPY $APPLICATION/cmd cmd
COPY $APPLICATION/pkg pkg

# Build
FROM code as build
RUN go build -o /main-go cmd/main.go
CMD /bin/sh

# Lint
FROM golangci/golangci-lint:v1.41-alpine as lint
WORKDIR /app
COPY --from=build /app /app
COPY golangci.yaml golangci.yaml
CMD golangci-lint run -c golangci.yaml

# Unit Tests
FROM code as unit
WORKDIR /app
COPY --from=build /app /app
CMD go test

# Production
FROM base AS production
WORKDIR /
COPY --from=build /main-go /main-go
EXPOSE 8080
CMD ["/main-go"]
