package services

import (
	"AdvAuthGo/config"
	"AdvAuthGo/internal/models"
	"AdvAuthGo/internal/repositories"
	"AdvAuthGo/internal/utils"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthService interface {
	Register(email, password, passwordConfirm string) (*TokenPair, error)
	Login(email, password string) (*TokenPair, error)
	Activate(token string) error
	Refresh(refreshToken string) (*TokenPair, error)
	GetAllUsers() ([]models.User, error)
	GetUserByToken(token string) (*models.User, error)
	AssignRoleToUser(userID, roleName string) error
	CreateRole(name string) error
	DeleteRole(name string) error
	GetAllRoles() ([]models.Role, error)
}

type authService struct {
	userRepo    repositories.UserRepository
	tokenRepo   repositories.TokenRepository
	emailSender *utils.EmailSender
	config      *config.Config
	roleRepo    repositories.RoleRepository
}

func NewAuthService(userRepo repositories.UserRepository, tokenRepo repositories.TokenRepository, emailSender *utils.EmailSender, cfg *config.Config, roleRepo repositories.RoleRepository) AuthService {
	return &authService{
		userRepo:    userRepo,
		tokenRepo:   tokenRepo,
		emailSender: emailSender,
		config:      cfg,
		roleRepo:    roleRepo,
	}
}

func (s *authService) Register(email, password, passwordConfirm string) (*TokenPair, error) {
	existing, _ := s.userRepo.GetByEmail(email)
	if existing != nil && existing.ID != 0 {
		return nil, errors.New("user already exists")
	}

	if password != passwordConfirm {
		return nil, errors.New("passwords don't match")
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
	accessToken, err := utils.GenerateToken(uint(user.ID), user.Email, utils.ExtractRoleNames(user.Roles), s.config.JWTSecret, 15*time.Minute)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateToken(uint(user.ID), user.Email, utils.ExtractRoleNames(user.Roles), s.config.JWTSecret, 30*24*time.Hour)
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

func (s *authService) GetUserByToken(token string) (*models.User, error) {
	claims, err := utils.ValidateToken(token, s.config.JWTSecret)
	if err != nil {
		return nil, errors.New("invalid or expired token")
	}

	return s.userRepo.GetByIDWithRoles(int(claims.UserId))
}

func (s *authService) AssignRoleToUser(userID, roleName string) error {
	role, err := s.roleRepo.GetByName(roleName)
	if err != nil {
		return err
	}

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return err
	}

	user, err := s.userRepo.GetByIDWithRoles(userIDInt)
	if err != nil {
		return err
	}

	for _, r := range user.Roles {
		if r.ID == role.ID {
			return errors.New("role already assigned")
		}
	}

	user.Roles = append(user.Roles, *role)
	return s.userRepo.Update(user)
}

func (s *authService) CreateRole(roleName string) error {
	existingRole, err := s.roleRepo.GetByName(roleName)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			existingRole = nil
		} else {
			return err
		}
	}

	if existingRole != nil {
		return errors.New("role already exists")
	}

	role := &models.Role{Name: roleName}
	return s.roleRepo.Create(role)
}

func (s *authService) DeleteRole(roleName string) error {
	existingRole, err := s.roleRepo.GetByName(roleName)

	if err != nil {
		return err
	}

	return s.roleRepo.Delete(existingRole)
}

func (s *authService) GetAllRoles() ([]models.Role, error) {
	return s.roleRepo.GetAll()
}
