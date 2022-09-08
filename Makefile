.PHONY: version
version:
	tinygo build -o version.wasm -target=wasi -wasm-abi=generic ./_examples/version
	lunatic version.wasm

.PHONY: bindings
bindings:
	wit-go --dest=internal/bindings ./wit/*
