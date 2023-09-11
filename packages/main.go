package main

import (
	"ecommerece/packages/database"
	pb "ecommerece/packages/proto/transaction"
	"ecommerece/packages/transaction/store"
	"ecommerece/packages/transaction/v1"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
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

	listener, err := net.Listen("tcp", ":15935")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterTransactionServiceServer(s, transactionService)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
