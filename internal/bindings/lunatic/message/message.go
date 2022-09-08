package message

//go:wasm-module lunatic::message
//go:export create-data
func CreateData(tag int64, capacity uint64)

//go:wasm-module lunatic::message
//go:export write-data
func WriteData(data uint32, dataLen uint32) uint32

//go:wasm-module lunatic::message
//go:export read-data
func ReadData(data uint32, dataLen uint32) uint32

//go:wasm-module lunatic::message
//go:export seek-data
func SeekData(position uint64)

//go:wasm-module lunatic::message
//go:export get-tag
func GetTag() int64

//go:wasm-module lunatic::message
//go:export data-size
func DataSize() uint64

//go:wasm-module lunatic::message
//go:export push-process
func PushProcess(processId uint64) uint64

//go:wasm-module lunatic::message
//go:export take-process
func TakeProcess(index uint64) uint64

//go:wasm-module lunatic::message
//go:export push-tcp-stream
func PushTcpStream(tcpStreamId uint64) uint64

//go:wasm-module lunatic::message
//go:export take-tcp-stream
func TakeTcpStream(index uint64) uint64

//go:wasm-module lunatic::message
//go:export send
func Send(processId uint64)

//go:wasm-module lunatic::message
//go:export send-receive-skip-search
func SendReceiveSkipSearch(processId uint64, timeout uint32) uint32

//go:wasm-module lunatic::message
//go:export receive
func Receive(tag uint32, tagLen uint32, timeout uint32) uint32


