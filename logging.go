package serving

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"time"
)

type logWriter struct{}

func (w logWriter) Write(bytes []byte) (int, error) {
	return fmt.Fprint(os.Stderr, time.Now().UTC().Format(time.RFC3339)+" "+string(bytes))
}

func init() {
	color.NoColor = false
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
}
