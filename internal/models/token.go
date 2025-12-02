package models

import "time"

type Token struct {
	ID           int    `gorm:"primaryKey;autoIncrement"`
	UserID       int    `gorm:"not null;index"`
	RefreshToken string `gorm:"unique;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (Token) TableName() string {
	return "tokens"
}
