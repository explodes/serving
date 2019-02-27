package logz

import (
	"context"
	"errors"
	"github.com/explodes/serving"
	"google.golang.org/grpc/metadata"
)

const (
	stackMetadata = "stackz"
)

var (
	errStackMetadataNotFound = errors.New("stack metadata not found")
)

// PutNewStackInOutgoingContext creates a new Frame and puts it into a Context for outgoing RPCs.
func PutNewStackInOutgoingContext(ctx context.Context, parent *Frame, id *ServerID) (context.Context, error){
	stack := StackForOutgoingContext(parent, id)
	return PutStackInOutgoingContext(ctx, stack)
}

// PutStackInOutgoingContext puts a Frame into a Context for outgoing RPCs.
func PutStackInOutgoingContext(ctx context.Context, stack *Frame) (context.Context, error){
	s, err := serving.SerializeProtoBase64(stack)
	if err != nil {
		return nil, err
	}
	ctx = metadata.AppendToOutgoingContext(ctx, stackMetadata, s)
	return ctx, nil
}

// StackForOutgoingContext creates a Frame for outgoing RPCs.
func StackForOutgoingContext(parent *Frame, id *ServerID) *Frame {
	return &Frame{
		ServerId:     id,
		Parent: parent,
	}
}

// StackFromIncomingContext gets a stack from incoming RPCs so that it can be logged or forwarded to
// outgoing RPCs as a parent.
func StackFromIncomingContext(ctx context.Context) (*Stack, error){
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil,errStackMetadataNotFound
	}
	serializedStackStrings, ok := md[stackMetadata]
	if !ok || len(serializedStackStrings) == 0 {
		return nil,errStackMetadataNotFound
	}
	stack := &Stack{}
	err := serving.DeserializeProtoBase64(serializedStackStrings[0], stack)
	return stack, err
}
