# Variables
APPLICATION=address_validation
CURRENT_DIR = $(shell pwd)
TAG=go-$(APPLICATION)
TARGET=production
PLATFORM=linux/amd64

RUN_BUILD=echo "build skip";
ifeq ($(BUILD_TOO),true)
	RUN_BUILD=make build;
endif
RUN_FLAGS=-i

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
	docker rm $(TAG) || true;

build:
	docker buildx build . \
		--build-arg APPLICATION=$(APPLICATION) \
		--tag $(TAG) \
		--target $(TARGET) \
   		--platform $(PLATFORM);

lint:
	docker rm $(TAG)-lint || true;
	docker buildx build . \
		--tag $(TAG)-lint \
		--target lint \
   		--platform $(PLATFORM);
	make run TAG=$(TAG)-lint;

unit:
	docker rm $(TAG)-unit || true;
	docker buildx build . \
		--target unit \
		--tag $(TAG)-unit \
   		--platform $(PLATFORM);
	make run TAG=$(TAG)-unit;

e2e:
	docker rm $(TAG)-unit || true;
	docker buildx build . \
		--target e2e \
		--tag $(TAG)-e2e \
   		--platform $(PLATFORM);
	make run TAG=$(TAG)-e2e;

run:
	$(RUN_BUILD)
	docker run -v $(CURRENT_DIR)/.output:/app/.output $(RUN_FLAGS) -t $(TAG);

serve: build
	make cleanup;
	make build TARGET=production;
	docker run -it --env-file .env -p 8080:8080 $(TAG);

# Swagger
generate_swagger:
	cd $(APPLICATION) && swagger generate spec -o ./api/rest/swagger.json

swagger:
	docker rm $(TAG)-swagger || true;
	docker buildx build . \
		--target swagger \
		--tag $(TAG)-swagger \
   		--platform $(PLATFORM);
	docker run -p 8081:8081 $(TAG)-swagger;
