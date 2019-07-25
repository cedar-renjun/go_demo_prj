package main

import (
	"fmt"
	"net"
	"flag"
	//"os"
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

		conn.Write(data)

		fmt.Printf("Receive data [%s] from [%s]", string(data[:n]), conn.RemoteAddr())
	}
}

func main() {

	port := flag.String("port", ":8888", "tcp listen port")

	flag.Parse()

	fmt.Println("Start server port: ", *port)

	listener, err := net.Listen("tcp", *port)
	if nil != err {
		fmt.Println(err)
		return
	}

	fmt.Println("Start listen localhost", *port)
	for {
		conn, err := listener.Accept()
		if nil != err {
			fmt.Println(err)
			return
		}

		go handle(conn)
	}
}
