// +build testing

package test_utils

import (
	"context"
	"time"
)

const (
	TestTimeout = 10 * time.Second
)

func TestContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), TestTimeout)
	return ctx
}
