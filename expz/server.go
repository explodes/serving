package expz

import (
	"github.com/explodes/serving/logz"
	spb "github.com/explodes/serving/proto"
	"github.com/explodes/serving/statusz"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"sync"
)

type expzServer struct {
	logz logz.Client
	mods ModFlags
}

var (
	registerOnce      = &sync.Once{}
	varGetExperiments = statusz.NewRateTracker("GetExperiments")
)

func registerServerVars() {
	statusz.Register("Expz", statusz.VarGroup{
		varGetExperiments,
	})
}

func NewExpzServer(logz logz.Client, mods ModFlags) ExpzServiceServer {
	registerOnce.Do(registerServerVars)
	return &expzServer{
		logz: logz,
		mods: mods,
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
	//spew.Dump(s.logz)
	defer s.logz.DeferRequestLog(frame).Send()

	hash, err := spb.CookieHash(req.Cookie)
	if err != nil {
		s.logz.Errorf(frame, "error hashing cookie: %v", err)
		return nil, errors.New("cookie error")
	}

	mod := hash % MaxMods
	s.logz.Debugf(frame, "mod=%d", mod)
	flags := s.mods[mod]
	res := &GetExperimentsResponse{
		Features: &Features{
			Flags: flags,
		},
	}

	return res, nil
}
