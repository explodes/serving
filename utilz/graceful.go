package utilz

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

const (
	httpShutdownTimeout = 10 * time.Second
)

var (
	gracefulMu        = &sync.Mutex{}
	gracefulShutdowns = make([]shutdowner, 0, 16)
	gracefulCh        = make(chan struct{})
)

type shutdowner interface {
	Name() string
	Shutdown() error
}

func init() {
	registerHandler()
}

func registerHandler() {
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt, os.Kill)
	go func() {
		exitSignal := <-exitChan

		gracefulMu.Lock()
		defer gracefulMu.Unlock()

		nShutdowns := len(gracefulShutdowns)
		if nShutdowns == 1 {
			log.Printf("SHUTDOWN STARTED (%s): 1 system", exitSignal)
		} else {
			log.Printf("SHUTDOWN STARTED (%s): %d systems", exitSignal, nShutdowns)
		}
		wg := &sync.WaitGroup{}
		wg.Add(nShutdowns)
		for _, sd := range gracefulShutdowns {
			go func(sd shutdowner) {
				defer wg.Done()
				name := sd.Name()
				err := sd.Shutdown()
				if err == nil {
					log.Printf("SHUTDOWN SUCCESS: %s", name)
				} else {
					log.Printf("SHUTDOWN FAILURE: %s: %v", name, err)
				}
			}(sd)
		}
		wg.Wait()

		close(gracefulCh)

		// We need to re-emit the exit signal because a normal use case is that stuff will be run
		// in a goroutine and since it has hijacked the exit signal it must re-emit.
		log.Print("SHUTDOWN COMPLETE")
		signal.Stop(exitChan)
		if currentProcess, err := os.FindProcess(os.Getpid()); err != nil {
			log.Printf("Error getting current process to re-emit exit signal: %v", err)
		} else {
			if err := currentProcess.Signal(exitSignal); err != nil {
				log.Printf("Error re-emitting exit signal: %v", err)
			}
		}
	}()
}

func registerShutdowner(sd shutdowner) {
	gracefulMu.Lock()
	gracefulShutdowns = append(gracefulShutdowns, sd)
	gracefulMu.Unlock()
}

func GracefulShutdown() <-chan struct{} {
	return gracefulCh
}

type closeShutdowner struct {
	name string
	c    io.Closer
}

func (sd *closeShutdowner) Name() string {
	return sd.name
}

func (sd *closeShutdowner) Shutdown() error {
	return sd.c.Close()
}

func RegisterGracefulShutdownCloser(name string, c io.Closer) {
	sd := &closeShutdowner{
		name: name,
		c:    c,
	}
	registerShutdowner(sd)
}

type grpcShutdowner struct {
	name string
	gs   *grpc.Server
}

func (sd *grpcShutdowner) Name() string {
	return sd.name
}

func (sd *grpcShutdowner) Shutdown() error {
	sd.gs.GracefulStop()
	return nil
}

func RegisterGracefulShutdownGrpcServer(name string, gs *grpc.Server) {
	sd := &grpcShutdowner{
		name: name,
		gs:   gs,
	}
	registerShutdowner(sd)
}

type httpShutdowner struct {
	name string
	hs   *http.Server
}

func (sd *httpShutdowner) Name() string {
	return sd.name
}

func (sd *httpShutdowner) Shutdown() error {
	ctx, _ := context.WithTimeout(context.Background(), httpShutdownTimeout)
	return sd.hs.Shutdown(ctx)
}

func RegisterGracefulShutdownHttpServer(name string, hs *http.Server) {
	sd := &httpShutdowner{
		name: name,
		hs:   hs,
	}
	registerShutdowner(sd)
}
