

# Get the latest Git tag for versioning
GIT_TAG := $(shell git describe --tags --always --dirty)


# DEV

.PHONY: run
run:
	go build -ldflags="-X 'GoStreamRecord/internal/db.Version=$(GIT_TAG)'" -o server main.go && ./server





# DOCKER
.PHONY: build
build: base app

.PHONY: base
base: push-base
	docker build \
		--build-arg TAG=$(GIT_TAG) \
		-t lunanightbyte/gorecord-base:$(GIT_TAG) . \
		-f ./docker/Dockerfile.base \

.PHONY: app
app: push-app
	docker build \
		--build-arg TAG=$(GIT_TAG) \
		-t lunanightbyte/gorecord:$(GIT_TAG) . \
		-f ./docker/Dockerfile.run \
	docker push lunanightbyte/gorecord:$(GIT_TAG)

.PHONY: push-app
push-app:
	docker push lunanightbyte/gorecord:$(GIT_TAG)

.PHONY: push-app
push-base:
	docker push lunanightbyte/gorecord-base:$(GIT_TAG)

# Clean up dangling images
.PHONY: clean
clean:
	docker image prune -f

