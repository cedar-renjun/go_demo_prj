package main

import (
	"fmt"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Connect : ", conn.RemoteAddr())

	for {
		data := make([]byte, 2048)
		n, err := conn.Read(data)

		if 0 == n {
			fmt.Println("%s has disconnect", conn.RemoteAddr())
			break
		}

		if nil != err {
			fmt.Println(err)
			continue
		}

		fmt.Println("Receive data [%s] from [%s]", string(data[:n]), conn.RemoteAddr())
	}
}

func main() {

	fmt.Println("Start server")

	listener, err := net.Listen("tcp", ":8899")
	if nil != err {
		fmt.Println(err)
		return
	}

	fmt.Println("Start listen localhost :8899")
	for {
		conn, err := listener.Accept()
		if nil != err {
			fmt.Println(err)
			return
		}

		go handle(conn)
	}
}
