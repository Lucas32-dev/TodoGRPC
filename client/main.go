package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/Lucas32-dev/TodoGRPC/pb"
	"google.golang.org/grpc"
)

func AddMultiple(ctx context.Context, c pb.TodoClient, repeats int) {
	for i := 0; i < repeats; i++ {
		r, err := c.AddItem(ctx, &pb.Item{Title: strconv.Itoa(rand.Int()), Completed: true})

		if err != nil {
			log.Printf("could not add:%v", err)
		}

		log.Printf("add req -> success: %v, res: %v", r.GetSuccess(), r.GetMessage())
	}
}

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

	AddMultiple(ctx, c, 4)

	r, err := c.GetItems(ctx, &pb.GetItemsRequest{})
	if err != nil {
		log.Panicf("fail to get all:%v", err)
	}
	log.Printf("All:%v", r.GetItems())
}
