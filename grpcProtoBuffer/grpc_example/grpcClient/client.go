package main

import (
	"google.golang.org/grpc"
	"log"
	pb "Isabela/grpcProtoBuffer/grpc_example/datafiles"
	"context"
)

const (
	address = "localhost:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}

	defer conn.Close()
	c := pb.NewMoneyTransactionClient(conn)

	from := "5334"
	to := "6914"
	amount := float32(1250.75)

	r, err := c.MakeTransaction(context.Background(), &pb.TransactionRequest{From: from, To: to, Amount: amount})
	if err != nil {
		log.Fatalf("Could not transact %v", err)
	}

	log.Printf("Transaction confirmed %t", r.Confirmation)

}
