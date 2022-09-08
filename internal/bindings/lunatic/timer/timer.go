package timer

//go:wasm-module lunatic::timer
//go:export send_after
func SendAfter(processId uint64, duration uint64) uint64

//go:wasm-module lunatic::timer
//go:export cancel_timer
func CancelTimer(timerId uint64) uint32


