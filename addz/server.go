package addz

import (
	"context"
	"github.com/explodes/serving/expz"
	"github.com/explodes/serving/logz"
	"github.com/explodes/serving/statusz"
	"github.com/pkg/errors"
	"time"
)

var (
	varAdd      = statusz.NewRateTracker("Add")
	varSubtract = statusz.NewRateTracker("Subtract")
)

func registerAddzStatusz() {
	statusz.Register("Addz", statusz.VarGroup{
		varAdd,
		varSubtract,
	})
}

type addzServer struct {
	expz *expz.Client
	logz *logz.Client
}

func NewAddzServer(logz *logz.Client, expz *expz.Client) AddzServiceServer {
	registerAddzStatusz()
	return &addzServer{
		expz: expz,
		logz: logz,
	}
}

type deps struct {
	frame *logz.Frame
	exps  *expz.ExperimentFlags
	log   *logz.DeferredLog
}

func (s *addzServer) mathDeps(requestContext context.Context, operation string, cookie string) (*deps, error) {
	frame, err := logz.FrameFromIncomingContext(requestContext)
	if err != nil {
		s.logz.Errorf(frame, "error getting frame: %v", err)
		return nil, errors.Wrap(err, "unable to read incoming frame")
	}
	frame = logz.FrameForOutgoingContext(frame, operation)
	log := s.logz.Defer(frame, logz.Level_INFO, "request")

	expzCtx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	expzCtx, err = logz.PutFrameInOutgoingContext(expzCtx, frame)
	if err != nil {
		s.logz.Errorf(frame, "unable to update frame: %v", err)
		return nil, errors.Wrap(err, "unable to update frame")
	}
	exps, err := s.expz.GetExperiments(expzCtx, cookie)
	if err != nil {
		s.logz.Errorf(frame, "unable to get experiments: %v", err)
		return nil, errors.Wrap(err, "unable to get experiments")
	}
	d := &deps{
		frame: frame,
		exps:  exps,
		log:   log,
	}
	return d, nil
}

func (s *addzServer) Add(requestContext context.Context, req *AddRequest) (*AddResponse, error) {
	defer varAdd.Start().End()

	deps, err := s.mathDeps(requestContext, "Addz.Add", req.Cookie)
	if err != nil {
		return nil, err
	}
	defer deps.log.Send()

	extraAddition := deps.exps.Int64Value("extra_addition", 0)
	s.logz.Debugf(deps.frame, "extra_addition=%d", extraAddition)

	sum := extraAddition
	for _, v := range req.Values {
		sum += v
	}

	res := &AddResponse{
		Result: sum,
	}
	return res, nil
}

func (s *addzServer) Subtract(requestContext context.Context, req *SubtractRequest) (*SubtractResponse, error) {
	defer varSubtract.Start().End()

	deps, err := s.mathDeps(requestContext, "Addz.Subtract", req.Cookie)
	if err != nil {
		return nil, err
	}
	defer deps.log.Send()

	var sum int64
	for _, v := range req.Values {
		sum -= v
	}

	res := &SubtractResponse{
		Result: sum,
	}
	return res, nil
}
