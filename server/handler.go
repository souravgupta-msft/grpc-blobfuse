package server

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"grpc-blobfuse/generated/dcache"

	"github.com/golang/protobuf/ptypes/empty"
)

// type check to ensure that StripeServiceHandler implements dcache.StripeServiceServer interface
var _ dcache.StripeServiceServer = &StripeServiceHandler{}

type StripeServiceHandler struct {
	dcache.UnimplementedStripeServiceServer
	mu       sync.Mutex
	cacheDir string
}

func NewStripeServiceHandler() *StripeServiceHandler {
	return &StripeServiceHandler{
		cacheDir: "/home/sourav/dcache",
	}
}

func (h *StripeServiceHandler) Ping(ctx context.Context, opts *empty.Empty) (*empty.Empty, error) {
	fmt.Println("Ping!")
	return nil, nil
}

func (h *StripeServiceHandler) GetStripe(ctx context.Context, opts *dcache.GetStripeRequest) (*dcache.Stripe, error) {
	if opts == nil {
		return nil, fmt.Errorf("GetStripe: GetStripeRequest is nil")
	}
	fmt.Printf("GetStripe called for stripe ID %v\n", opts.StripeID)

	stripeFilePath := filepath.Join(h.cacheDir, opts.StripeID)
	data, err := os.ReadFile(stripeFilePath)
	if err != nil {
		fmt.Printf("Error reading stripe file [%v]\n", err.Error())
		return nil, err
	}

	return &dcache.Stripe{
		Id: opts.StripeID,
		// Offset: offset,
		// Length: stripeLength,
		Hash: "stripeHash",
		Data: data,
	}, nil
}

func (h *StripeServiceHandler) PutStripe(ctx context.Context, stripe *dcache.Stripe) (*empty.Empty, error) {
	if stripe == nil {
		return nil, fmt.Errorf("PutStripe: Stripe is nil")
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	// should be written once, take locks
	fmt.Printf("PutStripe called for stripe ID %v, offset %v, stripe length %v, hash %v, data length %v\n",
		stripe.Id, stripe.Offset, stripe.Length, stripe.Hash, len(stripe.Data))

	stripeFilePath := filepath.Join(h.cacheDir, fmt.Sprintf("%s-%d-%d", stripe.Id, stripe.Offset, stripe.Length))
	err := os.WriteFile(stripeFilePath, stripe.Data, 0400)
	if err != nil {
		fmt.Printf("Error writing stripe file, hash %v [%v]\n", stripe.Hash, err.Error())
		return nil, err
	}

	fmt.Printf("PutStripe: Stripe written successfully, hash: %v\n", stripe.Hash)

	return nil, nil
}

func (h *StripeServiceHandler) RemoveStripe(ctx context.Context, opts *dcache.RemoveStripeRequest) (*empty.Empty, error) {
	if opts == nil {
		return nil, fmt.Errorf("RemoveStripe: RemoveStripeRequest is nil")
	}

	fmt.Printf("RemoveStripe called for stripe ID %v\n", opts.StripeID)

	stripeFilePath := filepath.Join(h.cacheDir, opts.StripeID)
	err := os.Remove(stripeFilePath)
	if err != nil {
		fmt.Printf("Error removing stripe file [%v]\n", err.Error())
		return nil, err
	}

	return nil, nil
}
