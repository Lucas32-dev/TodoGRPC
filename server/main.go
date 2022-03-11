package main

import (
	"context"
	"log"
	"net"

	"github.com/Lucas32-dev/TodoGRPC/pb"
	"google.golang.org/grpc"
)

// Server is used to implement the todo service
type server struct {
	pb.UnimplementedTodoServer
}

// Make an map Title -> Item with with inicial capacity of 50 entrys
var todoList = make(map[string]*pb.Item, 50)

// AddItem implement todo service interface
func (s *server) AddItem(ctx context.Context, in *pb.Item) (*pb.CommonActionReply, error) {
	log.Printf("Item received: %v", in)
	if todoList[in.GetTitle()] != nil {
		return &pb.CommonActionReply{Success: false, Message: "iten already exist"}, nil
	}
	// Add new item
	todoList[in.Title] = in
	return &pb.CommonActionReply{Success: true, Message: "saved!"}, nil
}

// RemoveItem implement todo service interface
func (s *server) RemoveItem(ctx context.Context, in *pb.DeleteRequest) (*pb.CommonActionReply, error) {
	// Remove from map
	delete(todoList, in.Title)
	log.Printf("Item %s removed", in.GetTitle())
	return &pb.CommonActionReply{Success: true, Message: "deleted"}, nil
}

// GetAll implement todo service interface
func (s *server) GetItems(ctx context.Context, in *pb.GetItemsRequest) (*pb.GetItemsReply, error) {

	items := make([]*pb.Item, 0, len(todoList))

	for _, v := range todoList {
		items = append(items, v)
	}

	return &pb.GetItemsReply{Items: items}, nil
}

func main() {
	// listen to port 9000
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen to port :9000, err: %v", err)
	}
	// Create a new grpc server
	grpcServer := grpc.NewServer()

	// Register service
	pb.RegisterTodoServer(grpcServer, &server{})

	// Start the server with the listener
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server: %s", err)
	}

}
