package service

import (
	"context"

	"app.io/data/model"
	pb "app.io/service/proto_service"
)

func NewTransactionService(transactionGrpcClient pb.BlockServiceClient, ctx context.Context) TransactionService {
	return &transactionServiceImpl{
		ctx:                   ctx,
		transactionGrpcClient: transactionGrpcClient,
	}
}

type transactionServiceImpl struct {
	ctx                   context.Context
	transactionGrpcClient pb.BlockServiceClient
}

func (service *transactionServiceImpl) Single(request model.CreateTransactionRequest) (response model.CreateTransactionResponse) {
	// validation.Validate(request)
	service.transactionGrpcClient.GetBlock(service.ctx, &pb.BlockRequest{
		BlockNumber: 124454,
	})

	response = model.CreateTransactionResponse{
		Id:       "1",
		Name:     "transaction.Name",
		Price:    10,
		Quantity: 10,
	}
	return response
}

func (service *transactionServiceImpl) List() (responses []model.GetTransactionResponse) {
	item := model.GetTransactionResponse{
		Id:       "1",
		Name:     "transaction.Name",
		Price:    10,
		Quantity: 10,
	}
	responses = append(responses, item)
	return
}
