package logz

import (
	"context"
	"google.golang.org/grpc"
	"sync"
	"testing"
)

type fakeLogzClient struct{}

func (f fakeLogzClient) Record(ctx context.Context, req *RecordRequest, opts ...grpc.CallOption) (*RecordResponse, error) {
	return nil, nil
}

func newLogzServiceClient() LogzServiceClient {
	return fakeLogzClient{}
}

type waitingBackend struct {
	wg *sync.WaitGroup
}

func newWaitingBackend(wg *sync.WaitGroup) *waitingBackend {
	return &waitingBackend{wg: wg}
}

func (w *waitingBackend) Record(request *RecordRequest) (err error) {
	w.wg.Done()
	return nil
}

func TestClientImpl_Debug_sanity(t *testing.T) {
	m := newLogzServiceClient()
	c := NewClient(m).(*clientImpl)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	be := newWaitingBackend(wg)
	c.console = be

	c.Debug(newFrame(), "")

	wg.Wait()
}

func TestClientImpl_Debugf_sanity(t *testing.T) {
	m := newLogzServiceClient()
	c := NewClient(m).(*clientImpl)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	be := newWaitingBackend(wg)
	c.console = be

	c.Debugf(newFrame(), "")

	wg.Wait()
}

func TestClientImpl_Info_sanity(t *testing.T) {
	m := newLogzServiceClient()
	c := NewClient(m).(*clientImpl)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	be := newWaitingBackend(wg)
	c.console = be

	c.Info(newFrame(), "")

	wg.Wait()
}

func TestClientImpl_Infof_sanity(t *testing.T) {
	m := newLogzServiceClient()
	c := NewClient(m).(*clientImpl)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	be := newWaitingBackend(wg)
	c.console = be

	c.Infof(newFrame(), "")

	wg.Wait()
}

func TestClientImpl_Warn_sanity(t *testing.T) {
	m := newLogzServiceClient()
	c := NewClient(m).(*clientImpl)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	be := newWaitingBackend(wg)
	c.console = be

	c.Warn(newFrame(), "")

	wg.Wait()
}

func TestClientImpl_Warnf_sanity(t *testing.T) {
	m := newLogzServiceClient()
	c := NewClient(m).(*clientImpl)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	be := newWaitingBackend(wg)
	c.console = be

	c.Warnf(newFrame(), "")

	wg.Wait()
}

func TestClientImpl_Error_sanity(t *testing.T) {
	m := newLogzServiceClient()
	c := NewClient(m).(*clientImpl)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	be := newWaitingBackend(wg)
	c.console = be

	c.Error(newFrame(), "")

	wg.Wait()
}

func TestClientImpl_Errorf_sanity(t *testing.T) {
	m := newLogzServiceClient()
	c := NewClient(m).(*clientImpl)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	be := newWaitingBackend(wg)
	c.console = be

	c.Errorf(newFrame(), "")

	wg.Wait()
}

func TestClientImpl_DeferRequestLog_sanity(t *testing.T) {
	m := newLogzServiceClient()
	c := NewClient(m).(*clientImpl)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	be := newWaitingBackend(wg)
	c.console = be

	c.DeferRequestLog(newFrame()).Send()

	wg.Wait()
}

func TestClientImpl_Defer_sanity(t *testing.T) {
	m := newLogzServiceClient()
	c := NewClient(m).(*clientImpl)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	be := newWaitingBackend(wg)
	c.console = be

	c.Defer(newFrame(), 0, "msg").Send()

	wg.Wait()
}
