FROM golang:1.18 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o kustomize-kcl


FROM kusionstack/kclvm

WORKDIR /app
USER root
COPY --from=builder /app/kustomize-kcl .
RUN mkdir -p /go/bin

CMD ["/app/kustomize-kcl"]
