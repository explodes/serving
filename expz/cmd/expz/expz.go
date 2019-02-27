package main

import (
	"flag"
	"github.com/explodes/serving"
	"github.com/explodes/serving/expz"
	"github.com/explodes/serving/logz"
	"github.com/explodes/serving/statusz"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	configFlag = flag.String("config", "config.textproto", "configuration file location")
)

func main() {
	flag.Parse()

	config := &expz.ExpzConfig{}
	if err := serving.ReadConfigFile(*configFlag, config); err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	experiments, err := config.Validate()
	if err != nil {
		log.Fatal(err)
	}

	addr := config.BindAddress.Address()
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("expz listening on %s", addr)

	logzClient, err := logz.NewClient(config.LogzAddress)
	if err != nil {
		log.Fatalf("error connecting to logz: %v", err)
	}

	grpcServer := grpc.NewServer()
	expz.RegisterExpzServiceServer(grpcServer, expz.NewExpzServer(logzClient, experiments))
	statusz.RegisterStatuszServiceServer(grpcServer, statusz.NewStatuszServer())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("serving error: %v", err)
	}

}
