package distributed

//go:wasm-module lunatic::distributed
//go:export nodes_count
func NodesCount() uint32

//go:wasm-module lunatic::distributed
//go:export get_nodes
func GetNodes(nodesPtr uint32, nodesLen uint32) uint32

//go:wasm-module lunatic::distributed
//go:export node_id
func NodeId() uint64

//go:wasm-module lunatic::distributed
//go:export module_id
func ModuleId() uint64

//go:wasm-module lunatic::distributed
//go:export spawn
func Spawn(nodeId uint64, configId int64, moduleId uint64, funcStrPtr uint32, funcStrLen uint32, paramsPtr uint32, paramsLen uint32, idPtr uint32) uint32

//go:wasm-module lunatic::distributed
//go:export send
func Send(nodeId uint64, processId uint64)

//go:wasm-module lunatic::distributed
//go:export send_receive_skip_search
func SendReceiveSkipSearch(nodeId uint64, processId uint64, timeoutDuration uint64) uint32


