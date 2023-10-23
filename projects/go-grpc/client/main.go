package main

import (
	"context"
	"log"

	pb "bookshop/client/pb/inventory"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}

	defer conn.Close()

	client := pb.NewInventoryClient(conn)
	bookList, err := client.GetBookList(context.Background(), &pb.GetBookListRequest{})
	if err != nil {
		log.Fatalf("Failed to get book list: %v", err)
	}

	log.Printf("Book list: %v", bookList)

}
