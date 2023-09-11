package v1_test

import (
	"context"
	pb "ecommerece/packages/proto/transaction"
	"ecommerece/packages/transaction/store"
	_ "ecommerece/packages/transaction/store/mock"
	store_mock "ecommerece/packages/transaction/store/mock"
	svc "ecommerece/packages/transaction/v1"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

type ServiceTestSuite struct {
	suite.Suite
	svc  pb.TransactionServiceServer
	mock *store_mock.Store
	ctx  context.Context
}

func (suite *ServiceTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	suite.mock = &store_mock.Store{}
	suite.svc = svc.NewTransactionService(suite.mock)
}

func (suite *ServiceTestSuite) TearDownSuite() {

}

// =================================
// Test: CreateTransaction
// =================================

func (suite *ServiceTestSuite) TestTransactionService_CreateTransaction_SUCCESS() {
	id := uuid.New()
	createdAt := time.Now()
	customerId := uuid.New()
	productId := uuid.New()

	request := pb.CreateTransactionRequest{
		CustomerId: customerId.String(),
		ProductId:  productId.String(),
		Quantity:   1,
		TotalPrice: 20,
	}

	suite.mock.
		On("InsertTransaction", suite.ctx, store.Transaction{
			CustomerId: customerId,
			ProductId:  productId,
			Quantity:   1,
			TotalPrice: 20,
		}).Return(&store.Transaction{
		Id:         id,
		CreatedAt:  createdAt,
		CustomerId: customerId,
		ProductId:  productId,
		Quantity:   1,
		TotalPrice: 20,
	}, nil).Once()

	service := svc.NewTransactionService(suite.mock)
	result, err := service.CreateTransaction(suite.ctx, &request)
	suite.NoError(err)
	suite.NotNil(result)
	suite.Equal(result.Id, id.String())
	suite.Equal(result.IsSuccess, true)

}

func (suite *ServiceTestSuite) TestTransactionService_CreateTransaction_Error_InvalidArgument() {
	productId := uuid.New()

	request := pb.CreateTransactionRequest{
		CustomerId: "",
		ProductId:  productId.String(),
		Quantity:   1,
		TotalPrice: 20,
	}

	service := svc.NewTransactionService(suite.mock)
	result, err := service.CreateTransaction(suite.ctx, &request)
	suite.Error(err)
	suite.Equal(status.Error(codes.InvalidArgument, "invalid customer id"), err)
	suite.Nil(result)
}

func (suite *ServiceTestSuite) TestTransactionService_CreateTransaction_Error_Internal() {
	customerId := uuid.New()
	productId := uuid.New()

	request := pb.CreateTransactionRequest{
		CustomerId: customerId.String(),
		ProductId:  productId.String(),
		Quantity:   1,
		TotalPrice: 20,
	}

	suite.mock.
		On("InsertTransaction", suite.ctx, store.Transaction{
			CustomerId: customerId,
			ProductId:  productId,
			Quantity:   1,
			TotalPrice: 20,
		}).Return(nil, status.Error(codes.Internal, "Ops")).Once()

	service := svc.NewTransactionService(suite.mock)
	result, err := service.CreateTransaction(suite.ctx, &request)
	suite.Equal(status.Error(codes.Internal, "Ops"), err)
	suite.Nil(result)
}

//=================================

// =================================
// Test: GetTransactionById
// =================================

func (suite *ServiceTestSuite) TestTransactionService_GetTransactionById_SUCCESS() {
	transaction := store.Transaction{
		Id:         uuid.New(),
		CreatedAt:  time.Now(),
		ProductId:  uuid.New(),
		CustomerId: uuid.New(),
		Quantity:   1,
		TotalPrice: 20,
	}
	suite.mock.On("GetTransactionById", suite.ctx, transaction.Id).
		Return(&transaction, nil).
		Once()

	service := svc.NewTransactionService(suite.mock)
	result, err := service.GetTransactionById(suite.ctx, &pb.GetTransactionByIdRequest{
		Id: transaction.Id.String(),
	})

	suite.NoError(err)
	suite.NotNil(result)
	suite.Equal(transaction.Id.String(), result.Transaction.Id)
	suite.Equal(transaction.CreatedAt.UTC(), result.Transaction.CreatedAt.AsTime().UTC())
	suite.Equal(transaction.CustomerId.String(), result.Transaction.CustomerId)
	suite.Equal(transaction.ProductId.String(), result.Transaction.ProductId)
	suite.Equal(transaction.Quantity, result.Transaction.Quantity)
	suite.Equal(transaction.TotalPrice, result.Transaction.TotalPrice)
}

func (suite *ServiceTestSuite) TestTransactionService_GetTransactionById_Error_WrongUUIDFormat() {
	transaction := store.Transaction{
		Id:         uuid.New(),
		CreatedAt:  time.Now(),
		ProductId:  uuid.New(),
		CustomerId: uuid.New(),
		Quantity:   1,
		TotalPrice: 20,
	}
	suite.mock.On("GetTransactionById", suite.ctx, transaction.Id).
		Return(&transaction, nil).
		Once()

	service := svc.NewTransactionService(suite.mock)
	result, err := service.GetTransactionById(suite.ctx, &pb.GetTransactionByIdRequest{
		Id: "wrong uuid format",
	})

	suite.Equal(status.Error(codes.InvalidArgument, "invalid transaction id"), err)
	suite.Nil(result)

}

func (suite *ServiceTestSuite) TestTransactionService_GetTransactionById_Error_Internal() {
	transaction := store.Transaction{
		Id:         uuid.New(),
		CreatedAt:  time.Now(),
		ProductId:  uuid.New(),
		CustomerId: uuid.New(),
		Quantity:   1,
		TotalPrice: 20,
	}
	suite.mock.On("GetTransactionById", suite.ctx, transaction.Id).
		Return(nil, errors.New("")).
		Once()

	service := svc.NewTransactionService(suite.mock)
	result, err := service.GetTransactionById(suite.ctx, &pb.GetTransactionByIdRequest{
		Id: transaction.Id.String(),
	})

	suite.Equal(status.Error(codes.Internal, ""), err)
	suite.Nil(result)
}

func (suite *ServiceTestSuite) TestTransactionService_GetTransactionById_Error_NotFound() {
	transaction := store.Transaction{
		Id:         uuid.New(),
		CreatedAt:  time.Now(),
		ProductId:  uuid.New(),
		CustomerId: uuid.New(),
		Quantity:   1,
		TotalPrice: 20,
	}
	suite.mock.On("GetTransactionById", suite.ctx, transaction.Id).
		Return(nil, nil).
		Once()

	service := svc.NewTransactionService(suite.mock)
	result, err := service.GetTransactionById(suite.ctx, &pb.GetTransactionByIdRequest{
		Id: transaction.Id.String(),
	})

	suite.Equal(status.Error(codes.NotFound, "no record"), err)
	suite.Nil(result)

}

//=================================

// this func to run the test suite
func TestStoreTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
