package wasm

import _ "embed"

//go:embed spectest-interp
var SpectestInterp []byte

//go:embed wasm-decompile
var WasmDecompile []byte

//go:embed wasm-interp
var WasmInterp []byte

//go:embed wasm-objdump
var WasmObjDump []byte

//go:embed wasm-stats
var WasmStats []byte

//go:embed wasm-strip
var WasmStrip []byte

//go:embed wasm-validate
var WasmValidate []byte

//go:embed wasm2c
var Wasm2c []byte

//go:embed wasm2wat
var Wasm2Wat []byte

//go:embed wast2json
var Wast2Json []byte

//go:embed wat-desugar
var WatDesugar []byte

//go:embed wat2wasm
var Wat2Wasm []byte
