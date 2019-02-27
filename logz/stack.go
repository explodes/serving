package logz

import (
	"context"
	"errors"
	spb "github.com/explodes/serving/proto"
	"github.com/satori/go.uuid"
	"google.golang.org/grpc/metadata"
)

const (
	frameMetadataKey = "logz-frame"
)

var (
	errFrameMetadataNotFound = errors.New("frame metadata not found")
)

// PutNewFrameInOutgoingContext creates a new Frame and puts it into a Context for outgoing RPCs.
func PutNewFrameInOutgoingContext(ctx context.Context, parent *Frame, operationName string) (context.Context, error) {
	frame := FrameForOutgoingContext(parent, operationName)
	return PutFrameInOutgoingContext(ctx, frame)
}

// PutFrameInOutgoingContext puts a Frame into a Context for outgoing RPCs.
func PutFrameInOutgoingContext(ctx context.Context, frame *Frame) (context.Context, error) {
	s, err := spb.SerializeProtoBase64(frame)
	if err != nil {
		return nil, err
	}
	ctx = metadata.AppendToOutgoingContext(ctx, frameMetadataKey, s)
	return ctx, nil
}

// FrameForOutgoingContext creates a Frame for outgoing RPCs.
func FrameForOutgoingContext(parent *Frame, operationName string) *Frame {
	var stackID, parentOperationID string
	if parent == nil {
		stackID = getUuid()
		parentOperationID = ""

	} else {
		stackID = parent.StackId
		parentOperationID = parent.OperationId
	}
	return &Frame{
		StackId:           stackID,
		OperationId:       getUuid(),
		ParentOperationId: parentOperationID,
		OperationName:     operationName,
	}
}

// FrameFromIncomingContext gets a stack from incoming RPCs so that it can be logged or forwarded to
// outgoing RPCs as a parent.
func FrameFromIncomingContext(ctx context.Context) (*Frame, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, nil
	}
	serializedFrame, ok := md[frameMetadataKey]
	if !ok || len(serializedFrame) == 0 {
		return nil, nil
	}
	frame := &Frame{}
	err := spb.DeserializeProtoBase64(serializedFrame[0], frame)
	return frame, err
}

func getUuid() string {
	id, err := uuid.NewV4()
	if err != nil {
		return ""
	}
	return id.String()
}
