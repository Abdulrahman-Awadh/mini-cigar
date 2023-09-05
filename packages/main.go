package main

import (
	"ecommerece/packages/database"
	"ecommerece/packages/transaction"
	"ecommerece/packages/transaction/store"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"log"
)

func main() {
	fmt.Print("test")
	//conn, err := database.ConnectToCockroachDB("postgresql://root@localhost:26257/ecommerece")
	conn, err := database.NewCockroachDB("postgresql://root@localhost:26257/ecommerece")
	if err != nil {
		fmt.Println("Error connecting to CockroachDB:", err)
		return
	}
	tx, err := conn.Begin(context.Background())
	defer conn.Pool.Close()

	transactionStore := store.NewTransactionStore(tx)
	transactionService := transaction.NewTransactionService(transactionStore)
	//print(transactionService)
	//txn := transactionService.GetTransactionById(context.Background(), uuid.MustParse("04e6919b-1cde-429f-a238-9199731f67db"))
	//print(txn)

	//listener, err := net.Listen("tcp", ":8080")
	//if err != nil {
	//	panic(err)
	//}
	//
	//s := grpc.NewServer()
	//pb.RegisterTransactionServiceServer(s, &server{})
	//if err := s.Serve(listener); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}

	//for i := 0; i < 20; i++ {
	//	InsertTransaction(context.Background(), conn, Transaction{
	//		Id:         uuid.New(),
	//		CreatedAt:  time.Now(),
	//		ProductId:  uuid.New(),
	//		CustomerId: uuid.New(),
	//		Quantity:   rand.Int(),
	//		TotalPrice: rand.Float32(),
	//	})
	//}

	transactionID := uuid.MustParse("d9d6d5f1-dc61-40c5-ba38-4887d09f122c")
	//transactionID, err := uuid.Parse("d9d6d5f1-dc61-40c5-ba38-4887d09f122c")
	//if err != nil {
	//	log.Fatalf("Error In Parsing the ID: %v", err)
	//}
	//
	//
	//
	//
	transaction := transactionService.GetTransactionById(context.Background(), transactionID)
	if transaction == nil {
		fmt.Println("Transaction not found or an error occurred.")
		return
	}

	fmt.Printf("Transaction ID: %s, CreatedAt: %s, Customer ID: %s, Product ID: %s, Quantity: %d, Total Price: %.2f\n",
		transaction.Id, transaction.CreatedAt, transaction.CustomerId, transaction.ProductId, transaction.Quantity, transaction.TotalPrice)

	transactions, err := transactionService.GetAllTransactions(context.Background())
	if err != nil {
		log.Fatalf("Error getting transactions: %v", err)
	}

	for _, transaction := range transactions {
		fmt.Printf("Transaction ID: %s, CreatedAt: %s, Customer ID: %s, Product ID: %s, Quantity: %d, Total Price: %.2f\n",
			transaction.Id, transaction.CreatedAt, transaction.CustomerId, transaction.ProductId, transaction.Quantity, transaction.TotalPrice)
	}
}
