package logz

import (
	"log"
	"time"
)

type Backend interface {
	Record(request *RecordRequest) (err error)
}

type consoleBackend struct {
}

func NewConsoleBackend() Backend {
	return &consoleBackend{}
}

func (cb *consoleBackend) Record(request *RecordRequest) (err error) {
	for _, entry := range request.Entries {
		cb.recordEntry(request.Stack, entry)
	}
	return nil
}

func (cb *consoleBackend) recordEntry(stack *Frame, entry *Entry) {
	if entry.EndTimestamp == nil {
		log.Printf("level=%-5s id=%-36s parent=%-36s operation=%-32s message=%s\n",
			entry.Level,
			stack.OperationId,
			stack.ParentOperationId,
			stack.OperationName,
			entry.Message,
		)
	} else {
		log.Printf("level=%-5s id=%-36s parent=%-36s operation=%-32s duration=%-10s message=%s\n",
			entry.Level,
			stack.OperationId,
			stack.ParentOperationId,
			stack.OperationName,
			time.Duration(entry.EndTimestamp.GetNanoseconds()-entry.Timestamp.GetNanoseconds()),
			entry.Message,
		)
	}
}
