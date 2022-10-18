package main

import (
	//	"net/rpc"
	"flag"
	"net/rpc"
	"os"
	"uk.ac.bris.cs/distributed2/secretstrings/stubs"

	//	"bufio"
	//	"uk.ac.bris.cs/distributed2/secretstrings/stubs"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	server := flag.String("server", "127.0.0.1:8030", "IP:port string to connect to as server")
	flag.Parse()
	fmt.Println("Server: ", *server)
	//TODO: connect to the RPC server and send the request(s)
	client, _ := rpc.Dial("tcp", *server)
	defer client.Close()
	// this is how you open a file in go

	data, err := os.ReadFile("wordlist")
	check(err)
	str := string(data)
	request := stubs.Request{Message: str}
	response := new(stubs.Response)
	// TODO: connect to the RPC server and send the requests
	client.Call(stubs.PremiumReverseHandler, request, response)
	fmt.Println("Responded:" + response.Message)

}
