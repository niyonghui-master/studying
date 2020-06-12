package main

import (
	"NYH/Day20191224/rpc"
	"log"
)

func main()  {
	client, err := rpc.DailHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(reply)
}
