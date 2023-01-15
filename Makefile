SERVICE_NAME=template
DOCKER_IMAGE_NAME=registry.ecobin.ir/ecomicro/${SERVICE_NAME}
BINARY=${SERVICE_NAME}.o

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

build: clean
	go build -o ${BINARY} main.go

compose:
	docker compose up -d

compose-dev:
	docker compose -f docker-compose-dev.yml up -d

run-foo:
	go run ./extra/foo-main.go

run: build
	$(build)
	./${BINARY}

watch: compose-dev 
	/bin/sh -c 'go run ./extra/foo-main.go &'
	sleep 1
	air

docker:
	docker build --build-arg USERNAME=$(USERNAME) --build-arg APIKEY=$(APIKEY) -t ${DOCKER_IMAGE_NAME}:$(VERSION) .

docker-push:
	docker push ${DOCKER_IMAGE_NAME}:$(VERSION)

test:
	go test ./...

