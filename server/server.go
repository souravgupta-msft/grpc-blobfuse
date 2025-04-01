package server

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"grpc-blobfuse/generated/dcache"
)

func RunServer(addr string, secure bool) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return err
	}

	var opts []grpc.ServerOption
	if secure {
		certFile := "x509/server_cert.pem" // Path to server certificate
		keyFile := "x509/server_key.pem"   // Path to server key
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			fmt.Printf("Failed to generate credentials: %v\n", err)
			return err
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	// update max message size to 64 MB, default being 4 MB
	opts = append(opts, grpc.MaxRecvMsgSize(64*1024*1024), grpc.MaxSendMsgSize(64*1024*1024)) // 64 MB

	grpcServer := grpc.NewServer(opts...)

	fmt.Println("Starting the simple server... on ", addr)
	dcache.RegisterStripeServiceServer(grpcServer, NewStripeServiceHandler())
	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

	return err
}
