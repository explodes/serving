package main

import (
	"flag"
	"github.com/rs/zerolog/log"
	"os"
)

var (
	configFlag = flag.String("config", "sample.pbascii", "configuration file location")
)

func main() {
	flag.Parse()

	configFile,err := os.Open(*configFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

}

func readConfig() *logz.ServerConfig