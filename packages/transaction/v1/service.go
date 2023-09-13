package v1

import (
	"context"
	pb "ecommerece/packages/proto/transaction"
	"ecommerece/packages/transaction/store"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	//_ "google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	store store.Store
	pb.UnimplementedTransactionServiceServer
}

func (s *server) GetTotalSales(ctx context.Context, req *pb.GetTotalSalesRequest) (*pb.GetTotalSalesResponse, error) {
	//TODO implement me
	totalPrice, err := s.store.GetTotalSales(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Error while getting total price from transaction store: %s", err))
	}

	return &pb.GetTotalSalesResponse{
		TotalPrice: *totalPrice,
	}, nil
}

func (s *server) GetSalesByProductId(ctx context.Context, req *pb.GetSalesByProductIdRequest) (*pb.GetSalesByProductIdResponse, error) {
	productId, err := uuid.Parse(req.ProductId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid customer id")
	}

	result, err := s.store.GetSalesByProductId(ctx, productId)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Error while getting total price from transaction store: %s", err))
	}

	return &pb.GetSalesByProductIdResponse{
		TotalPrice: *result,
	}, nil
}

func (s *server) GetTopFiveCustomersId(ctx context.Context, request *pb.GetTopFiveCustomersIdRequest) (*pb.GetTopFiveCustomersIdResponse, error) {

	result, err := s.store.GetTopFiveCustomersId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Error while getting total price from transaction store: %s", err))
	}
	var customers []string
	for _, customerId := range result {
		customers = append(customers, customerId.String())

	}
	return &pb.GetTopFiveCustomersIdResponse{
		CustomerId: customers,
	}, nil
}

func (s *server) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.CreateTransactionResponse, error) {
	customerId, err := uuid.Parse(req.CustomerId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid customer id")
	}

	productId, err := uuid.Parse(req.ProductId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid product id")
	}

	if !(req.Quantity > 0) {
		return nil, status.Error(codes.InvalidArgument, "invalid quantity value")
	}

	if !(req.TotalPrice > 0) {
		return nil, status.Error(codes.InvalidArgument, "invalid total price value")
	}

	transaction := store.Transaction{
		CustomerId: customerId,
		ProductId:  productId,
		Quantity:   req.Quantity,
		TotalPrice: req.TotalPrice,
	}

	result, err := s.store.InsertTransaction(ctx, transaction)
	if err != nil {
		return nil, status.Error(codes.Internal, "Ops")
	}

	return &pb.CreateTransactionResponse{
		Id:        result.Id.String(),
		IsSuccess: true,
	}, nil
}

func (s *server) GetTransactionById(ctx context.Context, req *pb.GetTransactionByIdRequest) (*pb.GetTransactionByIdResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid transaction id")

	}

	transaction, err := s.store.GetTransactionById(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Internal, "")
	}
	if transaction == nil && err == nil {
		return nil, status.Error(codes.NotFound, "no record")
	}

	return &pb.GetTransactionByIdResponse{Transaction: &pb.Transaction{
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

func (s *server) GetAllTransactions(ctx context.Context, req *pb.GetAllTransactionsRequest) (*pb.GetAllTransactionsResponse, error) {
	transactions, err := s.store.GetAllTransaction(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Ops")
	}

	if transactions == nil {
		return nil, status.Error(codes.NotFound, "no transaction found")
	}

	var transactionsList []*pb.Transaction
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

	return &pb.GetAllTransactionsResponse{
		Transaction: transactionsList,
	}, nil
}

func NewTransactionService(store store.Store) pb.TransactionServiceServer {
	return &server{
		store: store,
	}
}
