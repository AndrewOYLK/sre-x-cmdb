.PHONY: build
build:
	go build -o cmdb main.go

.PHONE: swag
swag:
	swag init

.PHONY: run
run: swag build
	./cmdb

.PHONE: test
test:
	go test -v ./ -cover

.PHONE: docker
docker:
	CGO_ENABLED=0 GOOS=linux go build -o cmdb main.go
	docker build . -t devops/cmdb:v1.0
	rm -rf cmdb