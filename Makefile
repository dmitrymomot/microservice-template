.PHONY: proto
proto:
	protoc -I /usr/local/include -I . \
		-I$(GOPATH)/src \
		-I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
		--go_out=plugins=grpc:. \
		--twirp_out=. \
		--validate_out=lang=go:. \
		proto/microservice-template/*.proto

.PHONY: proto-dart
proto-dart:
	protoc -I /usr/local/include -I . \
		-I$(GOPATH)/src \
		--dart_out=grpc:. \
		proto/microservice-template/*.proto

.PHONY: build
build: proto
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix nocgo -o app .

.PHONY: docker
docker:
	docker build . -t microservice-template:latest

.PHONY: up
up:
	kubectl apply -f k8s/deploy.yml -f k8s/service.yml

.PHONY: lb
lb:
	kubectl apply -f k8s/deploy.yml -f k8s/lb.yml

.PHONY: down
down:
	kubectl delete -f k8s/deploy.yml -f k8s/lb.yml -f k8s/service.yml