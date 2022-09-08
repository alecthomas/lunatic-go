package networking

//go:wasm-module lunatic::networking
//go:export resolve
func Resolve(nameStr uint32, nameStrLen uint32, timeout uint32, id uint64) uint32

//go:wasm-module lunatic::networking
//go:export drop-dns-iterator
func DropDnsIterator(dnsIterId uint64)

//go:wasm-module lunatic::networking
//go:export resolve-next
func ResolveNext(dnsIterId uint64, addrType uint32, addr uint32, port uint16, flowInfo uint32, scopeId uint32) uint32

//go:wasm-module lunatic::networking
//go:export tcp-bind
func TcpBind(addrType uint32, addr uint32, port uint32, flowInfo uint32, scopeId uint32, id uint64) uint32

//go:wasm-module lunatic::networking
//go:export udp-bind
func UdpBind(addrType uint32, addr uint32, port uint32, flowInfo uint32, scopeId uint32, id uint64) uint32

//go:wasm-module lunatic::networking
//go:export drop-tcp-listener
func DropTcpListener(tcpListenerId uint64)

//go:wasm-module lunatic::networking
//go:export drop-udp-socket
func DropUdpSocket(udpSocketId uint64)

//go:wasm-module lunatic::networking
//go:export tcp-local-addr
func TcpLocalAddr(tcpListenerId uint64, addrDnsIter uint64) uint32

//go:wasm-module lunatic::networking
//go:export udp-local-addr
func UdpLocalAddr(udpSocketId uint64, addrDnsIter uint64) uint32

//go:wasm-module lunatic::networking
//go:export tcp-accept
func TcpAccept(listenerId uint64, id uint64, peerDnsIter uint64) uint32

//go:wasm-module lunatic::networking
//go:export tcp-connect
func TcpConnect(addrType uint32, addr uint32, port uint32, flowInfo uint32, scopeId uint32, timeout uint32, id uint64) uint32

//go:wasm-module lunatic::networking
//go:export drop-tcp-stream
func DropTcpStream(tcpStreamId uint64)

//go:wasm-module lunatic::networking
//go:export clone-tcp-stream
func CloneTcpStream(tcpStreamId uint64) uint64

//go:wasm-module lunatic::networking
//go:export tcp-write-vectored
func TcpWriteVectored(tcpStreamId uint64, ciovecArray uint32, ciovecArrayLen uint32, timeout uint32, opaque uint64) uint32

//go:wasm-module lunatic::networking
//go:export tcp-read
func TcpRead(tcpStreamId uint64, buffer uint32, bufferLen uint32, timeout uint32, opaque uint64) uint32

//go:wasm-module lunatic::networking
//go:export tcp-flush
func TcpFlush(tcpStreamId uint64, errorId uint64) uint32


