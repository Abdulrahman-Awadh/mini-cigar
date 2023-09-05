package transaction

import (
	"context"
	pb "ecommerece/packages/proto/transaction"
	"ecommerece/packages/transaction/store"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc"
)

type Transaction interface {
	CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest, opts ...grpc.CallOption) (*pb.CreateTransactionResponse, error)
	GetTransaction(ctx context.Context, in *pb.GetTransactionRequest, opts ...grpc.CallOption) (*pb.GetTransactionResponse, error)
}

type server struct {
	store.Store
}

func NewTransactionService(store store.Store) *server {
	return &server{store}
}

func (s *server) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest, opts ...grpc.CallOption) (*pb.CreateTransactionResponse, error) {

	customerId, err := uuid.Parse(req.CustomerId)
	if err != nil {
		return nil, err
	}

	productId, err := uuid.Parse(req.ProductId)
	if err != nil {
		return nil, err
	}

	transaction, err := s.InsertTransaction(ctx, store.Transaction{
		CustomerId: customerId,
		ProductId:  productId,
		Quantity:   req.Quantity,
		TotalPrice: req.TotalPrice,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateTransactionResponse{IsSuccess: true, Id: transaction.Id.String()}, nil
}

//func (s *server) GetTransaction(ctx context.Context, in *pb.GetTransactionRequest, opts ...grpc.CallOption) (*pb.GetTransactionResponse, error) {
//	return nil, nil
//}
//
//func (s *server) StreamTransactions(ctx context.Context, in *pb.StreamTransactionsRequest, opts ...grpc.CallOption) (pb.TransactionService_StreamTransactionsClient, error) {
//	return nil, nil
//}
