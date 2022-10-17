package auth

import (
	"GraphNeo4jGO/config"
	"GraphNeo4jGO/repository"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type ServiceImpl struct {
	cfg   config.Secrets
	cache repository.Cache
}

type JwtClaims struct {
	Username string
	ID       uint
	jwt.RegisteredClaims
}

func (s *ServiceImpl) GenerateToken(id uint, username string) (string, error) {
	expTime := time.Now().Add(s.cfg.ExpTime)
	claims := &JwtClaims{
		Username: username,
		ID:       id,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	jc := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return jc.SignedString([]byte(s.cfg.JwtSecret))
}

func (s *ServiceImpl) ClaimsFromToken(t string) (any, error) {
	_, found := s.cache.Get(t)
	if found {
		return nil, fmt.Errorf("token is banned")
	}
	token, err := jwt.ParseWithClaims(t, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	_, ok := token.Claims.(*JwtClaims)
	if !ok {
		return nil, fmt.Errorf("cant get claims from token")
	}

	if err = token.Claims.Valid(); err != nil {
		return nil, err
	}

	return token.Claims, nil
}

func (r *ServiceImpl) BlackList(t string) {
	r.cache.Set(t, struct{}{})
}

func New(cfg config.Secrets, repo repository.Repository) *ServiceImpl {
	return &ServiceImpl{
		cfg:   cfg,
		cache: repo.Cache(),
	}
}
