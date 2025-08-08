package ports

import (
	"context"
	"github.com/2u35-Wooje/gorder-v2/common/genproto/orderpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCServer struct {
}

func (G GRPCServer) CreateOrder(ctx context.Context, request *orderpb.CreateOrderRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (G GRPCServer) GetOrder(ctx context.Context, request *orderpb.GetOrderRequest) (*orderpb.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (G GRPCServer) UpdateOrder(ctx context.Context, order *orderpb.Order) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func NewGrpcServer() *GRPCServer {
	return &GRPCServer{}
}
