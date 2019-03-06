package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"github.com/explodes/serving"
	"github.com/explodes/serving/monitor"
	"github.com/explodes/serving/statusz"
	"github.com/explodes/serving/utilz"
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
	addr := service.GrpcServer.Address.Address()
	name := fmt.Sprintf("%s/%s:", addr, service.Name)

	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	var cc *connClient

	colorBuff := color.New(fgColor)

	updateNanos := service.UpdateFrequency.Duration().Nanoseconds()
	randomNanos := time.Nanosecond * time.Duration(rand.Int63n(updateNanos))
	<-time.After(randomNanos)

	for {
		buf.Reset()

		if cc == nil {
			var err error
			cc, err = makeClient(service)
			if err != nil {
				printf(buf, "%s ERROR %v", name, err)
				if err := cc.conn.Close(); err != nil {
					log.Printf("error closing client: %v", err)
				}
				cc = nil
			}
		}
		if cc != nil {
			reqCtx, _ := context.WithTimeout(ctx, service.Timeout.Duration())
			now := time.Now()
			res, err := cc.client.GetStatus(reqCtx, &statusz.GetStatusRequest{})
			then := time.Now()
			if err != nil {
				printf(buf, "%s ERROR %v", name, err)
				cc = nil
			} else {
				printf(buf, "%s (%s) ", name, then.Sub(now))
				if err := compactTextMarshaller.Marshal(buf, res.Status); err != nil {
					log.Printf("error printing metric: %v", err)
				}
			}
			buf.WriteRune('\n')
			colorPrint(colorBuff, buf)
		}

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

type connClient struct {
	conn   *grpc.ClientConn
	client statusz.StatuszServiceClient
}

func makeClient(service *monitor.Config_Service) (*connClient, error) {
	conn, err := utilz.DialGrpc(service.GrpcServer)
	if err != nil {
		return nil, err
	}
	client := statusz.NewStatuszServiceClient(conn)
	cc := &connClient{
		conn:   conn,
		client: client,
	}
	return cc, nil
}
