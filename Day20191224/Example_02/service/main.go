package main

import (
	"NYH/Day20191224/rpc"
	"log"
	"net"
	rpc2 "net/rpc"
)

type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main()  {
	rpc.RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc2.ServeConn(conn)
	}
}
