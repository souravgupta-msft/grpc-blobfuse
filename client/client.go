package client

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"grpc-blobfuse/generated/dcache"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func computeMD5Hash(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

func getStripe(ctx context.Context, client dcache.StripeServiceClient, stripeID string) (err error) {
	stripe, err := client.GetStripe(ctx, &dcache.GetStripeRequest{
		StripeID: stripeID,
	})
	if err != nil {
		fmt.Printf("error getting stripe %v : %v\n", stripeID, err)
	} else {
		fmt.Printf("Got Stripe %v, ID: %s, Offset: %d, Length: %d, Hash: %s, Data length: %v\n", stripeID, stripe.Id, stripe.Offset, stripe.Length, stripe.Hash, len(stripe.Data))
		fmt.Printf("Stripe Hash: %v\n", computeMD5Hash(stripe.Data))
	}

	return nil
}

func handleClient(client dcache.StripeServiceClient) (err error) {
	var defaultCtx = context.Background()
	var stripeSize uint64 = 16 * 1024 * 1024 //16MB

	// get stripe of 16MB
	for i := 1; i < 5; i++ {
		go func(i int) {
			err = getStripe(defaultCtx, client, fmt.Sprintf("stripeID%d-0-%d", i, stripeSize))
			if err != nil {
				fmt.Printf("error getting stripe %v : %v\n", i, err)
			} else {
				fmt.Printf("Stripe get successfully %v\n", i)
			}
		}(i)
	}

	time.Sleep(1 * time.Second)

	go func() {
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
	}()

	go func() {
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
	}()

	time.Sleep(4 * time.Second)

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
