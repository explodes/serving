package main

import (
	"context"
	"github.com/explodes/serving/addz"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:4003", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := addz.NewAddzServiceClient(conn)

	values := []int64{2, 3, 4}
	var cookie int64
	for {
		req := &addz.SubtractRequest{Cookie: cookie, Values: values}
		now := time.Now()
		res, err := client.Subtract(context.Background(), req)
		then := time.Now()
		if err != nil {
			log.Printf("Add ERROR: %v (cookie=%d)", err, cookie)
		} else {
			log.Printf("Add: %d (cookie=%d) (%s)", res.Result, cookie, then.Sub(now))
		}
		cookie++
		<-time.After(1 * time.Millisecond)
	}
}
