# go-wabt

go-wabt is a distribution of [wabt][1], that can be built with Go. It does not actually reimplement any
functionality of wabt in Go, instead the original WebAssembly binaries, and 
executing with the pure Go Wasm runtime [wazero][2]. This means that `go install` or `go run`
can be used to execute it, with no need to rely on external package managers such as Homebrew,
on any platform that Go supports.

## Installation

Precompiled binaries are available in the [releases](https://github.com/wasilibs/go-wabt/releases).
Alternatively, install the plugin you want using `go install`.

```bash
$ go install github.com/wasilibs/go-wabt/cmd/wasm2wat@latest
```

To avoid installation entirely, it can be convenient to use `go run`

```bash
$ go run github.com/wasilibs/go-wabt/cmd/wasm2wat@latest my.wasm
```

Note that due to the sandboxing of the filesystem when using Wasm, currently only files that descend
from the current directory when executing the tool are accessible to it, i.e., `../wasm/my.wasm` or
`/separate/root/my.wasm` will not be found.

[1]: https://github.com/WebAssembly/wabt
[2]: https://wazero.io/
