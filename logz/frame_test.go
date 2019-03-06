package logz

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPutNewFrameInOutgoingContext_sanity(t *testing.T) {
	ctx, err := PutNewFrameInOutgoingContext(context.Background(), newFrame(), "operation")
	assert.NoError(t, err)
	assert.NotNil(t, ctx)
}

func TestPutNewFrameInOutgoingContext_nilParent_sanity(t *testing.T) {
	ctx, err := PutNewFrameInOutgoingContext(context.Background(), nil, "operation")
	assert.NoError(t, err)
	assert.NotNil(t, ctx)
}

func TestFrameForOutgoingContext_nilParent(t *testing.T) {
	frame := FrameForOutgoingContext(nil, "op")
	assert.Equal(t, "", frame.ParentFrameId)
	assert.Equal(t, "op", frame.FrameName)
	assert.Len(t, frame.FrameId, 36)
	assert.Len(t, frame.StackId, 36)
}

func TestFrameForOutgoingContext_withParent(t *testing.T) {
	parent := newFrame()
	frame := FrameForOutgoingContext(parent, "op")
	assert.Equal(t, parent.FrameId, frame.ParentFrameId)
	assert.Equal(t, "op", frame.FrameName)
	assert.Len(t, frame.FrameId, 36)
	assert.Equal(t, parent.StackId, frame.StackId)
}

func TestFrameFromIncomingContext_sanity(t *testing.T) {
	ctx := context.Background()
	_, err := FrameFromIncomingContext(ctx)
	assert.NoError(t, err)
}
