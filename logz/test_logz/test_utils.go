// +build testing

package test_logz

import (
	"github.com/explodes/serving/logz"
	"github.com/explodes/serving/logz/mock_logz"
	"github.com/golang/mock/gomock"
)

type noopDeferredLog struct{}

func (d *noopDeferredLog) Send() {}

func NoopDeferredLog() logz.DeferredLog {
	return &noopDeferredLog{}
}

func NoopLogzClient(ctrl *gomock.Controller) *mock_logz.MockClient {
	mockLogz := mock_logz.NewMockClient(ctrl)
	mockLogz.EXPECT().Debug(gomock.Any(), gomock.Any()).AnyTimes()
	mockLogz.EXPECT().Debugf(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	mockLogz.EXPECT().Info(gomock.Any(), gomock.Any()).AnyTimes()
	mockLogz.EXPECT().Infof(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	mockLogz.EXPECT().Warn(gomock.Any(), gomock.Any()).AnyTimes()
	mockLogz.EXPECT().Warnf(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	mockLogz.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
	mockLogz.EXPECT().Errorf(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	mockLogz.EXPECT().DeferRequestLog(gomock.Any()).Return(NoopDeferredLog()).AnyTimes()
	mockLogz.EXPECT().Defer(gomock.Any(), gomock.Any(), gomock.Any()).Return(NoopDeferredLog()).AnyTimes()
	return mockLogz
}
