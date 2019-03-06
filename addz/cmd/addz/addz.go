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

	logzConn, err := utilz.DialGrpc(config.LogzServer)
	if err != nil {
		log.Fatalf("error dialing logz: %v", err)
	}
	utilz.RegisterGracefulShutdownCloser("logz-conn", logzConn)
	logzClient := logz.NewClient(logz.NewLogzServiceClient(logzConn))

	expzConn, err := utilz.DialGrpc(config.ExpzServer)
	if err != nil {
		log.Fatalf("error dialing expz: %v", err)
	}
	utilz.RegisterGracefulShutdownCloser("expz-conn", expzConn)
	expzClient := expz.NewClient(expz.NewExpzServiceClient(expzConn))

	userzConn, err := utilz.DialGrpc(config.UserzServer)
	if err != nil {
		log.Fatalf("error dialing userz: %v", err)
	}
	utilz.RegisterGracefulShutdownCloser("userz-conn", userzConn)
	userzClient := userz.NewClient(userz.NewUserzServiceClient(userzConn))

	addzServer := addz.NewAddzServer(logzClient, expzClient, userzClient)
	statuszServer := statusz.NewStatuszServer()

	if config.JsonBindAddress != nil {
		go func() {
			log.Printf("Serving JSON at http://%s...\n", config.JsonBindAddress.Address())
			log.Printf("Serving status page at http://%s/statusz\n", config.JsonBindAddress.Address())
			if err := jsonpb.ServeJson(config.JsonBindAddress, addzServer, statuszServer); err != nil {
				log.Printf("error serving json: %v", err)
			}
		}()
	}

	grpcServer := grpc.NewServer()
	addz.RegisterAddzServiceServer(grpcServer, addzServer)
	statusz.RegisterStatuszServiceServer(grpcServer, statuszServer)
	utilz.RegisterGracefulShutdownGrpcServer("grpc-server", grpcServer)

	go func() {
		log.Printf("addz listening on %s...", addr)
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("grpc server error: %v", err)
		}
	}()

	<-utilz.GracefulShutdown()
}
