package expz

import (
	"github.com/explodes/serving/logz"
	"github.com/explodes/serving/statusz"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type expzServer struct {
	logz *logz.Client
	exps Experiments
}

var (
	varGetExperiments = statusz.NewRateTracker("GetExperiments")
)

func registerExpzStatusz() {
	statusz.Register("Expz", statusz.VarGroup{
		varGetExperiments,
	})
}

func NewExpzServer(logz *logz.Client, exps Experiments) ExpzServiceServer {
	registerExpzStatusz()
	return &expzServer{
		logz: logz,
		exps: exps,
	}
}

func (e *expzServer) GetExperiments(ctx context.Context, req *GetExperimentsRequest) (*GetExperimentsResponse, error) {
	defer varGetExperiments.Start().End()

	frame, err := logz.FrameFromIncomingContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read incoming frame")
	}
	frame = logz.FrameForOutgoingContext(frame, "Expz.GetExperiments")
	defer e.logz.Defer(frame, logz.Level_INFO, "request").Send()

	mod := req.Cookie % MaxMods
	flags := e.exps[mod]
	res := &GetExperimentsResponse{
		Features: &Features{
			Flags: flags,
		},
	}

	return res, nil
}
