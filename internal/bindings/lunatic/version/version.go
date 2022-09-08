package version

//go:wasm-module lunatic::version
//go:export major
func Major() uint32

//go:wasm-module lunatic::version
//go:export minor
func Minor() uint32

//go:wasm-module lunatic::version
//go:export patch
func Patch() uint32


