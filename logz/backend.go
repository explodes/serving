package logz

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
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
		cb.recordEntry(request.Frame, entry)
	}
	return nil
}

func (cb *consoleBackend) recordEntry(frame *Frame, entry *Entry) {
	var message string
	if entry.EndTimestamp == nil {
		message = fmt.Sprintf("level=%-5s id=%-36s parent=%-36s operation=%-32s message=%s\n",
			entry.Level,
			frame.FrameId,
			frame.ParentFrameId,
			frame.FrameName,
			entry.Message,
		)
	} else {
		message = fmt.Sprintf("level=%-5s id=%-36s parent=%-36s operation=%-32s duration=%-10s message=%s\n",
			entry.Level,
			frame.FrameId,
			frame.ParentFrameId,
			frame.FrameName,
			time.Duration(entry.EndTimestamp.GetNanoseconds()-entry.Timestamp.GetNanoseconds()),
			entry.Message,
		)
	}

	if color.NoColor {
		log.Print(message)
		return
	}

	var c color.Attribute
	switch entry.Level {
	case Level_DEBUG:
		c = color.FgBlue
	case Level_INFO:
		c = color.FgGreen
	case Level_WARN:
		c = color.FgYellow
	case Level_ERROR:
		c = color.FgRed
	default:
		c = color.FgRed
	}
	_, err := color.New(c).Fprint(os.Stderr, message)
	if err != nil {
		log.Printf("error printing to stderr: %s", err)
	}
}
