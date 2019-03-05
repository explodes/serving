package main

import (
	"flag"
	"github.com/explodes/serving"
	"github.com/explodes/serving/jsonpb"
	"github.com/explodes/serving/logz"
	"github.com/explodes/serving/statusz"
	"github.com/explodes/serving/utilz"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	configFlag = flag.String("config", "config.textproto", "configuration file location")
	colorFlag  = flag.Bool("color", true, "whether or not to print in colors")
)

func main() {
	flag.Parse()

	// Enable/disable color output
	color.NoColor = !(*colorFlag)

	config := &logz.ServiceConfig{}
	if err := serving.ReadConfigFile(*configFlag, config); err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	addr := config.BindAddress.Address()
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if config.StatuszAddress != nil {
		go func() {
			log.Printf("Serving status page at %s/statusz\n", config.StatuszAddress.Address())
			if err := jsonpb.ServeStatusz(config.StatuszAddress); err != nil {
				log.Printf("statusz server error: %v", err)
			}
		}()
	}

	grpcServer := grpc.NewServer()
	logz.RegisterLogzServiceServer(grpcServer, logz.NewLogzServer(config))
	statusz.RegisterStatuszServiceServer(grpcServer, statusz.NewStatuszServer())
	utilz.RegisterGracefulShutdownGrpcServer("grpc-server", grpcServer)

	go func() {
		log.Printf("logz listening on %s...", addr)
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("grpc server error: %v", err)
		}
	}()

	<-utilz.GracefulShutdown()
}
