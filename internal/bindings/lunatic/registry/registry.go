package registry

//go:wasm-module lunatic::registry
//go:export put
func Put(name uint32, nameLen uint32, processId uint64)

//go:wasm-module lunatic::registry
//go:export get
func Get(name uint32, nameLen uint32, processId uint32) uint32

//go:wasm-module lunatic::registry
//go:export remove
func Remove(name uint32, nameLen uint32)


