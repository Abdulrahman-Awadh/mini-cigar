package v1

import (
	"context"
	pb "ecommerece/packages/proto/analytics"
	transactionPb "ecommerece/packages/proto/transaction"
	transactionpb "ecommerece/packages/proto/transaction"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	transaction transactionPb.TransactionServiceServer
	pb.UnimplementedAnalyticsServiceServer
}

func (s server) BringTotalSales(ctx context.Context, request *pb.BringTotalSalesRequest) (*pb.BringTotalSalesResponse, error) {
	totalPrice, err := s.transaction.GetTotalSales(ctx, &transactionpb.GetTotalSalesRequest{})
	if err != nil {
		return nil, status.Error(codes.Internal, "Error while getting total sales")
	}

	return &pb.BringTotalSalesResponse{TotalPrice: totalPrice.TotalPrice}, nil
}

func (s server) BringSalesByProductId(ctx context.Context, request *pb.BringSalesByProductIdRequest) (*pb.BringSalesByProductIdResponse, error) {
	productId, err := uuid.Parse(request.ProductId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid product id")
	}
	result, err := s.transaction.GetSalesByProductId(ctx, &transactionpb.GetSalesByProductIdRequest{ProductId: productId.String()})
	if err != nil {
		return nil, status.Error(codes.Internal, "Error while getting total sales")
	}

	return &pb.BringSalesByProductIdResponse{TotalPrice: result.TotalPrice}, nil
}

func (s server) ListTopFiveCustomers(ctx context.Context, request *pb.ListTopFiveCustomersRequest) (*pb.ListTopFiveCustomersResponse, error) {
	//TODO implement me
	transactions, err := s.transaction.GetTopFiveCustomersId(ctx, &transactionPb.GetTopFiveCustomersIdRequest{})
	if err != nil {
		return nil, status.Error(codes.Internal, "Error while getting top customers")
	}

	return &pb.ListTopFiveCustomersResponse{CustomerId: transactions.CustomerId}, nil
}

func NewAnalyticsService(transaction transactionPb.TransactionServiceServer) pb.AnalyticsServiceServer {
	return &server{
		transaction: transaction,
	}
}
