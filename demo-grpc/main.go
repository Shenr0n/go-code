package main

import (
	"context"
	"log"
	"net"

	"github.com/Shenr0n/go-code/demo-grpc/invoicer"
	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte(req.From),
		Docx: []byte("test"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("Cannot create listener %s", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &myInvoicerServer{}

	invoicer.RegisterInvoicerServer(serverRegistrar, service)

	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("Cannot server %s", err)
	}
}
