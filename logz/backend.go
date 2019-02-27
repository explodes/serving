package logz

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Backend interface {
	Record(request *LogRequest) (err error)
}

type ConsoleBackendOption func(cb *consoleBackend)

func WithWriter(out io.Writer) ConsoleBackendOption {
	return func(cb *consoleBackend) {
		cb.out = out
	}
}

type consoleBackend struct {
	out io.Writer
}

func NewConsoleBackend(opts ...ConsoleBackendOption) Backend {
	cb := &consoleBackend{
		out: os.Stderr,
	}
	for _, opt := range opts {
		opt(cb)
	}
	return cb
}

func (cb *consoleBackend) Record(request *LogRequest) (err error) {
	for _, entry := range request.LogEntries {
		if err := cb.recordEntry(request.Stack, entry); err != nil {
			return err
		}
	}
	return nil
}

func (cb *consoleBackend) recordEntry(stack *Frame, entry *LogEntry) error {
	_, err := fmt.Fprintf(cb.out, "%s %s:%s %s %s %s %s\n",
		time.Now(),
		stack.Id.Name,
		stack.Id.Instance,
		entry.Level,
		entry.Timestamp.Time(),
		entry.Type,
		entry.Message,
	)
	return err
}
