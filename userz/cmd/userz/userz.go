package main

import (
	"flag"
	"github.com/explodes/serving"
	"github.com/explodes/serving/expz"
	"github.com/explodes/serving/jsonpb"
	"github.com/explodes/serving/logz"
	"github.com/explodes/serving/statusz"
	"github.com/explodes/serving/userz"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	configFlag = flag.String("config", "config.textproto", "configuration file location")
)

func main() {
	flag.Parse()

	config := &userz.UserzConfig{}
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

	expzClient, err := expz.NewClient(config.ExpzAddress.Address())
	if err != nil {
		log.Fatalf("error connecting to expz: %v", err)
	}

	userzServer := userz.NewUserzServer(config.CookiePasscode, logzClient, expzClient)
	statuzServer := statusz.NewStatuszServer()

	if config.JsonBindAddress != nil {
		go func() {
			log.Printf("Serving JSON at %s...\n", config.JsonBindAddress.Address())
			log.Printf("Serving status page at %s/statusz\n", config.JsonBindAddress.Address())
			if err := jsonpb.ServeJson(config.JsonBindAddress, userzServer, statuzServer); err != nil {
				log.Fatal(err)
			}
		}()
	}

	grpcServer := grpc.NewServer()
	userz.RegisterUserzServiceServer(grpcServer, userzServer)
	statusz.RegisterStatuszServiceServer(grpcServer, statuzServer)

	log.Printf("userz listening on %s...", addr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("serving error: %v", err)
	}

}
