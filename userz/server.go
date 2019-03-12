package userz

import (
	"github.com/explodes/serving/expz"
	"github.com/explodes/serving/logz"
	spb "github.com/explodes/serving/proto"
	"github.com/explodes/serving/statusz"
	"github.com/explodes/serving/utilz"
	"github.com/golang/protobuf/proto"
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
	cookiePasscode string
	storage        Storage
	logz           logz.Client
	expz           expz.Client
}

type deps struct {
	frame *logz.Frame
	exps  expz.ExperimentFlags
	log   logz.DeferredLog
}

var (
	registerOnce = &sync.Once{}
	varValidate = statusz.NewRateTracker("Validate")
	varLogin    = statusz.NewRateTracker("Login")
)

func registerServerVars() {
	statusz.Register("Userz", statusz.VarGroup{
		varValidate,
		varLogin,
	})
}

func NewUserzServer(cookiePasscode string, storage Storage, logz logz.Client, expz expz.Client) UserzServiceServer {
	registerOnce.Do(registerServerVars)
	return &userzServer{
		cookiePasscode: cookiePasscode,
		storage:        storage,
		logz:           logz,
		expz:           expz,
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
	var exps expz.ExperimentFlags
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

	uid, err := s.storage.UserID(requestContext, req.Username, req.Password)
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

	serialCookie, err := s.serializeCookie(cookie)
	if err != nil {
		s.logz.Errorf(deps.frame, "unable to serialize cookie: %v", err)
		return nil, errors.New("login error")
	}

	err = s.storage.Save(requestContext, serialCookie, cookie)
	if err != nil {
		s.logz.Errorf(deps.frame, "unable to save cookie: %v", err)
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

func (s *userzServer) Validate(requestContext context.Context, req *ValidateRequest) (*ValidateResponse, error) {
	defer varValidate.Start().End()

	deps, err := s.getDeps(requestContext, "Userz.Validate", req.Cookie)
	if err != nil {
		return nil, err
	}
	defer deps.log.Send()

	valid, err := s.storage.Validate(requestContext, req.Cookie)
	if err != nil {
		s.logz.Errorf(deps.frame, "error validating cookie with storage: %v", err)
		return nil, errors.Wrap(err, "validation error")

	}
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

func (s *userzServer) serializeCookie(cookie *spb.Cookie) (string, error) {
	b, err := proto.Marshal(cookie)
	if err != nil {
		return "", errors.Wrap(err, "error serializing cookie")
	}
	serial, err := utilz.EncryptToBase64String(b, s.cookiePasscode)
	if err != nil {
		return "", errors.Wrap(err, "error encrypting cookie")
	}
	return serial, nil
}

func getUuid() string {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return id.String()
}
