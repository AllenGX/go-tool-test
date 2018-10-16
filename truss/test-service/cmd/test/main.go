package main

import (
	"context"
	"fmt"
	"os"
	"time"

	// This Service
	pb "test/truss/pb"
	"test/truss/test-service/svc/client/cli/handlers"
	grpcclient "test/truss/test-service/svc/client/grpc"

	"google.golang.org/grpc"
)

func main() {

	var (
		service pb.TestServer
		err     error
	)
	grpcAddr := "127.0.0.1:8888"
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while dialing grpc connection: %v", err)
	}
	defer conn.Close()
	service, err = grpcclient.New(conn)

	NameHello := "NameHello:'allen guo'"
	request, err := handlers.Hello(NameHello)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error calling handlers.Hello: %v\n", err)
	}

	v, err := service.Hello(context.Background(), request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error calling service.Hello: %v\n", err)
	}
	fmt.Println("Client Requested with:")
	fmt.Println(NameHello)
	fmt.Println("Server Responded with:")
	fmt.Println(v)
}
