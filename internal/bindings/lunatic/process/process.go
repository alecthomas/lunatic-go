package process

//go:wasm-module lunatic::process
//go:export compile-module
func CompileModule(data uint32, dataLen uint32, id uint32) int32

//go:wasm-module lunatic::process
//go:export drop-module
func DropModule(configId uint64)

//go:wasm-module lunatic::process
//go:export create-config
func CreateConfig() uint64

//go:wasm-module lunatic::process
//go:export drop-config
func DropConfig(configId uint64)

//go:wasm-module lunatic::process
//go:export config-set-max-memory
func ConfigSetMaxMemory(configId uint64, maxMemory uint64)

//go:wasm-module lunatic::process
//go:export config-get-max-memory
func ConfigGetMaxMemory(configId uint64) uint64

//go:wasm-module lunatic::process
//go:export config-set-max-fuel
func ConfigSetMaxFuel(configId uint64, maxFuel uint64)

//go:wasm-module lunatic::process
//go:export config-get-max-fuel
func ConfigGetMaxFuel(configId uint64) uint64

//go:wasm-module lunatic::process
//go:export config-can-compile-modules
func ConfigCanCompileModules(configId uint64) uint32

//go:wasm-module lunatic::process
//go:export config-set-can-compile-modules
func ConfigSetCanCompileModules(configId uint64, can uint32)

//go:wasm-module lunatic::process
//go:export config-can-create-configs
func ConfigCanCreateConfigs(configId uint64) uint32

//go:wasm-module lunatic::process
//go:export config-set-can-create-configs
func ConfigSetCanCreateConfigs(configId uint64, can uint32)

//go:wasm-module lunatic::process
//go:export config-can-spawn-processes
func ConfigCanSpawnProcesses(configId uint64) uint32

//go:wasm-module lunatic::process
//go:export config-set-can-spawn-processes
func ConfigSetCanSpawnProcesses(configId uint64, can uint32)

//go:wasm-module lunatic::process
//go:export spawn
func Spawn(link int64, configId int64, moduleId int64, function uint32, functionLen uint32, params uint32, paramsLen uint32, id uint32) uint32

//go:wasm-module lunatic::process
//go:export drop-process
func DropProcess(processId uint64)

//go:wasm-module lunatic::process
//go:export clone-process
func CloneProcess(processId uint64) uint64

//go:wasm-module lunatic::process
//go:export sleep-ms
func SleepMs(millis uint64)

//go:wasm-module lunatic::process
//go:export die-when-link-dies
func DieWhenLinkDies(trap uint32)

//go:wasm-module lunatic::process
//go:export this
func This() uint64

//go:wasm-module lunatic::process
//go:export id
func Id(processId uint64, uuid uint32)

//go:wasm-module lunatic::process
//go:export link
func Link(tag int64, processId uint64)

//go:wasm-module lunatic::process
//go:export unlink
func Unlink(processId uint64)

//go:wasm-module lunatic::process
//go:export kill
func Kill(processId uint64)


