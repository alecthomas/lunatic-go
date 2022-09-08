package error

//go:wasm-module lunatic::error
//go:export string_size
func StringSize(caller uint64) uint32

//go:wasm-module lunatic::error
//go:export to_string
func ToString(errorId uint64, errorStr uint32)

//go:wasm-module lunatic::error
//go:export drop
func Drop(errorId uint64)


