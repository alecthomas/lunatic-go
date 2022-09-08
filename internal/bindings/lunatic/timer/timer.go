package timer

//go:wasm-module lunatic::timer
//go:export send-after
func SendAfter(processId uint64, duration uint64) uint64

//go:wasm-module lunatic::timer
//go:export cancel-timer
func CancelTimer(timerId uint64) uint32


