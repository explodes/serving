package logz

import (
	"context"
	"fmt"
	spb "github.com/explodes/serving/proto"
	"log"
)

const (
	tagBuildResponse = "BuildResponse"
)

type Client interface {
	Debug(frame *Frame, message string)
	Debugf(frame *Frame, message string, args ...interface{})
	Info(frame *Frame, message string)
	Infof(frame *Frame, message string, args ...interface{})
	Warn(frame *Frame, message string)
	Warnf(frame *Frame, message string, args ...interface{})
	Error(frame *Frame, message string)
	Errorf(frame *Frame, message string, args ...interface{})
	DeferRequestLog(frame *Frame) DeferredLog
	Defer(frame *Frame, level Level, message string) DeferredLog
}

type frameEntry struct {
	entry *Entry
	frame *Frame
}

type clientImpl struct {
	logz    LogzServiceClient
	entries chan frameEntry
	console Backend
}

func NewClient(logz LogzServiceClient) Client {
	client := &clientImpl{
		logz:    logz,
		entries: make(chan frameEntry),
		console: NewConsoleBackend(),
	}
	go client.loop()
	return client
}

func (c *clientImpl) loop() {
	for frameEntry := range c.entries {
		req := &RecordRequest{
			Stack:   frameEntry.frame,
			Entries: []*Entry{frameEntry.entry},
		}
		if err := c.console.Record(req); err != nil {
			log.Printf("error logging to console: %v", err)
		}
		if _, err := c.logz.Record(context.Background(), req); err != nil {
			log.Printf("error sending log: %v", err)
		}
	}
}

func (c *clientImpl) makeEntry(level Level, message string) *Entry {
	return &Entry{
		Level:     level,
		Message:   message,
		Timestamp: spb.TimestampNow(),
	}
}

func (c *clientImpl) Debug(frame *Frame, message string) {
	c.log(frame, Level_DEBUG, message)
}

func (c *clientImpl) Debugf(frame *Frame, message string, args ...interface{}) {
	c.log(frame, Level_DEBUG, fmt.Sprintf(message, args...))
}

func (c *clientImpl) Info(frame *Frame, message string) {
	c.log(frame, Level_INFO, message)
}

func (c *clientImpl) Infof(frame *Frame, message string, args ...interface{}) {
	c.log(frame, Level_INFO, fmt.Sprintf(message, args...))
}

func (c *clientImpl) Warn(frame *Frame, message string) {
	c.log(frame, Level_WARN, message)
}

func (c *clientImpl) Warnf(frame *Frame, message string, args ...interface{}) {
	c.log(frame, Level_WARN, fmt.Sprintf(message, args...))
}

func (c *clientImpl) Error(frame *Frame, message string) {
	c.log(frame, Level_ERROR, message)
}

func (c *clientImpl) Errorf(frame *Frame, message string, args ...interface{}) {
	c.log(frame, Level_ERROR, fmt.Sprintf(message, args...))
}

func (c *clientImpl) log(frame *Frame, level Level, message string) {
	entry := c.makeEntry(level, message)
	c.queueEntry(frame, entry)
}

func (c *clientImpl) queueEntry(frame *Frame, entry *Entry) {
	go func() {
		c.entries <- frameEntry{frame: frame, entry: entry}
	}()
}

func (c *clientImpl) DeferRequestLog(frame *Frame) DeferredLog {
	return c.Defer(frame, Level_INFO, tagBuildResponse)
}

func (c *clientImpl) Defer(frame *Frame, level Level, message string) DeferredLog {
	entry := c.makeEntry(level, message)
	return &deferredLogImpl{
		logz:       c,
		frameEntry: &frameEntry{frame: frame, entry: entry},
	}
}

type DeferredLog interface {
	Send()
}

type deferredLogImpl struct {
	logz       *clientImpl
	frameEntry *frameEntry
}

func (d *deferredLogImpl) Send() {
	d.frameEntry.entry.EndTimestamp = spb.TimestampNow()
	d.logz.queueEntry(d.frameEntry.frame, d.frameEntry.entry)
}
