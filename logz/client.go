package logz

import (
	"context"
	"fmt"
	spb "github.com/explodes/serving/proto"
	"log"
)

type frameEntry struct {
	entry *Entry
	frame *Frame
}

type Client struct {
	logz    LogzServiceClient
	entries chan frameEntry
	console Backend
}

func NewClient(logz LogzServiceClient) *Client {
	client := &Client{
		logz:    logz,
		entries: make(chan frameEntry),
		console: NewConsoleBackend(),
	}
	go client.loop()
	return client
}

func (c *Client) loop() {
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

func (c *Client) makeEntry(level Level, message string) *Entry {
	return &Entry{
		Level:     level,
		Message:   message,
		Timestamp: spb.TimestampNow(),
	}
}

func (c *Client) Debug(frame *Frame, message string) {
	c.log(frame, Level_DEBUG, message)
}

func (c *Client) Debugf(frame *Frame, message string, args ...interface{}) {
	c.log(frame, Level_DEBUG, fmt.Sprintf(message, args...))
}

func (c *Client) Info(frame *Frame, message string) {
	c.log(frame, Level_INFO, message)
}

func (c *Client) Infof(frame *Frame, message string, args ...interface{}) {
	c.log(frame, Level_INFO, fmt.Sprintf(message, args...))
}

func (c *Client) Warn(frame *Frame, message string) {
	c.log(frame, Level_WARN, message)
}

func (c *Client) Warnf(frame *Frame, message string, args ...interface{}) {
	c.log(frame, Level_WARN, fmt.Sprintf(message, args...))
}

func (c *Client) Error(frame *Frame, message string) {
	c.log(frame, Level_ERROR, message)
}

func (c *Client) Errorf(frame *Frame, message string, args ...interface{}) {
	c.log(frame, Level_ERROR, fmt.Sprintf(message, args...))
}

func (c *Client) log(frame *Frame, level Level, message string) {
	entry := c.makeEntry(level, message)
	c.queueEntry(frame, entry)
}

func (c *Client) queueEntry(frame *Frame, entry *Entry) {
	go func() {
		c.entries <- frameEntry{frame: frame, entry: entry}
	}()
}

func (c *Client) DeferRequestLog(frame *Frame) *DeferredLog {
	return c.Defer(frame, Level_INFO, "BuildResponse")
}

func (c *Client) Defer(frame *Frame, level Level, message string) *DeferredLog {
	entry := c.makeEntry(level, message)
	return &DeferredLog{
		logz:       c,
		frameEntry: &frameEntry{frame: frame, entry: entry},
	}
}

type DeferredLog struct {
	logz       *Client
	frameEntry *frameEntry
}

func (d *DeferredLog) Send() {
	d.frameEntry.entry.EndTimestamp = spb.TimestampNow()
	d.logz.queueEntry(d.frameEntry.frame, d.frameEntry.entry)
}
