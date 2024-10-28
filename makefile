VERSION:=$(shell cat VERSION)
test:
	go test ./...

fmt:
	go fmt ./...		

image:
	docker build . -t docker.io/kcllang/kustomize-kcl:v$(VERSION)
	docker push docker.io/kcllang/kustomize-kcl:v$(VERSION)

release:
	git tag v$(VERSION)
	git push origin v$(VERSION)
	gh release create v$(VERSION) --draft --generate-notes --title "$(VERSION) Release"
