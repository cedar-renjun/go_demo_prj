package main

import (
	"flag"
	"fmt"
	"net"
	"runtime"
	//"os"
)

// ------ job & worker arch ------

type Job interface {
	Do()
}

type Worker struct {
	JobQueue chan Job
}

func NewWorker() Worker {
	return Worker{JobQueue: make(chan Job)}
}

func (w Worker) Run(wq chan chan Job) {
	go func() {
		for {
			wq <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				job.Do()
			}
		}
	}()
}

type WorkerPool struct {
	workerlen   int
	JobQueue    chan Job
	WorkerQueue chan chan Job
}

func NewWorkerPool(workerlen int) *WorkerPool {
	return &WorkerPool{
		workerlen:   workerlen,
		JobQueue:    make(chan Job),
		WorkerQueue: make(chan chan Job, workerlen),
	}
}

func (wp *WorkerPool) Run() {
	fmt.Println("init worker")
	for i := 0; i < wp.workerlen; i++ {
		worker := NewWorker()
		worker.Run(wp.WorkerQueue)
	}

	go func() {
		for {
			select {
			case job := <-wp.JobQueue:
				worker := <-wp.WorkerQueue
				worker <- job
			}
		}
	}()
}

func tcpHandle(conn net.Conn) {
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

//-----
type mmmConn struct {
	idx   int
	mconn net.Conn
}

func (c *mmmConn) Do() {
	tcpHandle(c.mconn)
}

func main() {

	port := flag.String("port", ":8888", "tcp listen port")

	flag.Parse()
	fmt.Println("Start server port: ", *port)

	runtime.GOMAXPROCS(runtime.NumCPU())

	max_worker := 100 * 100 * 20
	p := NewWorkerPool(max_worker)
	p.Run()

	cnt := 0

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

		tc := &mmmConn{idx: cnt, mconn: conn}
		p.JobQueue <- tc
		cnt++
	}
}
