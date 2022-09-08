package wasi

//go:wasm-module lunatic::wasi
//go:export config_add_environment_variable
func ConfigAddEnvironmentVariable(configId uint64, key uint32, keyLen uint32, value uint32, valueLen uint32)

//go:wasm-module lunatic::wasi
//go:export config_add_command_line_argument
func ConfigAddCommandLineArgument(configId uint64, key uint32, keyLen uint32)

//go:wasm-module lunatic::wasi
//go:export config_preopen_dir
func ConfigPreopenDir(configId uint64, key uint32, keyLen uint32)


