package v1

import (
	"context"
	pb "ecommerece/packages/proto/transaction"
	"ecommerece/packages/transaction/store"
	"errors"
	"github.com/google/uuid"
	_ "google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (s *server) GetTransaction(ctx context.Context, req *pb.GetTransactionRequest) (*pb.GetTransactionResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errors.New("invalid transaction id")
	}

	transaction := s.GetTransactionById(ctx, id)
	if transaction == nil {
		return nil, errors.New("no transaction found")
	}

	return &pb.GetTransactionResponse{Transaction: &pb.Transaction{
		Id:         transaction.Id.String(),
		CreatedAt:  timestamppb.New(transaction.CreatedAt),
		CustomerId: transaction.CustomerId.String(),
		ProductId:  transaction.ProductId.String(),
		Quantity:   transaction.Quantity,
		TotalPrice: transaction.TotalPrice,
	}}, nil

}
func (s *server) StreamTransactions(*pb.StreamTransactionsRequest, pb.TransactionService_StreamTransactionsServer) error {
	//todo later
	return nil

}
func (s *server) GetAllTransactions(ctx context.Context, req *pb.GetTransactionRequest) (*pb.GetAllTransactionsResponse, error) {
	transactions, err := s.GetAllTransaction(ctx)
	if err != nil {
		return nil, err
	}

	if transactions == nil {
		return nil, errors.New("no transaction found")
	}

	transactionsList := []*pb.Transaction{}
	for _, transaction := range transactions {
		transactionsList = append(transactionsList, &pb.Transaction{
			Id:         transaction.Id.String(),
			CreatedAt:  timestamppb.New(transaction.CreatedAt),
			CustomerId: transaction.CustomerId.String(),
			ProductId:  transaction.ProductId.String(),
			Quantity:   transaction.Quantity,
			TotalPrice: transaction.TotalPrice,
		})
	}

	return &pb.GetAllTransactionsResponse{Transaction: transactionsList}, nil
}
