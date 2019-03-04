package main

import (
	"context"
	"fmt"
	"github.com/explodes/serving/addz"
	"github.com/explodes/serving/userz"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:4003", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer mustClose(conn)

	client := addz.NewAddzServiceClient(conn)

	values := []int64{2, 3, 4}

	cookieInt := 0

	for {
		cookie := &userz.Cookie{
			SessionId: fmt.Sprint(cookieInt),
		}
		cookieStr, err := userz.SerializeCookie(cookie)
		if err != nil {
			log.Fatalf("cookie failure: %v", err)
		}

		req := &addz.SubtractRequest{Cookie: cookieStr, Values: values}
		now := time.Now()
		res, err := client.Subtract(context.Background(), req)
		then := time.Now()
		if err != nil {
			log.Printf("Subtract ERROR: %v (cookie=%s)", err, cookieStr)
		} else {
			log.Printf("Subtract: %d (cookie=%s) (%s)", res.Result, cookieStr, then.Sub(now))
		}
		cookieInt++
		<-time.After(800 * time.Millisecond)
	}
}

func mustClose(c io.Closer) {
	if err := c.Close(); err != nil {
		panic(err)
	}
}
