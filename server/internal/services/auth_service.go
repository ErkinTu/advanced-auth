package services

import (
	"AdvAuthGo/config"
	"AdvAuthGo/internal/models"
	"AdvAuthGo/internal/repositories"
	"AdvAuthGo/internal/utils"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthService interface {
	Register(email, password string) (*TokenPair, error)
	Login(email, password string) (*TokenPair, error)
	Activate(token string) error
	Refresh(refreshToken string) (*TokenPair, error)
	GetAllUsers() ([]models.User, error)
}

type authService struct {
	userRepo    repositories.UserRepository
	tokenRepo   repositories.TokenRepository
	emailSender *utils.EmailSender
	config      *config.Config
}

func NewAuthService(userRepo repositories.UserRepository, tokenRepo repositories.TokenRepository, emailSender *utils.EmailSender, cfg *config.Config) AuthService {
	return &authService{
		userRepo:    userRepo,
		tokenRepo:   tokenRepo,
		emailSender: emailSender,
		config:      cfg,
	}
}

func (s *authService) Register(email, password string) (*TokenPair, error) {
	existing, _ := s.userRepo.GetByEmail(email)
	if existing != nil && existing.ID != 0 {
		return nil, errors.New("user already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:          email,
		PasswordHash:   string(hashed),
		IsActivated:    false,
		ActivationLink: uuid.New().String(),
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	activationURL := fmt.Sprintf("http://localhost:8080/api/activate/%s", user.ActivationLink)

	err = s.emailSender.SendActivationEmail(email, activationURL)
	if err != nil {
		return nil, fmt.Errorf("failed to send email: %w", err)
	}

	return s.generateAndSaveTokens(user)
}

func (s *authService) Login(email, password string) (*TokenPair, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !user.IsActivated {
		return nil, errors.New("account not activated")
	}

	return s.generateAndSaveTokens(user)
}

func (s *authService) Activate(activationLink string) error {
	user, err := s.userRepo.GetByActivationLink(activationLink)
	if err != nil {
		return errors.New("invalid activation link")
	}

	user.IsActivated = true
	user.ActivationLink = ""
	return s.userRepo.Update(user)
}

func (s *authService) Refresh(refreshToken string) (*TokenPair, error) {
	token, err := s.tokenRepo.GetByToken(refreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	claims, err := utils.ValidateToken(refreshToken, s.config.JWTSecret)
	if err != nil {
		return nil, errors.New("invalid or expired token")
	}

	user, err := s.userRepo.GetByID(int(claims.UserId))
	if err != nil {
		return nil, errors.New("user not found")
	}

	s.tokenRepo.DeleteByUserID(token.UserID)

	return s.generateAndSaveTokens(user)
}

func (s *authService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAll()
}

func (s *authService) generateAndSaveTokens(user *models.User) (*TokenPair, error) {
	accessToken, err := utils.GenerateToken(uint(user.ID), user.Email, s.config.JWTSecret, 15*time.Minute)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateToken(uint(user.ID), user.Email, s.config.JWTSecret, 30*24*time.Hour)
	if err != nil {
		return nil, err
	}

	s.tokenRepo.DeleteByUserID(user.ID)

	token := &models.Token{
		UserID:       user.ID,
		RefreshToken: refreshToken,
	}
	if err := s.tokenRepo.Create(token); err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
