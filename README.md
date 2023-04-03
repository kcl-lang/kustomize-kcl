# Kustomize KCL Function

[![Go Report Card](https://goreportcard.com/badge/github.com/KusionStack/kustomize-kcl)](https://goreportcard.com/report/github.com/KusionStack/kustomize-kcl)
[![GoDoc](https://godoc.org/github.com/KusionStack/kustomize-kcl?status.svg)](https://godoc.org/github.com/KusionStack/kustomize-kcl)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/KusionStack/kustomize-kcl/blob/main/LICENSE)

This is an example of implementing a KCL function for Kustomize. [KCL](https://github.com/KusionStack/KCLVM) is a constraint-based record & functional domain language. Full documents of KCL can be found [here](https://kcl-lang.io/).

## Function Implementation

The function is implemented as an image, and built using `make image`.

The function is implemented as a go program, which reads a collection of input Resource configuration, passing them to KCL.

## Function Configuration

See the API struct definition in `main.go` for documentation.

+ `source` - the KCL function code.
+ `params` - top-level arguments for KCL

## Function invocation

The function is invoked by authoring a local Resource with `metadata.annotations.[config.kubernetes.io/function]` and running:

```shell
sudo kustomize fn run examples/set-annotation/local-resource/ --as-current-user --dry-run
```

This exists non-zero if the KCL code has no errors.

## Guides for Developing KCL

Here's what you can do in the KCL script:

+ Read resources from `option("resource_list")`. The `option("resource_list")` complies with the [KRM Functions Specification](https://kpt.dev/book/05-developing-functions/01-functions-specification). You can read the input resources from `option("resource_list")["items"]` and the `functionConfig` from `option("resource_list")["functionConfig"]`.
+ Return a KPM list for output resources.
+ Read the environment variables. e.g. `option("PATH")`.
+ Read the OpenAPI schema. e.g. `option("open_api")["definitions"]["io.k8s.api.apps.v1.Deployment"]`
+ Return an error using `assert {condition}, {error_message}`.

## Library

You can directly use [KCL standard libraries](https://kcl-lang.io/docs/reference/model/overview) without importing them, such as `regex.match`, `math.log`.
