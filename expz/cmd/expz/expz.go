package main

import (
	"flag"
	"github.com/explodes/serving"
	"github.com/explodes/serving/expz"
	"github.com/explodes/serving/jsonpb"
	"github.com/explodes/serving/logz"
	"github.com/explodes/serving/statusz"
	"github.com/explodes/serving/utilz"
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

	logzClient, err := logz.NewClient(config.LogzAddress.Address())
	if err != nil {
		log.Fatalf("error connecting to logz: %v", err)
	}
	utilz.RegisterGracefulShutdownCloser("logz-client", logzClient)

	if config.StatuszAddress != nil {
		go func() {
			log.Printf("Serving status page at %s/statusz\n", config.StatuszAddress.Address())
			if err := jsonpb.ServeStatusz(config.StatuszAddress); err != nil {
				log.Printf("statusz server error: %v", err)
			}
		}()
	}

	grpcServer := grpc.NewServer()
	expz.RegisterExpzServiceServer(grpcServer, expz.NewExpzServer(logzClient, experiments))
	statusz.RegisterStatuszServiceServer(grpcServer, statusz.NewStatuszServer())
	utilz.RegisterGracefulShutdownGrpcServer("grpc-server", grpcServer)

	go func() {
		log.Printf("expz listening on %s...", addr)
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("grpc server error: %v", err)
		}
	}()

	<-utilz.GracefulShutdown()

}
