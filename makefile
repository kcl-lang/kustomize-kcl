test:
	go test ./...

fmt:
	go fmt ./...		

image:
	docker build . -t docker.io/peefyxpf/kustomize-kcl:v0.1.0
	docker push docker.io/peefyxpf/kustomize-kcl:v0.1.0
