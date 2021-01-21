# go-ihex

![CI](https://github.com/mihaigalos/go-ihex/workflows/CI/badge.svg) [![codecov](https://codecov.io/gh/mihaigalos/go-ihex/branch/main/graph/badge.svg?token=YRA05WPK3f)](https://codecov.io/gh/mihaigalos/go-ihex) [![license](https://img.shields.io/badge/license-GPLv3-brightgreen.svg)](LICENSE) [![LoC](https://tokei.rs/b1/github/mihaigalos/go-ihex)](https://github.com/Aaronepower/tokei)

Parser for intel hex file format, bazelized and easily importable as a Go library.

### How to use

`WORKSPACE`:
```bazel
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "207fad3e6689135c5d8713e5a17ba9d1290238f47b9ba545b63d9303406209c6",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.24.7/rules_go-v0.24.7.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.24.7/rules_go-v0.24.7.tar.gz",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "b85f48fa105c4403326e9525ad2b2cc437babaa6e15a3fc0b1dbab0ab064bc7c",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.2/bazel-gazelle-v0.22.2.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.2/bazel-gazelle-v0.22.2.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_rules_dependencies()

go_register_toolchains()

gazelle_dependencies()

go_repository(
    name = "go-ihex",
    commit = "93e825890b53ec109a8846b87d63bba826c17056",
    importpath = "github.com/mihaigalos/go-ihex",
)
```

`BUILD.bazel`:
```bazel
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "host_lib",
    srcs = [
        "logic.go",
        "main.go",
    ],
    importpath = "github.com/mihaigalos/uartboot/host",
    visibility = ["//visibility:private"],
    deps = ["@go-ihex//parser:parser_lib"],
)

go_binary(
    name = "host",
    embed = [":host_lib"],
    visibility = ["//visibility:public"],
)
```

`logic.go`:

```go
package main

import (
	"fmt"

	myparser "github.com/mihaigalos/go-ihex/parser"
)

func parseFile() {
	file := []string{
		":100000000C9434000C944F000C944F000C944F004F",
		":0800F00091E00E945700F0CFDF",
		":00000001FF"}

	actual := myparser.IsFileValid(file)
	fmt.Printf("File check: %d\n", actual)
}
```