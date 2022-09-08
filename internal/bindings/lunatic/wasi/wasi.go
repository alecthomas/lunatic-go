package wasi

//go:wasm-module lunatic::wasi
//go:export config-add-environment-variable
func ConfigAddEnvironmentVariable(configId uint64, key uint32, keyLen uint32, value uint32, valueLen uint32)

//go:wasm-module lunatic::wasi
//go:export config-add-command-line-argument
func ConfigAddCommandLineArgument(configId uint64, key uint32, keyLen uint32)

//go:wasm-module lunatic::wasi
//go:export config-preopen-dir
func ConfigPreopenDir(configId uint64, key uint32, keyLen uint32)


