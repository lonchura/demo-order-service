package service

import (
	"context"
	"github.com/lonchura/demo-order-service/libs/wm"
	"github.com/lonchura/demo-order-service/proto"
	"log"
	"time"
)

type OrderServiceImpl struct {}

func (s *OrderServiceImpl) Create(ctx context.Context, in *proto.OrderReq) (resp *proto.OrderResp, err error)  {
	// log req

	// TODO mock resp
	resp = &proto.OrderResp{
		Id: 9,
	}

	// TODO send message(async)
	msg := wm.CreateMessage(resp)
	wm.WorkerManager.SendMessage(msg)

	// TODO async
	wm.WorkerManager.Async(func() {
		// TODO do something
		log.Println("do something")
		time.Sleep(time.Second * 3)
		log.Println("do something after sleep")
	})

	return
}