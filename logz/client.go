package logz

import (
	"context"
	spb "github.com/explodes/serving/proto"
	"google.golang.org/grpc"
	"log"
	"sync"
)

type frameEntry struct {
	entry *Entry
	frame *Frame
}

type Client struct {
	clientMu *sync.RWMutex
	addr     *spb.Address
	conn     *grpc.ClientConn
	logz     LogzServiceClient
	entries  chan frameEntry
}

func NewClient(addr *spb.Address) (*Client, error) {
	client := &Client{
		clientMu: &sync.RWMutex{},
		addr:     addr,
		entries:  make(chan frameEntry),
	}
	err := client.restoreClient()
	if err != nil {
		return nil, err
	}
	go client.loop()
	return client, nil
}

func (c *Client) loop() {
	for frameEntry := range c.entries {
		_, err := c.logz.Record(context.Background(), &RecordRequest{
			Stack:   frameEntry.frame,
			Entries: []*Entry{frameEntry.entry},
		})
		if err != nil {
			log.Printf("error sending log: %v", err)
		}
	}
}

func (c *Client) restoreClient() error {
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return err
		}
		c.conn = nil
		c.logz = nil
	}
	conn, err := grpc.Dial(c.addr.Address(), grpc.WithInsecure())
	if err != nil {
		return err
	}
	c.conn = conn
	c.logz = NewLogzServiceClient(conn)
	return nil
}

func (c *Client) makeEntry(level Level, message string) *Entry {
	return &Entry{
		Level:     level,
		Message:   message,
		Timestamp: spb.Now(),
	}
}

func (c *Client) Log(frame *Frame, level Level, message string) {
	entry := c.makeEntry(level, message)
	c.LogEntry(frame, entry)
}

func (c *Client) LogEntry(frame *Frame, entry *Entry) {
	go func() {
		c.entries <- frameEntry{frame: frame, entry: entry}
	}()
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
	d.frameEntry.entry.EndTimestamp = spb.Now()
	d.logz.LogEntry(d.frameEntry.frame, d.frameEntry.entry)
}
