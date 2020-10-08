package service

import (
	"context"
	"github.com/lonchura/demo-order-service/proto"
)

type OrderServiceImpl struct {}

func (s *OrderServiceImpl) Create(ctx context.Context, in *proto.OrderReq) (resp *proto.OrderResp, err error)  {
	// log req

	// TODO mock resp
	resp = &proto.OrderResp{
		Id: 9,
	}

	return
}