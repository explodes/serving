package main

import (
	"flag"
	"github.com/explodes/serving"
	"github.com/explodes/serving/addz"
	"github.com/explodes/serving/expz"
	"github.com/explodes/serving/jsonpb"
	"github.com/explodes/serving/logz"
	"github.com/explodes/serving/statusz"
	"github.com/explodes/serving/userz"
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

	config := &addz.AddzConfig{}
	if err := serving.ReadConfigFile(*configFlag, config); err != nil {
		log.Fatalf("error reading config file: %v", err)
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

	expzClient, err := expz.NewClient(config.ExpzAddress.Address())
	if err != nil {
		log.Fatalf("error connecting to expz: %v", err)
	}
	utilz.RegisterGracefulShutdownCloser("expz-client", expzClient)

	userzClient, err := userz.NewClient(config.UserzAddress.Address())
	if err != nil {
		log.Fatalf("error connecting to userz: %v", err)
	}
	utilz.RegisterGracefulShutdownCloser("userz-client", userzClient)

	addzServer := addz.NewAddzServer(logzClient, expzClient, userzClient)
	statuszServer := statusz.NewStatuszServer()

	if config.JsonBindAddress != nil {
		go func() {
			log.Printf("Serving JSON at http://%s...\n", config.JsonBindAddress.Address())
			log.Printf("Serving status page at http://%s/statusz\n", config.JsonBindAddress.Address())
			if err := jsonpb.ServeJson(config.JsonBindAddress, addzServer, statuszServer); err != nil {
				log.Fatal(err)
			}
		}()
	}

	grpcServer := grpc.NewServer()
	addz.RegisterAddzServiceServer(grpcServer, addzServer)
	statusz.RegisterStatuszServiceServer(grpcServer, statuszServer)
	utilz.RegisterGracefulShutdownGrpcServer("grpc-server", grpcServer)

	go func() {
		log.Printf("userz listening on %s...", addr)
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("grpc server error: %v", err)
		}
	}()

	<-utilz.GracefulShutdown()
}
