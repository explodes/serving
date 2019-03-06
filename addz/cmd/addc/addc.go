package main

import (
	"context"
	"flag"
	"github.com/explodes/serving/addz"
	"github.com/explodes/serving/userz"
	"github.com/explodes/serving/utilz"
	"google.golang.org/grpc"
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

	addzConn, err := grpc.Dial(*flagAddzAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	utilz.RegisterGracefulShutdownCloser("addz-conn", addzConn)
	addzClient := addz.NewAddzServiceClient(addzConn)

	userzConn, err := grpc.Dial(*flagUserzAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	utilz.RegisterGracefulShutdownCloser("userz-conn", userzConn)
	userzClient := userz.NewClient(userz.NewUserzServiceClient(userzConn))

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
