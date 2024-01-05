package main

import (
	"github.com/wasilibs/go-wabt/internal/runner"
	"github.com/wasilibs/go-wabt/internal/wasm"
)

func main() {
	runner.Run("wasm-interp", wasm.WasmInterp)
}
