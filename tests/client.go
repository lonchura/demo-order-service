package main

import (
	"context"
	"fmt"
	"github.com/lonchura/demo-order-service/proto"
	"google.golang.org/grpc"
	"log"
)

func main()  {
	conn, err := grpc.Dial("localhost:8089", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	orderReq := &proto.OrderReq{

	}

	orderService := proto.NewOrderServiceClient(conn)
	orderResp, err := orderService.Create(context.Background(), orderReq)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", orderResp)
}
