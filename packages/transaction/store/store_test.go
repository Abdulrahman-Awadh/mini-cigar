package store_test

//todo: create setup method
//todo:
import (
	"context"
	"ecommerece/packages/database"
	"ecommerece/packages/transaction/store"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestTransactionStore_InsertTransaction_SUCCESS(t *testing.T) {
	//set up database connection
	//todo: move it to be used in all tests
	ctx := context.Background()
	conn, err := database.NewCockroachDB("postgresql://root@localhost:26257/ecommerece")
	if err != nil {
		fmt.Println("Error connecting to CockroachDB:", err)
		return
	}

	tx, err := conn.Begin(ctx)
	defer func() {
		log.Println("close func")
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
		conn.Close()
	}()
	transactionStore := store.NewTransactionStore(tx)

	//create the input transaction:
	inputTransaction := store.Transaction{
		CustomerId: uuid.New(),
		ProductId:  uuid.New(),
		Quantity:   3,
		TotalPrice: 30,
	}

	resultTransaction, err := transactionStore.InsertTransaction(ctx, inputTransaction)
	log.Println(resultTransaction.Id)
	databaseTransaction := transactionStore.GetTransactionById(ctx, resultTransaction.Id)
	log.Println(resultTransaction.Id)
	t.Log("ghfgjgv")
	assert.Nil(t, err)
	assert.NotNil(t, resultTransaction)
	assert.NotNil(t, resultTransaction.Id)
	assert.NotNil(t, resultTransaction.CreatedAt)
	assert.NotNil(t, resultTransaction.ProductId)
	assert.NotNil(t, resultTransaction.CustomerId)
	assert.NotNil(t, resultTransaction.Quantity)
	assert.NotNil(t, resultTransaction.TotalPrice)

	//checking the request with the response from the store
	assert.Equal(t, inputTransaction.CustomerId, resultTransaction.CustomerId)
	assert.Equal(t, inputTransaction.ProductId, resultTransaction.ProductId)
	assert.Equal(t, inputTransaction.Quantity, resultTransaction.Quantity)
	assert.Equal(t, inputTransaction.TotalPrice, resultTransaction.TotalPrice)

	assert.NotNil(t, databaseTransaction)
	//checking the response from the store with the response from the database
	assert.Equal(t, databaseTransaction.Id, resultTransaction.Id)
	assert.Equal(t, databaseTransaction.CreatedAt.UTC(), resultTransaction.CreatedAt.UTC())
	assert.Equal(t, databaseTransaction.CustomerId, resultTransaction.CustomerId)
	assert.Equal(t, databaseTransaction.ProductId, resultTransaction.ProductId)
	assert.Equal(t, databaseTransaction.Quantity, resultTransaction.Quantity)
	assert.Equal(t, databaseTransaction.TotalPrice, resultTransaction.TotalPrice)

	//assert.Equal(t, databaseTransaction.TotalPrice, resultTransaction.TotalPrice)
	//assert.NoError(t, err, "There are no errors")

}
