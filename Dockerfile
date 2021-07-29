# Base
FROM --platform=${BUILDPLATFORM} golang:1.16-alpine AS base
RUN apk add build-base

# Dependencies
FROM base as dependencies
WORKDIR /app

ARG APPLICATION="address_validation"
ENV APPLICATION=$APPLICATION

COPY $APPLICATION/go.mod .
COPY $APPLICATION/go.sum .
RUN go mod download
RUN go mod vendor
CMD /bin/sh

# Codebase
FROM dependencies as code
WORKDIR /app
COPY $APPLICATION/. .

# Build
FROM code as build
RUN go build -o /main-go cmd/main.go

# Lint
FROM golangci/golangci-lint:v1.41-alpine as lint
WORKDIR /app
COPY --from=build /app /app
COPY golangci.yaml golangci.yaml
CMD golangci-lint run -c golangci.yaml

# Unit Tests
FROM code as unit
CMD go test $(go list ./... | grep -v /test/) -coverprofile .output/unit-coverage.out

# E2E Tests
FROM code as e2e
CMD go test $(go list ./... | grep /test/) -coverprofile .output/e2e-coverage.out

# Swagger
FROM quay.io/goswagger/swagger AS swagger
WORKDIR /app
COPY --from=code /app /app
EXPOSE 8081
RUN swagger generate spec -o swagger.json
CMD ["serve", "swagger.json", "-p", "8081", "-F", "redoc", "--no-open"]

# Production
FROM base AS production
WORKDIR /
COPY --from=build /main-go /main-go
EXPOSE 8080
CMD ["/main-go"]
