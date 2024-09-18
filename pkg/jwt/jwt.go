package jwt

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TokenManager provides logic for JWT & Refresh tokens generation and parsing.
type TokenManager interface {
	NewJWT(payload Payload, createdAt time.Time) (string, error)
	Parse(accessToken string, expiredDuration time.Duration, payload Payload) error
	NewRefreshToken() (string, error)
	RefreshTokenExpired(expiredTime *time.Time) error
}

type Manager struct {
	signingKey string
}

func NewManager(signingKey string) (*Manager, error) {
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}

	return &Manager{signingKey: signingKey}, nil
}

type Payload interface {
	ToString() string
	ParseFromString(payloadString string)
}

func (m *Manager) NewJWT(payload Payload, createdAt time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:  payload.ToString(),
		IssuedAt: &jwt.NumericDate{Time: createdAt},
	})

	return token.SignedString([]byte(m.signingKey))
}

func (m *Manager) Parse(accessToken string, expiredDuration time.Duration, payload Payload) error {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.signingKey), nil
	})
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("error get claims from token")
	}

	issuedAd, err := claims.GetIssuedAt()
	if err != nil {
		return fmt.Errorf("error get issued at field from claims ")
	}

	payload.ParseFromString(claims["sub"].(string))
	if issuedAd.Add(expiredDuration).Before(time.Now()) {
		return jwt.ErrTokenExpired
	}

	return nil
}

func (m *Manager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

func (m *Manager) RefreshTokenExpired(expiredTime *time.Time) error {
	if expiredTime.Before(time.Now().UTC()) {
		return jwt.ErrTokenExpired
	}
	return nil
}
