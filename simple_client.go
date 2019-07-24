package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("client start!")

	conn, err := net.Dial("tcp", ":8899")
	if nil != err {
		fmt.Println(err)
		return
	}

	fmt.Println("Connect to localhost:8899 success")
	defer conn.Close()

	str := "hello,world"
	conn.Write([]byte(str))

	fmt.Println("client finish")

	return
}
