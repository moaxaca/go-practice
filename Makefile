# Variables
IMAGE=go_example
TARGET=production
PLATFORM=linux/amd64 # Docker

ARCH=$(shell go env GOOS)-$(shell go env GOARCH)
platform_temp=$(subst -, ,$(ARCH))

PKG := github.com/moaxaca/paracly_api
GOOS = $(word 1, $(platform_temp))
GOARCH = $(word 2, $(platform_temp))

APP=address_validation

# Local
install:
	cd $(APP) && go mod download && go mod vendor;
# Docker
all: cleanup build run

cleanup:
	docker rm $(IMAGE) || true;

build:
	docker buildx build . \
	--tag $(IMAGE) \
	--target $(TARGET) \
   	--platform $(PLATFORM);

lint:
	docker rm $(IMAGE)-lint || true;
	docker buildx build . \
		--tag $(IMAGE)-lint \
		--target lint \
   		--platform $(PLATFORM);
	docker run $(IMAGE)-lint;

unit:
	docker rm $(IMAGE)-unit || true;
	docker buildx build . \
		--tag $(IMAGE)-unit \
		--target unit \
   		--platform $(PLATFORM);
	docker run $(IMAGE)-unit;

run: cleanup build
	docker run \
	-it \
	-p 8080:8080 \
	$(IMAGE);
