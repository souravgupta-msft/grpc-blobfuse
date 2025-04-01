package client

import (
	"context"
	"fmt"
	"grpc-blobfuse/generated/dcache"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func handleClient(client dcache.StripeServiceClient) (err error) {
	var defaultCtx = context.Background()
	var stripeSize uint64 = 16 * 1024 * 1024 //16MB

	// get stripe of 16MB
	stripe, err := client.GetStripe(defaultCtx, &dcache.GetStripeRequest{
		StripeID: fmt.Sprintf("stripeID1-0-%d", stripeSize),
	})
	if err != nil {
		fmt.Println("error getting stripe:", err)
	} else {
		fmt.Printf("Got Stripe, ID: %s, Offset: %d, Length: %d, Hash: %s, Data length: %v\n", stripe.Id, stripe.Offset, stripe.Length, stripe.Hash, len(stripe.Data))
	}

	// go func() {
	// put stripe
	_, err = client.PutStripe(defaultCtx, &dcache.Stripe{
		Id:     "stripeID1",
		Offset: stripeSize,
		Length: stripeSize,
		Hash:   "stripeHash1",
		Data:   make([]byte, stripeSize),
	})
	if err != nil {
		fmt.Println("error putting stripe 1:", err)
	} else {
		fmt.Println("Stripe put successfully 1")
	}
	// }()

	// go func() {
	_, err = client.PutStripe(defaultCtx, &dcache.Stripe{
		Id:     "stripeID1",
		Offset: stripeSize,
		Length: stripeSize,
		Hash:   "stripeHash2",
		Data:   make([]byte, stripeSize),
	})
	if err != nil {
		fmt.Println("error putting stripe 2:", err)
	} else {
		fmt.Println("Stripe put successfully 2")
	}
	// }()

	// time.Sleep(2 * time.Second)

	// remove stripe
	_, err = client.RemoveStripe(defaultCtx, &dcache.RemoveStripeRequest{
		StripeID: fmt.Sprintf("stripeID1-%d-%d", stripeSize, stripeSize),
	})
	if err != nil {
		fmt.Println("error removing stripe:", err)
	} else {
		fmt.Println("Stripe removed successfully")
	}

	return nil
}

func RunClient(addr string, secure bool) error {
	var opts []grpc.DialOption
	if secure {
		caFile := "x509/ca_cert.pem"      // Path to CA certificate
		serverHostOverride := "localhost" // Server name override for TLS
		creds, err := credentials.NewClientTLSFromFile(caFile, serverHostOverride)
		if err != nil {
			fmt.Printf("Failed to create TLS credentials: %v\n", err)
			return err
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	// update max message size to 64 MB, default being 4 MB
	opts = append(opts, grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(64*1024*1024),
		grpc.MaxCallSendMsgSize(64*1024*1024),
	))

	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		fmt.Printf("fail to dial: %v\n", err)
		return err
	}
	defer conn.Close()
	client := dcache.NewStripeServiceClient(conn)

	// check connection
	_, err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Printf("Ping failed: %v\n", err)
		return err
	}

	return handleClient(client)
}
