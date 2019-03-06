package logz_test

import (
	"github.com/explodes/serving/logz"
	spb "github.com/explodes/serving/proto"
	"github.com/fatih/color"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConsoleBackend_Record(t *testing.T) {
	b := logz.NewConsoleBackend()

	ext, err := ptypes.MarshalAny(spb.TimestampNow())
	assert.NoError(t, err)

	err = b.Record(&logz.RecordRequest{
		Cookie: "foo",
		Frame: &logz.Frame{
			FrameId:       "bar",
			ParentFrameId: "baz",
			FrameName:     "bell",
			StackId:       "bang",
		},
		Entries: []*logz.Entry{
			{
				Timestamp:    spb.TimestampNow(),
				Level:        logz.Level_ERROR,
				Message:      "hello world",
				EndTimestamp: spb.TimestampNow(),
				Ext:          ext,
			},
		},
	})
	assert.NoError(t, err)
}

func TestConsoleBackend_Record_NoEndTimestamp(t *testing.T) {
	b := logz.NewConsoleBackend()
	err := b.Record(&logz.RecordRequest{
		Cookie: "foo",
		Frame: &logz.Frame{
			FrameId:       "bar",
			ParentFrameId: "baz",
			FrameName:     "bell",
			StackId:       "bang",
		},
		Entries: []*logz.Entry{
			logz.NewEntry(logz.Level_DEBUG, "hello, world"),
		},
	})
	assert.NoError(t, err)
}

func TestConsoleBackend_Record_Color(t *testing.T) {
	originalColor := color.NoColor
	color.NoColor = true
	defer func() {
		color.NoColor = originalColor
	}()
	b := logz.NewConsoleBackend()

	ext, err := ptypes.MarshalAny(spb.TimestampNow())
	assert.NoError(t, err)

	err = b.Record(&logz.RecordRequest{
		Cookie: "foo",
		Frame: &logz.Frame{
			FrameId:       "bar",
			ParentFrameId: "baz",
			FrameName:     "bell",
			StackId:       "bang",
		},
		Entries: []*logz.Entry{
			{
				Timestamp: spb.TimestampNow(),
				Level:     logz.Level_ERROR,
				Message:   "hello world",
				Ext:       ext,
			},
		},
	})
	assert.NoError(t, err)
}

func TestConsoleBackend_Record_NoColor(t *testing.T) {
	originalColor := color.NoColor
	color.NoColor = false
	defer func() {
		color.NoColor = originalColor
	}()
	b := logz.NewConsoleBackend()

	ext, err := ptypes.MarshalAny(spb.TimestampNow())
	assert.NoError(t, err)

	err = b.Record(&logz.RecordRequest{
		Cookie: "foo",
		Frame: &logz.Frame{
			FrameId:       "bar",
			ParentFrameId: "baz",
			FrameName:     "bell",
			StackId:       "bang",
		},
		Entries: []*logz.Entry{
			{
				Timestamp: spb.TimestampNow(),
				Level:     logz.Level_ERROR,
				Message:   "hello world",
				Ext:       ext,
			},
		},
	})
	assert.NoError(t, err)
}
