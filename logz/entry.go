package logz

import (
	spb "github.com/explodes/serving/proto"
)

func NewEntry(level Level, message string) *Entry {
	return &Entry{
		Timestamp: spb.TimestampNow(),
		Level:     level,
		Message:   message,
	}
}
