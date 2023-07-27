FROM golang:1.19 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o kustomize-kcl

FROM kcllang/kcl

WORKDIR /app
COPY --from=builder /app/kustomize-kcl .

ENV KCL_GO_DISABLE_ARTIFACT=on

CMD ["/app/kustomize-kcl"]
