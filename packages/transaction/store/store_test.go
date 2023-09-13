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

// struct to be used inside the suite
type StoreTestSuite struct {
	suite.Suite
	db    database.Database
	store store.Store
	tx    pgx.Tx
}

func setupTestDatabase() {
	db, err := pgx.Connect(context.Background(), "postgresql://root@localhost:26257")
	if err != nil {
		log.Fatal(err)
	}
	//todo: create database for test
	_, err = db.Exec(context.Background(), "CREATE DATABASE IF NOT EXISTS ecommerce_test")
	if err != nil {
		log.Fatal(err)
	}

	//todo: create transaction table
	_, err = db.Exec(context.Background(), "CREATE DATABASE IF NOT EXISTS ecommerce_test")
	if err != nil {
		log.Fatal(err)
	}

	q := `CREATE TABLE IF NOT EXISTS ecommerce_test.transaction(
			id UUID NOT NULL DEFAULT gen_random_uuid(),
			created_at TIMESTAMP NOT NULL,
			customer_id UUID NOT NULL,
			product_id UUID NOT NULL,
			quantity INT8 NOT NULL,
			total_price FLOAT8 NOT NULL,
			CONSTRAINT transaction_pkey PRIMARY KEY (id ASC)
		)`
	//_, err = db.Exec(context.Background(), 'CREATE TABLE IF NOT EXISTS ecommerce_test.transaction')
	_, err = db.Exec(context.Background(), q)
	if err != nil {
		log.Fatal(err)
	}

}

func (suite *StoreTestSuite) SetupSuite() {
	//todo create testdatabse
	setupTestDatabase()
	db, err := database.NewCockroachDB("postgresql://root@localhost:26257/ecommerce_test")
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

// =================================
// Testing: InsertTransaction
// =================================
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

	databaseTransaction, err := suite.store.GetTransactionById(ctx, resultTransaction.Id)
	suite.NoError(err)
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

//=================================

// =================================
// Testing: GetTransactionById
// =================================
func (suite *StoreTestSuite) TestTransactionStore_GetTransactionById_SUCCESS() {
	//create new transaction struct
	input := store.Transaction{
		ProductId:  uuid.New(),
		CustomerId: uuid.New(),
		TotalPrice: 30,
		Quantity:   2,
	}

	//insert it in database
	funcOutput, err := suite.store.InsertTransaction(context.Background(), input)
	suite.NoError(err)

	//get the Id in order to use it in GetTransactionById
	databaseOutput, err := suite.store.GetTransactionById(context.Background(), funcOutput.Id)
	suite.NoError(err)

	suite.NotNil(funcOutput)
	suite.NotNil(databaseOutput)
	suite.NoError(err)

	suite.Equal(funcOutput.Id, databaseOutput.Id)
	suite.Equal(funcOutput.CreatedAt, databaseOutput.CreatedAt)
	suite.Equal(funcOutput.CustomerId, databaseOutput.CustomerId)
	suite.Equal(funcOutput.ProductId, databaseOutput.ProductId)
	suite.Equal(funcOutput.Quantity, databaseOutput.Quantity)
	suite.Equal(funcOutput.TotalPrice, databaseOutput.TotalPrice)
}

func (suite *StoreTestSuite) TestTransactionStore_GetTransactionById_Error_NoRecord() {
	databaseOutput, err := suite.store.GetTransactionById(context.Background(), uuid.New())
	suite.NoError(err)
	suite.Nil(databaseOutput)
}

//=================================

// =================================
// Testing: GetAllTransaction
// =================================
func (suite *StoreTestSuite) TestTransactionStore_GetAllTransaction_SUCCESS() {
	var inputs []*store.Transaction
	for i := 0; i < 5; i++ {
		input := &store.Transaction{
			ProductId:  uuid.New(),
			CustomerId: uuid.New(),
			TotalPrice: 30,
			Quantity:   2,
		}
		inputs = append(inputs, input)
	}

	var funcOutputs []*store.Transaction
	for _, transaction := range inputs {
		funcOutput, err := suite.store.InsertTransaction(context.Background(), *transaction)
		suite.NoError(err)
		funcOutputs = append(funcOutputs, funcOutput)
	}

	databaseOutputs, err := suite.store.GetAllTransaction(context.Background())
	suite.NoError(err)
	suite.NotNil(funcOutputs)
	suite.NotNil(databaseOutputs)

	suite.Equal(len(funcOutputs), len(databaseOutputs))
	for i, transaction := range databaseOutputs {
		suite.NotNil(transaction)
		suite.NotNil(funcOutputs[i])

		suite.Equal(funcOutputs[i].Id, transaction.Id)
		suite.Equal(funcOutputs[i].CreatedAt, transaction.CreatedAt)
		suite.Equal(funcOutputs[i].CustomerId, transaction.CustomerId)
		suite.Equal(funcOutputs[i].ProductId, transaction.ProductId)
		suite.Equal(funcOutputs[i].Quantity, transaction.Quantity)
		suite.Equal(funcOutputs[i].TotalPrice, transaction.TotalPrice)
	}

}

func (suite *StoreTestSuite) TestTransactionStore_GetAllTransaction_Error_NoRecord() {
	result, err := suite.store.GetAllTransaction(context.Background())

	suite.NoError(err)
	suite.Nil(result)
}

// =================================

// =================================
//todo: GetTotalSales test
// =================================

func (suite *StoreTestSuite) TestTransactionStore_GetTotalSales_SUCCESS() {
	//todo insert record to db
	inputTransaction := store.Transaction{
		CustomerId: uuid.New(),
		ProductId:  uuid.New(),
		Quantity:   3,
		TotalPrice: 30,
	}
	_, err := suite.store.InsertTransaction(context.Background(), inputTransaction)
	suite.NoError(err)

	_, err = suite.store.InsertTransaction(context.Background(), inputTransaction)
	suite.NoError(err)

	_, err = suite.store.InsertTransaction(context.Background(), inputTransaction)
	suite.NoError(err)

	result, err := suite.store.GetTotalSales(context.Background())
	suite.NoError(err)
	suite.NotNil(result)
	suite.Equal(float32(90), *result)
}

// =================================

// =================================
//todo: GetSalesByProductId test
// =================================

// =================================

// =================================
//todo: GetTopFiveCustomersId test
// =================================

// =================================

// this func to run the test suite
func TestStoreTestSuite(t *testing.T) {
	suite.Run(t, new(StoreTestSuite))
}
