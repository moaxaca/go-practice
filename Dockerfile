# Base
FROM --platform=${BUILDPLATFORM} golang:1.16 AS base
ENV CGO_ENABLED=0
ENV GO111MODULE=on
RUN apt-get update
RUN apt install -y protobuf-compiler
RUN go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.25.0 \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0

# Dependencies
FROM base as dependencies
WORKDIR /app

ARG APPLICATION="address_validation"
ENV APPLICATION=$APPLICATION

COPY $APPLICATION/go.mod .
COPY $APPLICATION/go.sum .
RUN go mod download
CMD /bin/sh

# Codebase
FROM dependencies as code
COPY $APPLICATION/ /app
RUN go mod vendor

# Build
FROM code as build
RUN protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/grpc/proto/*.proto
RUN go build -o /main-go cmd/main.go

# Lint
FROM golangci/golangci-lint:v1.41-alpine as lint
WORKDIR /app
COPY --from=code /app /app
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
ARG PORT=8081
ENV CGO_ENABLED=0
EXPOSE 8081
RUN swagger generate spec -o swagger.json
CMD ["serve", "swagger.json", "-p", "8081", "-F", "redoc", "--no-open"]

# Production
FROM --platform=${BUILDPLATFORM} golang:1.16-alpine AS production
COPY --from=build /main-go /main-go
EXPOSE 8080
CMD ["/main-go"]
