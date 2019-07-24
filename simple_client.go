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

	defer conn.Close()

	str := "hello,world"
	conn.Write([]byte(str))
	fmt.Printf("[%04d] >>> %s\r\n", i, str)

	data := make([]byte, 2048)
	n, err := conn.Read(data)
	if nil != err {
		fmt.Println(err)
		return nil
	}
	fmt.Printf("       <<< %s\r\n\r\n", string(data[:n]))

	return conn
}

func main() {

	var s1 []net.Conn

	fmt.Println("client start!")

	for i := 0; i < 1000; i++ {
		//time.Sleep(time.Second * 1)
		conn := establishConn(i)
		if nil != conn {
			s1 = append(s1, conn)
		}
	}

	fmt.Println("client finish")

	return
}
