package expz

import (
	"github.com/explodes/serving/logz"
	"github.com/explodes/serving/statusz"
	"github.com/explodes/serving/userz"
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

func (s *expzServer) GetExperiments(ctx context.Context, req *GetExperimentsRequest) (*GetExperimentsResponse, error) {
	defer varGetExperiments.Start().End()
	frame, err := logz.FrameFromIncomingContext(ctx)
	if err != nil {
		s.logz.Errorf(frame, "error getting frame: %v", err)
		return nil, errors.Wrap(err, "unable to read incoming frame")
	}
	frame = logz.FrameForOutgoingContext(frame, "Expz.GetExperiments")
	defer s.logz.Defer(frame, logz.Level_INFO, "request").Send()

	hash, err := userz.CookieHash(req.Cookie)
	if err != nil {
		s.logz.Errorf(frame, "error deserializing cookie: %v", err)
		return nil, errors.New("cookie error")
	}

	mod := hash % MaxMods
	s.logz.Debugf(frame, "mod=%d", mod)
	flags := s.exps[mod]
	res := &GetExperimentsResponse{
		Features: &Features{
			Flags: flags,
		},
	}

	return res, nil
}
