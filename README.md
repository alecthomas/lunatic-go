# TinyGo bindings for [Lunatic](https://lunatic.solutions/)

## Progress

- [x] Raw bindings
- [ ] Semantic layer
  - [x] Version
  - [ ] Error
  - [ ] Message
  - [ ] Networking
  - [ ] Process
  - [ ] Registry
  - [ ] Timer
  - [ ] Distributed

## Developing

All of the tools needed to develop on this project are included in this repo
using [Hermit](https://cashapp.github.io/hermit). The easiest way to get started
is to activate the project and type `make`, which should build and run the
`version` example:

```
$ . ./bin/activate-hermit
$ make
tinygo build --no-debug -o version.wasm -target=wasi -wasm-abi=generic ./cmd/version
lunatic version.wasm
tinygo 0.25.0 wasm linux
Lunatic 0.10.0
```

Raw WASM bindings generated by [wit-go](https://github.com/alecthomas/wit-go)
from `wit/*.wit` are in `internal/bindings/`.  The
[WIT](https://github.com/WebAssembly/component-model/blob/main/design/mvp/WIT.md)
files themselves are originally from this [pull
request](https://github.com/lunatic-solutions/lunatic-rs/pull/31).
