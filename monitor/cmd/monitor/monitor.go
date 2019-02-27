package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"github.com/explodes/serving"
	"github.com/explodes/serving/monitor"
	"github.com/explodes/serving/statusz"
	"github.com/fatih/color"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	configFlag = flag.String("config", "config.textproto", "configuration file location")
	colorFlag  = flag.Bool("color", true, "whether or not to print in colors")
)

var (
	compactTextMarshaller = proto.TextMarshaler{Compact: true, ExpandAny: true}
)

var colors = []color.Attribute{
	color.FgRed,
	color.FgGreen,
	color.FgYellow,
	color.FgBlue,
	color.FgMagenta,
	color.FgCyan,
}

func init() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(colors), func(i, j int) {
		colors[i], colors[j] = colors[j], colors[i]
	})
}

func main() {
	flag.Parse()

	// Enable/disable color output
	color.NoColor = !(*colorFlag)

	config := &monitor.Config{}
	if err := serving.ReadConfigFile(*configFlag, config); err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	if len(config.Services) == 0 {
		log.Fatal("no services configured for monitoring")
	}

	ctx := context.Background()

	for index, service := range config.Services {
		colorIndex := index % len(colors)
		go monitorService(ctx, colors[colorIndex], service)
	}

	<-ctx.Done()
}

func monitorService(ctx context.Context, fgColor color.Attribute, service *monitor.Config_Service) {
	addr := service.Address.Address()
	name := fmt.Sprintf("%s/%s:", addr, service.Name)

	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	var client statusz.StatuszServiceClient

	colorBuff := color.New(fgColor)

	for {
		buf.Reset()

		if client == nil {
			var err error
			client, err = makeClient(service)
			if err != nil {
				printf(buf, "%s ERROR %v", name, err)
				client = nil
			}
		}
		if client != nil {
			reqCtx, _ := context.WithTimeout(ctx, service.Timeout.Duration())
			now := time.Now()
			res, err := client.GetStatus(reqCtx, &statusz.GetStatusRequest{})
			then := time.Now()
			if err != nil {
				printf(buf, "%s ERROR %v", name, err)
				client = nil
			} else {
				printf(buf, "%s (%s) ", name, then.Sub(now))
				if err := compactTextMarshaller.Marshal(buf, res.Status); err != nil {
					log.Fatalf("error print metric: %v", err)
				}
				buf.WriteRune('\n')
			}
		}

		colorPrint(colorBuff, buf)

		select {
		case <-time.After(service.UpdateFrequency.Duration()):
			continue
		case <-ctx.Done():
			break
		}
	}
}

func printf(buf io.Writer, msg string, args ...interface{}) {
	_, err := fmt.Fprintf(buf, msg, args...)
	if err != nil {
		panic(fmt.Errorf("buffer error: %v", err))
	}
}

func colorPrint(c *color.Color, buf fmt.Stringer) {
	_, err := c.Fprint(os.Stderr, buf.String())
	if err != nil {
		panic(fmt.Errorf("buffer error: %v", err))
	}
}

func makeClient(service *monitor.Config_Service) (statusz.StatuszServiceClient, error) {
	conn, err := grpc.Dial(service.Address.Address(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return statusz.NewStatuszServiceClient(conn), nil
}
