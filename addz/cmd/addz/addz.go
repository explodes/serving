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
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
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

	expzClient, err := expz.NewClient(config.ExpzAddress.Address())
	if err != nil {
		log.Fatalf("error connecting to expz: %v", err)
	}

	userzClient, err := userz.NewClient(config.UserzAddress.Address())
	if err != nil {
		log.Fatalf("error connecting to expz: %v", err)
	}

	addzServer := addz.NewAddzServer(logzClient, expzClient, userzClient)
	statuszServer := statusz.NewStatuszServer()

	if config.JsonBindAddress != nil {
		go func() {
			log.Printf("Serving JSON at %s...\n", config.JsonBindAddress.Address())
			log.Printf("Serving status page at %s/statusz\n", config.JsonBindAddress.Address())
			if err := jsonpb.ServeJson(config.JsonBindAddress, addzServer, statuszServer); err != nil {
				log.Fatal(err)
			}
		}()
	}

	grpcServer := grpc.NewServer()
	addz.RegisterAddzServiceServer(grpcServer, addzServer)
	statusz.RegisterStatuszServiceServer(grpcServer, statuszServer)

	log.Printf("addz listening on %s...", addr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("serving error: %v", err)
	}

}

func combine(muxs ...*http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, mux := range muxs {
			if handler, pattern := mux.Handler(r); pattern != "" {
				handler.ServeHTTP(w, r)
				return
			}
		}
		http.NotFoundHandler().ServeHTTP(w, r)
	})
}
