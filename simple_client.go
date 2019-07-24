package main

import (
	"fmt"
	"net"
)

func establishConn(i int) net.Conn {
	conn, err := net.Dial("tcp", ":8899")
	if nil != err {
		fmt.Println(err)
		return nil
	}

	fmt.Printf("[%d] Connect to localhost:8899 success\r\n", i)
	defer conn.Close()

	str := "hello,world"
	conn.Write([]byte(str))

	return conn
}

func main() {

	var s1 []net.Conn

	fmt.Println("client start!")

	for i := 0; i < 1000; i++ {
		conn := establishConn(i)
		if nil != conn {
			s1 = append(s1, conn)
		}
	}

	fmt.Println("client finish")

	return
}
