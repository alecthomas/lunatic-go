package message

//go:wasm-module lunatic::message
//go:export create_data
func CreateData(tag int64, capacity uint64)

//go:wasm-module lunatic::message
//go:export write_data
func WriteData(data uint32, dataLen uint32) uint32

//go:wasm-module lunatic::message
//go:export read_data
func ReadData(data uint32, dataLen uint32) uint32

//go:wasm-module lunatic::message
//go:export seek_data
func SeekData(position uint64)

//go:wasm-module lunatic::message
//go:export get_tag
func GetTag() int64

//go:wasm-module lunatic::message
//go:export data_size
func DataSize() uint64

//go:wasm-module lunatic::message
//go:export push_process
func PushProcess(processId uint64) uint64

//go:wasm-module lunatic::message
//go:export take_process
func TakeProcess(index uint64) uint64

//go:wasm-module lunatic::message
//go:export push_tcp_stream
func PushTcpStream(tcpStreamId uint64) uint64

//go:wasm-module lunatic::message
//go:export take_tcp_stream
func TakeTcpStream(index uint64) uint64

//go:wasm-module lunatic::message
//go:export send
func Send(processId uint64)

//go:wasm-module lunatic::message
//go:export send_receive_skip_search
func SendReceiveSkipSearch(processId uint64, timeout uint32) uint32

//go:wasm-module lunatic::message
//go:export receive
func Receive(tag uint32, tagLen uint32, timeout uint32) uint32


