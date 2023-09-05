package main

import (
	"ecommerece/packages/database"
	"ecommerece/packages/transaction"
	"ecommerece/packages/transaction/store"
	"fmt"
	"golang.org/x/net/context"
	"log"
)

func main() {
	log.Print("main start")

	conn, err := database.NewCockroachDB("postgresql://root@localhost:26257/ecommerece")
	if err != nil {
		fmt.Println("Error connecting to CockroachDB:", err)
		return
	}
	tx, err := conn.Begin(context.Background())
	defer conn.Pool.Close()

	transactionStore := store.NewTransactionStore(tx)
	transactionService := transaction.NewTransactionService(transactionStore)
	print(transactionService)

}
