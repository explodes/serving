package userz

import (
	"github.com/explodes/serving/expz"
	"github.com/explodes/serving/logz"
	spb "github.com/explodes/serving/proto"
	"github.com/explodes/serving/statusz"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"sync"
	"time"
)

const (
	cookieValidity = 7 * 24 * time.Hour
)

type userzServer struct {
	logz     *logz.Client
	expz     *expz.Client
	sessions *sessions
}

type deps struct {
	frame *logz.Frame
	exps  *expz.ExperimentFlags
	log   *logz.DeferredLog
}

var (
	varValidate = statusz.NewRateTracker("Validate")
	varLogin    = statusz.NewRateTracker("Login")
)

func registerExpzStatusz() {
	statusz.Register("Userz", statusz.VarGroup{
		varValidate,
		varLogin,
	})
}

func NewUserzServer(logz *logz.Client, expz *expz.Client) UserzServiceServer {
	registerExpzStatusz()
	return &userzServer{
		logz:     logz,
		expz:     expz,
		sessions: newSessions(),
	}
}

func (s *userzServer) getDeps(requestContext context.Context, operation string, cookie string) (*deps, error) {
	frame, err := logz.FrameFromIncomingContext(requestContext)
	if err != nil {
		s.logz.Errorf(frame, "error getting frame: %v", err)
		return nil, errors.Wrap(err, "unable to read incoming frame")
	}
	frame = logz.FrameForOutgoingContext(frame, operation)
	log := s.logz.DeferRequestLog(frame)

	expzCtx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	expzCtx, err = logz.PutFrameInOutgoingContext(expzCtx, frame)
	if err != nil {
		s.logz.Errorf(frame, "unable to update frame: %v", err)
		return nil, errors.Wrap(err, "unable to update frame")
	}
	var exps *expz.ExperimentFlags
	if cookie != "" {
		exps, err = s.expz.GetExperiments(expzCtx, cookie)
		if err != nil {
			s.logz.Errorf(frame, "unable to get experiments: %v", err)
			return nil, errors.Wrap(err, "unable to get experiments")
		}
	}
	d := &deps{
		frame: frame,
		exps:  exps,
		log:   log,
	}
	return d, nil
}

func (s *userzServer) Login(requestContext context.Context, req *LoginRequest) (*LoginResponse, error) {
	defer varLogin.Start().End()

	deps, err := s.getDeps(requestContext, "Userz.Login", "")
	if err != nil {
		return nil, err
	}
	defer deps.log.Send()

	uid, err := s.userId(requestContext, req.Username, req.Password)
	if err != nil {
		s.logz.Errorf(deps.frame, "unable to get user id: %v", err)
		res := &LoginResponse{
			Result: &LoginResponse_Failure{
				Failure: &LoginResponse_LoginFailure{
					Reason: LoginResponse_LoginFailure_BAD_LOGIN,
				},
			},
		}
		return res, nil
	}

	now := time.Now()
	cookie := &spb.Cookie{
		ExpirationTime: spb.TimestampTime(now.Add(cookieValidity)),
		CreationTime:   spb.TimestampTime(now),
		UserId:         uid,
		SessionId:      getUuid(),
	}
	serialCookie, err := spb.SerializeCookie(cookie)
	s.sessions.save(serialCookie, cookie)
	if err != nil {
		s.logz.Errorf(deps.frame, "unable to serialize cookie: %v", err)
		return nil, errors.New("login error")
	}

	res := &LoginResponse{
		Result: &LoginResponse_Success{
			Success: &LoginResponse_LoginSuccess{
				Cookie: serialCookie,
			},
		},
	}
	return res, nil
}

func (s *userzServer) userId(ctx context.Context, username, password string) (string, error) {
	if username != "test" || password != "test" {
		return "", errors.New("invalid login")
	}
	return "1.0", nil
}

func (s *userzServer) Validate(requestContext context.Context, req *ValidateRequest) (*ValidateResponse, error) {
	defer varValidate.Start().End()

	deps, err := s.getDeps(requestContext, "Userz.Validate", req.Cookie)
	if err != nil {
		return nil, err
	}
	defer deps.log.Send()

	valid := s.sessions.exists(req.Cookie)
	var result ValidateResponse_ValidateResult
	if valid {
		result = ValidateResponse_SUCCESS
	} else {
		result = ValidateResponse_INVALID
	}

	res := &ValidateResponse{
		Result: result,
	}
	return res, nil
}

type sessions struct {
	mu *sync.RWMutex
	m  map[string]*spb.Cookie
}

func newSessions() *sessions {
	return &sessions{
		mu: &sync.RWMutex{},
		m:  make(map[string]*spb.Cookie),
	}
}

func (s *sessions) exists(serial string) bool {
	s.mu.RLock()
	cookie, exists := s.m[serial]
	s.mu.RUnlock()

	expired := cookie.ExpirationTime != nil && cookie.ExpirationTime.Time().Before(time.Now())
	if expired {
		s.mu.Lock()
		delete(s.m, serial)
		s.mu.Unlock()
	}

	return exists && !expired
}

func (s *sessions) save(serial string, cookie *spb.Cookie) {
	s.mu.Lock()
	s.m[serial] = cookie
	s.mu.Unlock()
}

func getUuid() string {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return id.String()
}
