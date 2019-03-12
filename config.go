package serving

import (
	"github.com/golang/protobuf/proto"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func ReadConfigFile(configPath string, pb proto.Message) error {
	f, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("error closing config file: %v", err)
		}
	}()
	return ReadConfig(f, pb)
}

func ReadConfig(r io.Reader, pb proto.Message) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	if err := proto.UnmarshalText(string(b), pb); err != nil {
		return err
	}
	return nil
}
