package auth

import (
	"GraphNeo4jGO/config"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type ServiceImpl struct {
	cfg config.Secrets
}

type jwtClaims struct {
	username string
	id       uint
	jwt.RegisteredClaims
}

func (s *ServiceImpl) GenerateToken(id uint, username string) (string, error) {
	expTime := time.Now().Add(s.cfg.ExpTime)
	claims := &jwtClaims{
		username: username,
		id:       id,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	jc := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return jc.SignedString([]byte(s.cfg.JwtSecret))
}

func (s *ServiceImpl) ClaimsFromToken(t string) (any, error) {
	token, err := jwt.ParseWithClaims(t, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	_, ok := token.Claims.(*jwtClaims)
	if !ok {
		return nil, fmt.Errorf("cant get claims from token")
	}

	if err = token.Claims.Valid(); err != nil {
		return nil, err
	}

	return token.Claims, nil
}

func New(cfg config.Secrets) *ServiceImpl {
	return &ServiceImpl{
		cfg: cfg,
	}
}
