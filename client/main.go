package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/Lucas32-dev/TodoGRPC/pb"
	"google.golang.org/grpc"
)

func main() {
	// Create address
	addr := flag.String("addr", "localhost:9000", "the address to connect to")

	// Set up connection to the server
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTodoClient(conn)

	// Comunicate with server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Send new item
	r, err := c.AddItem(ctx, &pb.Item{Title: "Walk with my little dog :)", Completed: true})

	if err != nil {
		log.Printf("could not add:%v", err)
	}

	log.Printf("server res: %v %v", r.GetSuccess(), r.GetMessage())

	// Remove the first item
	r, err = c.RemoveItem(ctx, &pb.DeleteRequest{Title: "Walk with my little dog :)"})

	if err != nil {
		log.Printf("could not remove:%v", err)
	}

	log.Printf("server res: %v %v", r.GetSuccess(), r.GetMessage())

}
