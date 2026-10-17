package service

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/azrilpramudia/libre-go/internal/domain/member"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrEmailTaken         = errors.New("email already registered")
)

type AuthService struct {
	memberRepo member.Repository
}

func NewAuthService(memberRepo member.Repository) *AuthService {
	return &AuthService{memberRepo: memberRepo}
}

func (s *AuthService) Register(ctx context.Context, fullName, email, phone, password string) (*member.Member, error) {
	_, err := s.memberRepo.GetByEmail(ctx, email)
	if err == nil {
		return nil, ErrEmailTaken
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return s.memberRepo.Create(ctx, &member.Member{
		FullName:     fullName,
		Email:        email,
		Phone:        phone,
		PasswordHash: string(hash),
	})
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	m, err := s.memberRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(m.PasswordHash), []byte(password)); err != nil {
		return "", ErrInvalidCredentials
	}

	return generateToken(m.ID.String())
}

func generateToken(memberID string) (string, error) {
	claims := jwt.MapClaims{
		"member_id": memberID,
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
