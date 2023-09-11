package v1

import (
	"context"
	pb "ecommerece/packages/proto/analytics"
	transactionPb "ecommerece/packages/proto/transaction"
)

type server struct {
	transaction transactionPb.TransactionServiceClient
	pb.UnimplementedAnalyticsServiceServer
}

func (s server) GetTotalSales(ctx context.Context, request *pb.GetTotalSalesRequest) (*pb.GetTotalSalesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) GetSalesByProductId(ctx context.Context, request *pb.GetSalesByProductIdRequest) (*pb.GetSalesByProductIdResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) GetTopFiveCustomers(ctx context.Context, request *pb.GetTopFiveCustomersRequest) (*pb.GetTopFiveCustomersResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewAnalyticsService(transaction transactionPb.TransactionServiceClient) pb.AnalyticsServiceServer {
	return &server{
		transaction: transaction,
	}
}
