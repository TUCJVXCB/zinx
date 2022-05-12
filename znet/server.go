package znet

import (
	"errors"
	"fmt"
	"net"
	"zinx/ziface"
)

type Server struct {
	Name string

	IP string

	Port int

	IPVersion string

	Router ziface.IRouter
}

func (s *Server) Start() {
	fmt.Printf("[start] Server Listener at IP: %s, Port: %d, is starting\n", s.IP, s.Port)

	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, "err", err)
			return
		}

		fmt.Println("start Zinx server succ, ", s.Name, "succ, listening...")
		var cid uint32
		cid = 0

		for {
			conn, err := listener.AcceptTCP()

			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			go dealConn.Start()
		}
	}()
}

func CallBackToClient(conn *net.TCPConn, bytes []byte, i int) error {
	fmt.Println("[Conn Handle] CallbackToClient...")
	if _, err := conn.Write(bytes); err != nil {
		fmt.Println("write back buf err", err)
		return errors.New("CallbackToClient error")
	}
	return nil
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()
	select {}
}

func NewServe(name string) *Server {
	s := &Server{
		Name:      name,
		IP:        "0.0.0.0",
		Port:      8080,
		IPVersion: "tcp4",
		Router:    nil,
	}
	return s
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router

	fmt.Println("Add Router succ! ")
}
