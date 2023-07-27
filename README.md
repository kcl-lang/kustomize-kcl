# Kustomize KCL Function

[![Go Report Card](https://goreportcard.com/badge/github.com/kcl-lang/kustomize-kcl)](https://goreportcard.com/report/github.com/kcl-lang/kustomize-kcl)
[![GoDoc](https://godoc.org/github.com/kcl-lang/kustomize-kcl?status.svg)](https://godoc.org/github.com/kcl-lang/kustomize-kcl)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/kcl-lang/kustomize-kcl/blob/main/LICENSE)

This is an example of implementing a KCL function for Kustomize. [KCL](https://github.com/kcl-lang/kcl) is a constraint-based record & functional domain language. Full documents of KCL can be found [here](https://kcl-lang.io/).

KCL can be used to create functions to mutate and/or validate the YAML Kubernetes Resource Model (KRM) input/output format, and we provide Kustomize KCL functions to simplify the function authoring process.

## Quick Start

Let’s write a KCL function that adds annotation `managed-by=kustomize-kcl` only to Deployment resources.

### Prerequisites

+ Install [kustomize](https://github.com/kubernetes-sigs/kustomize)
+ Install [Docker](https://www.docker.com/)

### Test and Run

#### Local Test

```bash
cat ./examples/set-annotation/local-resource/example-use.yaml | go run main.go
```

The output YAML is

```yaml
apiVersion: krm.kcl.dev/v1alpha1
kind: KCLRun
metadata:
  name: set-annotation
  annotations:
    config.kubernetes.io/function: |
      container:
        image: docker.io/kcllang/kustomize-kcl:v0.2.0
    config.kubernetes.io/path: example-use.yaml
    internal.config.kubernetes.io/path: example-use.yaml
# EDIT THE SOURCE!
# This should be your KCL code which preloads the `ResourceList` to `option("resource_list")
spec:
  source: |
    [resource | {if resource.kind == "Deployment": metadata.annotations: {"managed-by" = "kustomize-kcl"}} for resource in option("resource_list").items]
---
apiVersion: v1
kind: Service
metadata:
  name: test
  annotations:
    config.kubernetes.io/path: example-use.yaml
    internal.config.kubernetes.io/path: example-use.yaml
spec:
  selector:
    app: MyApp
  ports:
  - protocol: TCP
    port: 80
    targetPort: 9376
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
  annotations:
    managed-by: kustomize-kcl
    config.kubernetes.io/path: example-use.yaml
    internal.config.kubernetes.io/path: example-use.yaml
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80
```

#### Integration Test

```bash
# Note: you need add sudo and --as-current-user flags to ensure KCL has permission to write temp files in the container filesystem.
kustomize fn run examples/set-annotation/local-resource/ --dry-run
```

The output YAML is

```yaml
apiVersion: krm.kcl.dev/v1alpha1
kind: KCLRun
metadata:
  name: set-annotation
  annotations:
    config.kubernetes.io/function: |
      container:
        image: docker.io/kcllang/kustomize-kcl:v0.2.0
    config.kubernetes.io/path: example-use.yaml
    internal.config.kubernetes.io/path: example-use.yaml
# EDIT THE SOURCE!
# This should be your KCL code which preloads the `ResourceList` to `option("resource_list")
spec:
  source: |
    [resource | {if resource.kind == "Deployment": metadata.annotations: {"managed-by" = "kustomize-kcl"}} for resource in option("resource_list").items]
---
apiVersion: v1
kind: Service
metadata:
  name: test
  annotations:
    config.kubernetes.io/path: example-use.yaml
    internal.config.kubernetes.io/path: example-use.yaml
spec:
  selector:
    app: MyApp
  ports:
  - protocol: TCP
    port: 80
    targetPort: 9376
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
  annotations:
    managed-by: kustomize-kcl
    config.kubernetes.io/path: example-use.yaml
    internal.config.kubernetes.io/path: example-use.yaml
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80
```

This exists non-zero if the KCL code has no errors.

## Function Implementation

The function is implemented as an image, and built using `make image`, and its function is implemented as a go program, which reads a collection of input Resource configuration, passing them to KCL.

## Function Configuration

See the API struct definition in `main.go` for documentation.

+ `source` - the KCL function code.
+ `params` - top-level arguments for KCL

## Function Invocation

The function is invoked by authoring a local Resource with `metadata.annotations.[config.kubernetes.io/function]`.

## Guides for Developing KCL

Here's what you can do in the KCL script:

+ Read resources from `option("resource_list")`. The `option("resource_list")` complies with the [KRM Functions Specification](https://kpt.dev/book/05-developing-functions/01-functions-specification). You can read the input resources from `option("resource_list")["items"]` and the `functionConfig` from `option("resource_list")["functionConfig"]`.
+ Return a KPM list for output resources.
+ Return an error using `assert {condition}, {error_message}`.
+ Read the environment variables. e.g. `option("PATH")` (Not yet implemented).
+ Read the OpenAPI schema. e.g. `option("open_api")["definitions"]["io.k8s.api.apps.v1.Deployment"]` (Not yet implemented).

## Examples

See [here](https://kcl-lang.io/krm-kcl/tree/main/examples) for more examples.

## Library

You can directly use [KCL standard libraries](https://kcl-lang.io/docs/reference/model/overview) without importing them, such as `regex.match`, `math.log`.
