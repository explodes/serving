package main

import (
	"context"
	"flag"
	"github.com/explodes/serving/addz"
	"github.com/explodes/serving/userz"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

const (
	defaultLoopFrequency = 10 * time.Second
)

var (
	flagAddzAddr  = flag.String("addz", "0.0.0.0:4003", "addz server address")
	flagUserzAddr = flag.String("userz", "0.0.0.0:4004", "userz server address")
	flagLoopFrq   = flag.Duration("frq", defaultLoopFrequency, "loop frequency")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*flagAddzAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer mustClose(conn)

	addzClient := addz.NewAddzServiceClient(conn)
	userzClient, err := userz.NewClient(*flagUserzAddr)
	if err != nil {
		log.Fatal(err)
	}

	values := []int64{2, 3, 4}

	cookieInt := 0

	for {
		cookieStr, err := userzClient.Login(context.Background(), "test", "test")
		if err != nil {
			log.Fatal(err)
			<-time.After(*flagLoopFrq)
			continue
		}
		log.Printf("got cookie %s", cookieStr)

		req := &addz.SubtractRequest{Cookie: cookieStr, Values: values}
		now := time.Now()
		res, err := addzClient.Subtract(context.Background(), req)
		then := time.Now()
		if err != nil {
			log.Printf("Subtract ERROR: %v", err)
		} else {
			log.Printf("Subtract: %d (%s)", res.Result, then.Sub(now))
		}
		cookieInt++
		<-time.After(*flagLoopFrq)
	}
}

func mustClose(c io.Closer) {
	if err := c.Close(); err != nil {
		panic(err)
	}
}
