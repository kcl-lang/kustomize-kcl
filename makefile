VERSION:=$(shell cat VERSION)
test:
	go test ./...

fmt:
	go fmt ./...		

image:
	docker build . -t docker.io/kcllang/kustomize-kcl:$(VERSION)
	docker push docker.io/kcllang/kustomize-kcl:$(VERSION)
