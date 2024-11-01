FROM golang:1.23 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

ENV CGO_ENABLED=0
RUN go build -o kustomize-kcl

FROM kcllang/kcl

WORKDIR /app
COPY --from=builder /app/kustomize-kcl .

CMD ["/app/kustomize-kcl"]
