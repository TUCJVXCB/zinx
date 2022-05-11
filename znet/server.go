package znet

import (
	"fmt"
	"net"
)

type Server struct {
	Name string

	IP string

	Port int

	IPVersion string
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

		for {
			conn, err := listener.Accept()

			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err:", err)
						continue
					}
					fmt.Println("recv from client:", string(buf))
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err:", err)
						continue
					}
				}
			}()
		}
	}()
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
	}
	return s
}
