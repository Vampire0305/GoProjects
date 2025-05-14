package auth

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/sudarshanmg/gotask/pkg/validation"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(req RegisterRequest) (*User, error)
	Login(req LoginRequest) (string, error)
}

type authService struct {
	repo      AuthRepository
	jwtSecret string
	validator *validator.Validate
}

func NewService(repo AuthRepository, jwtSecret string) AuthService {
	return &authService{
		repo:      repo,
		jwtSecret: jwtSecret,
		validator: validator.New(),
	}
}

func (s *authService) Register(req RegisterRequest) (*User, error) {
	if err := s.validator.Struct(req); err != nil {
		return nil, validation.FormatValidationError(err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	id, err := s.repo.CreateUser(req.Username, string(hash))
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       id,
		Username: req.Username,
	}, nil
}

func (s *authService) Login(req LoginRequest) (string, error) {
	if err := s.validator.Struct(req); err != nil {
		return "", validation.FormatValidationError(err)
	}

	user, err := s.repo.FindByUsername(req.Username)
	if err != nil || user == nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(15 * time.Minute).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return signed, nil
}
