package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("Client Start...")

	conn, err := net.Dial("tcp4", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("connect server err:", err)
		return
	}

	for {
		_, err := conn.Write([]byte("hello zinx..."))
		if err != nil {
			fmt.Println("write err:", err)
			continue
		}

		buf := make([]byte, 512)
		count, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read from server err:", err)
			return
		}

		fmt.Printf("echo from server:%s, count:%d\n", string(buf), count)

		time.Sleep(1 * time.Second)
	}
}
