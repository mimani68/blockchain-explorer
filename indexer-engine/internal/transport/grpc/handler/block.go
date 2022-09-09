package handler

import (
	"context"

	service "app.io/internal/transport/grpc/service"
)

type blockHandler struct {
	srv service.BlockServiceServer
}

func NewUserHandler(a service.BlockServiceServer) service.BlockServiceServer {
	return &blockHandler{
		srv: a,
	}
}

func (s *blockHandler) StoreBlockation(ctx context.Context, in *service.BlockRequest) (*service.BlockResponse, error) {
	return s.srv.StoreBlockation(ctx, in)
}

func (s *blockHandler) ShowBlockInBlock(ctx context.Context, in *service.BlockRequest) (*service.BlockResponse, error) {
	return s.srv.ShowBlockInBlock(ctx, in)
}

func (s *blockHandler) Estimate(ctx context.Context, in *service.BlockRequest) (*service.EstimateResponse, error) {
	return s.srv.Estimate(ctx, in)
}
