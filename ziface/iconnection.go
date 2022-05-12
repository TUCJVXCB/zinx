package ziface

import "net"

type IConnection interface {
	Start()

	Stop()

	GetTcpConnection() *net.TCPConn

	GetConnId() uint32

	RemoteAddr() net.Addr

	Send(data []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error
