package v1

import (
	"context"
	pb "ecommerece/packages/proto/transaction"
	"ecommerece/packages/transaction/store"
	"errors"
	"github.com/google/uuid"
	_ "google.golang.org/grpc"
)

type server struct {
	store.Store
}

func NewTransactionService(store store.Store) *server {
	return &server{store}
}

func (s *server) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.CreateTransactionResponse, error) {
	customerId, err := uuid.Parse(req.CustomerId)
	if err != nil {
		return nil, errors.New("invalid customer id")
	}

	productId, err := uuid.Parse(req.CustomerId)
	if err != nil {
		return nil, errors.New("invalid product id")
	}

	if !(req.Quantity > 0) {
		return nil, errors.New("invalid quantity value")
	}

	if !(req.TotalPrice > 0) {
		return nil, errors.New("invalid total price value")
	}

	transaction := store.Transaction{
		CustomerId: customerId,
		ProductId:  productId,
		Quantity:   req.Quantity,
		TotalPrice: req.TotalPrice,
	}

	result, err := s.InsertTransaction(ctx, transaction)
	if err != nil {
		return nil, errors.New("unable to add the transaction")
	}

	return &pb.CreateTransactionResponse{
		Id:        result.Id.String(),
		IsSuccess: true,
	}, nil

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
