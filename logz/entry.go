package logz

import (
	spb "github.com/explodes/serving/proto"
)

func NewEntry(level Level, message string) *Entry {
	return &Entry{
		Timestamp: spb.Now(),
		Level:     level,
		Message:   message,
	}
}
