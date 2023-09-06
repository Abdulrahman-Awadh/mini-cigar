package v1

import (
	"context"
	pb "ecommerece/packages/proto/transaction"
	"ecommerece/packages/transaction/store"
	_ "google.golang.org/grpc"
)

type server struct {
	store.Store
}

func NewTransactionService(store store.Store) *server {
	return &server{store}
}

func (s *server) CreateTransaction(context.Context, *pb.CreateTransactionRequest) (*pb.CreateTransactionResponse, error) {
	return nil, nil

}
func (s *server) GetTransaction(context.Context, *pb.GetTransactionRequest) (*pb.GetTransactionResponse, error) {
	return nil, nil

}
func (s *server) StreamTransactions(*pb.StreamTransactionsRequest, pb.TransactionService_StreamTransactionsServer) error {
	return nil

}
func (s *server) GetAllTransactions(context.Context, *pb.GetTransactionRequest) (*pb.GetTransactionResponse, error) {
	return nil, nil
}
