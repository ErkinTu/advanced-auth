package repositories

import (
	"AdvAuthGo/internal/models"

	"gorm.io/gorm"
)

type TokenRepository interface {
	Create(token *models.Token) error
	GetByToken(token string) (*models.Token, error)
	GetByUserID(userID int) (*models.Token, error)
	Update(token *models.Token) error
	DeleteByUserID(userID int) error
}

type tokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepository{db: db}
}

func (r *tokenRepository) Create(token *models.Token) error {
	return r.db.Create(token).Error
}

func (r *tokenRepository) GetByToken(token string) (*models.Token, error) {
	var t models.Token
	err := r.db.Where("refresh_token = ?", token).First(&t).Error
	return &t, err
}

func (r *tokenRepository) GetByUserID(userID int) (*models.Token, error) {
	var token models.Token
	err := r.db.Where("user_id = ?", userID).First(&token).Error
	return &token, err
}

func (r *tokenRepository) Update(token *models.Token) error {
	return r.db.Save(token).Error
}

func (r *tokenRepository) DeleteByUserID(userID int) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.Token{}).Error
}
