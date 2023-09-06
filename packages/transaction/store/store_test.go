package store_test

//todo: create setup method
//todo:
import (
	"context"
	"ecommerece/packages/database"
	"ecommerece/packages/transaction/store"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

//func TestTransactionStore_InsertTransaction_SUCCESS(t *testing.T) {
//	//set up database connection
//	//todo: move it to be used in all tests
//	ctx := context.Background()
//	conn, err := database.NewCockroachDB("postgresql://root@localhost:26257/ecommerece")
//	if err != nil {
//		fmt.Println("Error connecting to CockroachDB:", err)
//		return
//	}
//
//	tx, err := conn.Begin(ctx)
//	defer func() {
//		err := tx.Rollback(ctx)
//		if err != nil {
//			return
//		}
//	}()
//	transactionStore := store.NewTransactionStore(tx.Conn())
//
//	//create the input transaction:
//	inputTransaction := store.Transaction{
//		CustomerId: uuid.New(),
//		ProductId:  uuid.New(),
//		Quantity:   3,
//		TotalPrice: 30,
//	}
//
//	resultTransaction, err := transactionStore.InsertTransaction(ctx, inputTransaction)
//	log.Println(resultTransaction.Id)
//	databaseTransaction := transactionStore.GetTransactionById(ctx, resultTransaction.Id)
//	log.Println(resultTransaction.Id)
//
//	assert.Nil(t, err)
//	assert.NotNil(t, resultTransaction)
//	assert.NotNil(t, resultTransaction.Id)
//	assert.NotNil(t, resultTransaction.CreatedAt)
//	assert.NotNil(t, resultTransaction.ProductId)
//	assert.NotNil(t, resultTransaction.CustomerId)
//	assert.NotNil(t, resultTransaction.Quantity)
//	assert.NotNil(t, resultTransaction.TotalPrice)
//
//	//checking the request with the response from the store
//	assert.Equal(t, inputTransaction.CustomerId, resultTransaction.CustomerId)
//	assert.Equal(t, inputTransaction.ProductId, resultTransaction.ProductId)
//	assert.Equal(t, inputTransaction.Quantity, resultTransaction.Quantity)
//	assert.Equal(t, inputTransaction.TotalPrice, resultTransaction.TotalPrice)
//
//	assert.NotNil(t, databaseTransaction)
//	//checking the response from the store with the response from the database
//	assert.Equal(t, databaseTransaction.Id, resultTransaction.Id)
//	assert.Equal(t, databaseTransaction.CreatedAt.UTC(), resultTransaction.CreatedAt.UTC())
//	assert.Equal(t, databaseTransaction.CustomerId, resultTransaction.CustomerId)
//	assert.Equal(t, databaseTransaction.ProductId, resultTransaction.ProductId)
//	assert.Equal(t, databaseTransaction.Quantity, resultTransaction.Quantity)
//	assert.Equal(t, databaseTransaction.TotalPrice, resultTransaction.TotalPrice)
//
//	//assert.Equal(t, databaseTransaction.TotalPrice, resultTransaction.TotalPrice)
//	//assert.NoError(t, err, "There are no errors")
//
//}

type StoreTestSuite struct {
	suite.Suite
	db    database.Database
	store store.Store
	tx    pgx.Tx
}

func (suite *StoreTestSuite) SetupSuite() {
	// Set up a test database connection pool
	//pool, err := pgxpool.New(context.Background(), "postgresql://root@localhost:26257/ecommerece")
	db, err := database.NewCockroachDB("postgresql://root@localhost:26257/ecommerece")
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.db = db

	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.tx = tx

	suite.store = store.NewTransactionStore(suite.db)
}

func (suite *StoreTestSuite) TearDownSuite() {
	err := suite.tx.Rollback(context.Background())
	if err != nil {
		suite.T().Fatal(err)
	}
}

func (suite *StoreTestSuite) TestTransactionStore_InsertTransaction_SUCCESS() {
	ctx := context.Background()

	inputTransaction := store.Transaction{
		CustomerId: uuid.New(),
		ProductId:  uuid.New(),
		Quantity:   3,
		TotalPrice: 30,
	}

	resultTransaction, err := suite.store.InsertTransaction(ctx, inputTransaction)
	log.Println(resultTransaction.Id)

	databaseTransaction := suite.store.GetTransactionById(ctx, resultTransaction.Id)
	log.Println(resultTransaction.Id)

	suite.Nil(err)
	suite.NotNil(resultTransaction)
	suite.NotNil(resultTransaction.Id)
	suite.NotNil(resultTransaction.CreatedAt)
	suite.NotNil(resultTransaction.ProductId)
	suite.NotNil(resultTransaction.CustomerId)
	suite.NotNil(resultTransaction.Quantity)
	suite.NotNil(resultTransaction.TotalPrice)

	//checking the request with the response from the store
	suite.Equal(inputTransaction.CustomerId, resultTransaction.CustomerId)
	suite.Equal(inputTransaction.ProductId, resultTransaction.ProductId)
	suite.Equal(inputTransaction.Quantity, resultTransaction.Quantity)
	suite.Equal(inputTransaction.TotalPrice, resultTransaction.TotalPrice)

	suite.NotNil(databaseTransaction)
	//checking the response from the store with the response from the database
	suite.Equal(databaseTransaction.Id, resultTransaction.Id)
	suite.Equal(databaseTransaction.CreatedAt.UTC(), resultTransaction.CreatedAt.UTC())
	suite.Equal(databaseTransaction.CustomerId, resultTransaction.CustomerId)
	suite.Equal(databaseTransaction.ProductId, resultTransaction.ProductId)
	suite.Equal(databaseTransaction.Quantity, resultTransaction.Quantity)
	suite.Equal(databaseTransaction.TotalPrice, resultTransaction.TotalPrice)

}

func TestStoreTestSuite(t *testing.T) {
	suite.Run(t, new(StoreTestSuite))
}
