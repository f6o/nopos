PROTO_IMAGE ?= nopos-proto-builder
DOCKER ?= docker

.PHONY: proto proto-image prisma

proto-image:
	$(DOCKER) build -f Dockerfile.proto -t $(PROTO_IMAGE) .

proto: proto-image hands/hands.proto
	$(DOCKER) run --rm \
		-u $(shell id -u):$(shell id -g) \
		-v $(CURDIR):/work \
		-w /work \
		$(PROTO_IMAGE) \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		./hands/hands.proto

prisma:
	go run github.com/steebchen/prisma-client-go db push
