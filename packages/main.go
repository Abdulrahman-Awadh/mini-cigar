package main

import (
	"context"
	analyticssvc "github.com/Abdulrahman-Awadh/mini-cigar/packages/analytics/v1"
	"github.com/Abdulrahman-Awadh/mini-cigar/packages/database"
	analyticspb "github.com/Abdulrahman-Awadh/mini-cigar/packages/proto/analytics"
	pb "github.com/Abdulrahman-Awadh/mini-cigar/packages/proto/transaction"
	"google.golang.org/grpc"
	"net"

	"fmt"
	"github.com/Abdulrahman-Awadh/mini-cigar/packages/transaction/store"
	"github.com/Abdulrahman-Awadh/mini-cigar/packages/transaction/v1"
	"log"
)

func main() {
	log.Print("main start")
	//Database connection setup
	conn, err := database.NewCockroachDB("postgresql://root@localhost:26257/ecommerece")
	if err != nil {
		fmt.Println("Error connecting to CockroachDB:", err)
		return
	}
	defer conn.Close(context.Background())

	grpcServer := grpc.NewServer()

	//initialize transaction service:
	transactionStore := store.NewTransactionStore(conn)
	transactionService := v1.NewTransactionService(transactionStore)
	pb.RegisterTransactionServiceServer(grpcServer, transactionService)

	//initialize analytics service:
	analyticsService := analyticssvc.NewAnalyticsService(transactionService)
	analyticspb.RegisterAnalyticsServiceServer(grpcServer, analyticsService)

	lis, err := net.Listen("tcp", ":15935")
	if err != nil {
		log.Fatalln("Fail to listen")
	}

	//to check the services registered in the grpc
	log.Println(grpcServer.GetServiceInfo())
	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatalln("Fail to serve")
	}

}
