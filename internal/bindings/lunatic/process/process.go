package process

//go:wasm-module lunatic::process
//go:export compile_module
func CompileModule(data uint32, dataLen uint32, id uint32) int32

//go:wasm-module lunatic::process
//go:export drop_module
func DropModule(configId uint64)

//go:wasm-module lunatic::process
//go:export create_config
func CreateConfig() uint64

//go:wasm-module lunatic::process
//go:export drop_config
func DropConfig(configId uint64)

//go:wasm-module lunatic::process
//go:export config_set_max_memory
func ConfigSetMaxMemory(configId uint64, maxMemory uint64)

//go:wasm-module lunatic::process
//go:export config_get_max_memory
func ConfigGetMaxMemory(configId uint64) uint64

//go:wasm-module lunatic::process
//go:export config_set_max_fuel
func ConfigSetMaxFuel(configId uint64, maxFuel uint64)

//go:wasm-module lunatic::process
//go:export config_get_max_fuel
func ConfigGetMaxFuel(configId uint64) uint64

//go:wasm-module lunatic::process
//go:export config_can_compile_modules
func ConfigCanCompileModules(configId uint64) uint32

//go:wasm-module lunatic::process
//go:export config_set_can_compile_modules
func ConfigSetCanCompileModules(configId uint64, can uint32)

//go:wasm-module lunatic::process
//go:export config_can_create_configs
func ConfigCanCreateConfigs(configId uint64) uint32

//go:wasm-module lunatic::process
//go:export config_set_can_create_configs
func ConfigSetCanCreateConfigs(configId uint64, can uint32)

//go:wasm-module lunatic::process
//go:export config_can_spawn_processes
func ConfigCanSpawnProcesses(configId uint64) uint32

//go:wasm-module lunatic::process
//go:export config_set_can_spawn_processes
func ConfigSetCanSpawnProcesses(configId uint64, can uint32)

//go:wasm-module lunatic::process
//go:export spawn
func Spawn(link int64, configId int64, moduleId int64, function uint32, functionLen uint32, params uint32, paramsLen uint32, id uint32) uint32

//go:wasm-module lunatic::process
//go:export sleep_ms
func SleepMs(millis uint64)

//go:wasm-module lunatic::process
//go:export die_when_link_dies
func DieWhenLinkDies(trap uint32)

//go:wasm-module lunatic::process
//go:export process_id
func ProcessId() uint64

//go:wasm-module lunatic::process
//go:export environment_id
func EnvironmentId() uint64

//go:wasm-module lunatic::process
//go:export link
func Link(tag int64, processId uint64)

//go:wasm-module lunatic::process
//go:export unlink
func Unlink(processId uint64)

//go:wasm-module lunatic::process
//go:export kill
func Kill(processId uint64)


