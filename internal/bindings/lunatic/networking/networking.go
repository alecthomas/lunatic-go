package networking

//go:wasm-module lunatic::networking
//go:export resolve
func Resolve(nameStr uint32, nameStrLen uint32, timeout uint32, id uint64) uint32

//go:wasm-module lunatic::networking
//go:export drop_dns_iterator
func DropDnsIterator(dnsIterId uint64)

//go:wasm-module lunatic::networking
//go:export resolve_next
func ResolveNext(dnsIterId uint64, addrType uint32, addr uint32, port uint16, flowInfo uint32, scopeId uint32) uint32

//go:wasm-module lunatic::networking
//go:export tcp_bind
func TcpBind(addrType uint32, addr uint32, port uint32, flowInfo uint32, scopeId uint32, id uint64) uint32

//go:wasm-module lunatic::networking
//go:export udp_bind
func UdpBind(addrType uint32, addr uint32, port uint32, flowInfo uint32, scopeId uint32, id uint64) uint32

//go:wasm-module lunatic::networking
//go:export drop_tcp_listener
func DropTcpListener(tcpListenerId uint64)

//go:wasm-module lunatic::networking
//go:export drop_udp_socket
func DropUdpSocket(udpSocketId uint64)

//go:wasm-module lunatic::networking
//go:export tcp_local_addr
func TcpLocalAddr(tcpListenerId uint64, addrDnsIter uint64) uint32

//go:wasm-module lunatic::networking
//go:export udp_local_addr
func UdpLocalAddr(udpSocketId uint64, addrDnsIter uint64) uint32

//go:wasm-module lunatic::networking
//go:export tcp_accept
func TcpAccept(listenerId uint64, id uint64, peerDnsIter uint64) uint32

//go:wasm-module lunatic::networking
//go:export tcp_connect
func TcpConnect(addrType uint32, addr uint32, port uint32, flowInfo uint32, scopeId uint32, timeout uint32, id uint64) uint32

//go:wasm-module lunatic::networking
//go:export drop_tcp_stream
func DropTcpStream(tcpStreamId uint64)

//go:wasm-module lunatic::networking
//go:export clone_tcp_stream
func CloneTcpStream(tcpStreamId uint64) uint64

//go:wasm-module lunatic::networking
//go:export tcp_write_vectored
func TcpWriteVectored(tcpStreamId uint64, ciovecArray uint32, ciovecArrayLen uint32, timeout uint32, opaque uint64) uint32

//go:wasm-module lunatic::networking
//go:export tcp_read
func TcpRead(tcpStreamId uint64, buffer uint32, bufferLen uint32, timeout uint32, opaque uint64) uint32

//go:wasm-module lunatic::networking
//go:export tcp_flush
func TcpFlush(tcpStreamId uint64, errorId uint64) uint32


