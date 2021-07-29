# Variables
APPLICATION=address_validation
CURRENT_DIR = $(shell pwd)
IMAGE=go-$(APPLICATION)
TARGET=production
PLATFORM=linux/amd64

ARCH=$(shell go env GOOS)-$(shell go env GOARCH)
platform_temp=$(subst -, ,$(ARCH))

PKG := github.com/moaxaca/paracly_api
GOOS = $(word 1, $(platform_temp))
GOARCH = $(word 2, $(platform_temp))

# Versioning Information
HASH := $(shell git rev-parse HEAD)

# Local
install:
	cd $(APPLICATION) && go mod download && go mod vendor;

# Docker
all: install cleanup build run

cleanup:
	docker rm $(IMAGE) || true;

build:
	docker buildx build . \
		--build-arg APPLICATION=$(APPLICATION) \
		--tag $(IMAGE) \
		--target $(TARGET) \
   		--platform $(PLATFORM);

lint:
	docker rm $(IMAGE)-lint || true;
	docker buildx build . \
		--tag $(IMAGE)-lint \
		--target lint \
   		--platform $(PLATFORM);
	make run IMAGE=$(IMAGE)-lint;

unit:
	docker rm $(IMAGE)-unit || true;
	docker buildx build . \
		--target unit \
		--tag $(IMAGE)-unit \
   		--platform $(PLATFORM);
	make run IMAGE=$(IMAGE)-unit;

e2e:
	docker rm $(IMAGE)-unit || true;
	docker buildx build . \
		--target e2e \
		--tag $(IMAGE)-e2e \
   		--platform $(PLATFORM);
	make run IMAGE=$(IMAGE)-e2e;

run:
	docker run --env-file .env -v $(CURRENT_DIR)/.output:/app/.output -it $(IMAGE);

serve: build
	make cleanup;
	make build TARGET=production;
	docker run -it --env-file .env -p 8080:8080 $(IMAGE);

# Swagger
generate_swagger:
	cd $(APPLICATION) && swagger generate spec -o ./api/rest/swagger.json

swagger:
	docker rm $(IMAGE)-swagger || true;
	docker buildx build . \
		--target swagger \
		--tag $(IMAGE)-swagger \
   		--platform $(PLATFORM);
	docker run -p 8081:8081 $(IMAGE)-swagger;
