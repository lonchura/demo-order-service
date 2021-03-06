package main

import (
	"github.com/lonchura/demo-order-service/libs/wm"
	"github.com/lonchura/demo-order-service/proto"
	"github.com/lonchura/demo-order-service/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main()  {
	// WorkerManger start
	workerManager := wm.NewWorkerManager()
	workerManager.Start()

	// grpc server
	grpcServer := grpc.NewServer()

	// register service
	proto.RegisterOrderServiceServer(grpcServer, new(service.OrderServiceImpl))

	// listen
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatal(err)
	}

	// serve
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}

	// wait worker stop
	workerManager.Stop()
}