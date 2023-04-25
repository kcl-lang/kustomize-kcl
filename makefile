VERSION:=$(shell cat VERSION)
test:
	go test ./...

fmt:
	go fmt ./...		

image:
	docker build . -t docker.io/peefyxpf/kustomize-kcl:$(VERSION)
	docker push docker.io/peefyxpf/kustomize-kcl:$(VERSION)
