package main

import (
	pb "bookshop/server/pb/inventory"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedInventoryServer
}

func (s *server) GetBookList(c context.Context, in *pb.GetBookListRequest) (*pb.GetBookListResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	return &pb.GetBookListResponse{
		Books: getSampleBooks(),
	}, nil
}

func getSampleBooks() []*pb.Book {
	return []*pb.Book{
		{
			Title:     "The Alchemist",
			Author:    "Paulo Coelho",
			PageCount: 197,
			Language:  "English",
		},
		{
			Title:     "Kafka on the Shore",
			Author:    "Haruki Murakami",
			PageCount: 505,
			Language:  "English",
		},
	}

}

func main() {
	listner, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterInventoryServer(s, &server{})
	log.Printf("Starting server on port :8000")
	if err := s.Serve(listner); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
