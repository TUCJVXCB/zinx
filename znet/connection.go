package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Connection struct {
	Conn *net.TCPConn

	ConnID uint32

	isClosed bool

	Router ziface.IRouter

	ExitChan chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		Router:   router,
		ExitChan: make(chan bool, 1),
	}
	return c
}

func (c *Connection) Start() {
	fmt.Println("Conn Start()...ConnID = ", c.ConnID)
	go c.StartReader()
}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop()...ConnID = ", c.ConnID)

	if c.isClosed {
		return
	}
	c.isClosed = true

	c.Conn.Close()

	close(c.ExitChan)
}

func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnId() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	_, err := c.Conn.Write(data)
	return err
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("ConnID = ", c.ConnID, "Reader is exit, remote addr is "+c.RemoteAddr().String())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			c.ExitChan <- true
			continue
		}

		req := Request{
			connection: c,
			data:       buf,
		}

		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

	}
}
