# intel-hex-parser

![CI](https://github.com/mihaigalos/intel-hex-parser/workflows/CI/badge.svg) [![codecov](https://codecov.io/gh/mihaigalos/intel-hex-parser/branch/main/graph/badge.svg?token=YRA05WPK3f)](https://codecov.io/gh/mihaigalos/intel-hex-parser) [![license](https://img.shields.io/badge/license-GPLv3-brightgreen.svg)](LICENSE) [![LoC](https://tokei.rs/b1/github/mihaigalos/intel-hex-parser)](https://github.com/Aaronepower/tokei)

Parser for intel hex file format, bazelized and easily importable as a Go library.

### How to use

Example usage, BUILD.bazel:
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
    deps = ["@intel-hex-parser//parser:parser_lib"],
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

	myparser "github.com/mihaigalos/intel-hex-parser/parser"
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