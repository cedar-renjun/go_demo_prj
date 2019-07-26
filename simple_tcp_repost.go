package main

import (
	"fmt"
	"io"
	"net"
	"runtime"
	"sync"
)

var server_addrs = []string{":8890", ":8891", ":8892", ":8893"}

//var server_addrs = []string{":8890", ":8890", ":8890", ":8890"}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("repost server start")
	forword()
}

func forword() {
	cnt := 0
	lis, err := net.Listen("tcp", ":8898")
	if err != nil {
		fmt.Println(err)
	}
	defer lis.Close()

	fmt.Println("listening port 8898")

	for {
		localConn, err := lis.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handle(localConn, cnt)
		cnt++
	}
}

func handle(localConn net.Conn, cnt int) {
	var wg sync.WaitGroup

	i_port := server_addrs[cnt%4]

	remoteConn, err := net.Dial("tcp", i_port)
	if err != nil {
		localConn.Close()
		fmt.Println(err)
	}

	fmt.Printf("[%08d] repost msg\r\n", cnt)

	wg.Add(2)
	go func(local net.Conn, remote net.Conn) {
		defer wg.Done()
		io.Copy(remote, local)
		remote.Close()
	}(localConn, remoteConn)
	go func(local net.Conn, remote net.Conn) {
		defer wg.Done()
		io.Copy(local, remote)
		local.Close()
	}(localConn, remoteConn)
	wg.Wait()

	// this
}
