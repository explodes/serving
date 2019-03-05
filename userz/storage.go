package userz

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	spb "github.com/explodes/serving/proto"
	"github.com/pkg/errors"
	"time"
)

type Storage interface {
	UserID(ctx context.Context, username, password string) (string, error)
	CreateUser(ctx context.Context, username, password string) error
	Validate(ctx context.Context, cookie string) (bool, error)
	Save(ctx context.Context, serial string, cookie *spb.Cookie) error
}

type postgresStorage struct {
	db     *sql.DB
	secret string
}

func NewPostgresStorage(db *sql.DB, secret string) Storage {
	return &postgresStorage{
		db:     db,
		secret: secret,
	}
}

func (s *postgresStorage) UserID(ctx context.Context, username, password string) (string, error) {
	hash := s.hash(password)
	var userID int64
	if err := s.db.QueryRowContext(ctx, `SELECT id FROM users WHERE username = $1 AND password = $2`, username, hash).Scan(&userID); err != nil {
		return "", errors.Wrap(err, "unable to query user")
	}
	return fmt.Sprint(userID), nil
}

func (s *postgresStorage) CreateUser(ctx context.Context, username, password string) error {
	hash := s.hash(password)
	if _, err := s.db.ExecContext(ctx, `INSERT INTO users (username, password) VALUES ($1, $2)`, username, hash); err != nil {
		return errors.Wrap(err, "unable to save session")
	}
	return nil
}

func (s *postgresStorage) hash(v string) string {
	h := sha256.New()
	h.Write([]byte(s.secret + v))
	return hex.EncodeToString(h.Sum(nil))
}

func (s *postgresStorage) Validate(ctx context.Context, cookie string) (bool, error) {
	now := s.expiry(time.Now())
	var found int64
	if err := s.db.QueryRowContext(ctx, `SELECT 1 FROM sessions WHERE cookie = $1 AND expiry > $2`, cookie, now).Scan(&found); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, errors.Wrap(err, "unable to query session")
	}
	return true, nil
}

func (s *postgresStorage) Save(ctx context.Context, serial string, cookie *spb.Cookie) error {
	expiry := s.expiry(cookie.ExpirationTime.Time())
	if _, err := s.db.ExecContext(ctx, `INSERT INTO sessions (cookie, expiry) VALUES ($1, $2)`, serial, expiry); err != nil {
		return errors.Wrap(err, "unable to save session")
	}
	return nil
}

func (s *postgresStorage) expiry(t time.Time) int64 {
	return t.Unix()
}
