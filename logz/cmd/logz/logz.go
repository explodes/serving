package main

import (
	"flag"
	"github.com/explodes/serving"
	"github.com/explodes/serving/logz"
	"github.com/explodes/serving/statusz"
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
	log.Printf("logz listening on %s", addr)

	grpcServer := grpc.NewServer()
	logz.RegisterLogzServiceServer(grpcServer, logz.NewLogzServer(config))
	statusz.RegisterStatuszServiceServer(grpcServer, statusz.NewStatuszServer())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("serving error: %v", err)
	}
}
