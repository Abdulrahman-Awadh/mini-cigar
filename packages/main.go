package main

import (
	"ecommerece/packages/database"
	"ecommerece/packages/transaction/store"
	"ecommerece/packages/transaction/v1"
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
	defer conn.Close(context.Background())

	transactionStore := store.NewTransactionStore(tx.Conn())
	transactionService := v1.NewTransactionService(transactionStore)
	print(transactionService)

}
