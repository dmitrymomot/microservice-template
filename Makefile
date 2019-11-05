.PHONY: help
help: ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: proto
proto: ## Compile protobuf for golang
	protoc -I /usr/local/include -I . \
		-I$(GOPATH)/src \
		-I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
		--go_out=plugins=grpc:. \
		--twirp_out=. \
		--validate_out=lang=go:. \
		pb/**/*.proto

.PHONY: proto-dart
proto-dart: ## Compile protobuf for dart
	protoc -I /usr/local/include -I . \
		-I$(GOPATH)/src \
		--dart_out=grpc:. \
		pb/**/*.proto

.PHONY: build
build: proto ## Build application and compile protobuf for golang
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix nocgo -o app .

.PHONY: docker
docker: ## Build docker image
	docker build . -t microservice-template:latest

.PHONY: deploy
deploy: ## Deploy pods to kubernetes
	kubectl apply -f k8s/deploy.yml

.PHONY: lb
lb: ## Run load balancer service
	kubectl apply -f k8s/lb.yml

.PHONY: svc
svc: ## Run service without external port
	kubectl apply -f k8s/service.yml

.PHONY: down
down: ## Down pods
	kubectl delete -f k8s/deploy.yml

.PHONY: down-lb
down-lb: ## Down load balancer
	kubectl delete -f k8s/deploy.yml -f k8s/lb.yml

.PHONY: down-svc
down-svc: ## Down service wrapper
	kubectl delete -f k8s/deploy.yml -f k8s/service.yml

.PHONY: reload
reload: down deploy info ## Reload after app was rebuilt

.PHONY: info
info: ## Get cluster info
	@kubectl get all

.PHONY: logs
log: ## Show logs
	@kubectl logs $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS)) --all-containers=true