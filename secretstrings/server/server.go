package main

import (
	"flag"
	"math/rand"
	"net"
	"net/rpc"
	"uk.ac.bris.cs/distributed2/secretstrings/stubs"

	//	"errors"
	//	"flag"
	//	"fmt"
	//	"net"
	"time"
	//"uk.ac.bris.cs/distributed2/secretstrings/stubs"
	//	"net/rpc"
)

/** Super-Secret `reversing a string' method we can't allow clients to see. **/
func ReverseString(s string, i int) string {
	time.Sleep(time.Duration(rand.Intn(i)) * time.Second)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

type SecretStringOperations struct{}

func (s *SecretStringOperations) Reverse(req stubs.Request, response *stubs.Response) (err error) {

	response.Message = ReverseString(req.Message, 10)
	return
}

// terrible delay , new fast server
func (s *SecretStringOperations) FastReverse(req stubs.Request, response *stubs.Response) (err error) {
	response.Message = ReverseString(req.Message, 2)
	return

}

func main() {
	//we want the server to accpt some config
	pAddr := flag.String("port", "8030", "Port to listen on")
	flag.Parse()
	//part of the reverse string impl
	rand.Seed(time.Now().UnixNano())
	// use the net/rpc package to register service
	rpc.Register(&SecretStringOperations{})
	listener, _ := net.Listen("tcp", ":"+*pAddr)
	//defer means delay the execution of a function or a statement until the nearby function returns.
	//In simple words, defer will move the execution of the statement to the very end inside a function.
	defer listener.Close()
	// we want the service to start accepting the communications
	rpc.Accept(listener)

}
