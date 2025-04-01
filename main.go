package main

import (
	"flag"
	"fmt"
	"grpc-blobfuse/client"
	"grpc-blobfuse/server"
)

func main() {
	isServer := flag.Bool("server", false, "Run server")
	addr := flag.String("addr", "localhost:9090", "Address to listen to")
	secure := flag.Bool("secure", false, "Use tls secure transport")

	flag.Parse()

	if *isServer {
		if err := server.RunServer(*addr, *secure); err != nil {
			fmt.Println("error running server:", err)
		}
	} else {
		if err := client.RunClient(*addr, *secure); err != nil {
			fmt.Println("error running client:", err)
		}
	}
}
