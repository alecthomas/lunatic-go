package error

//go:wasm-module lunatic::error
//go:export string-size
func StringSize(caller uint64) uint32

//go:wasm-module lunatic::error
//go:export to-string
func ToString(errorId uint64, errorStr uint32)

//go:wasm-module lunatic::error
//go:export drop
func Drop(errorId uint64)


