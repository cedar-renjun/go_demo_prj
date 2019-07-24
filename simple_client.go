package main

import "fmt"
import "net"

func main() {

	ipAddr := "192.168.1.179"
	addr := net.ParseIP(ipAddr)
	if nil == addr {
		fmt.Println("unavaliable addr")
	} else {
		fmt.Println("got ip", addr.To16())
	}

	fmt.Println("hello,world")
}
