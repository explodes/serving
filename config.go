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
	defer mustClose(f)

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	if err := proto.UnmarshalText(string(b), pb); err != nil {
		return err
	}

	return nil
}

func mustClose(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}
