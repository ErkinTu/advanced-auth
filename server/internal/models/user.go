package models

import "time"

type User struct {
	ID             int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Email          string    `gorm:"unique;not null" json:"email"`
	PasswordHash   string    `gorm:"not null" json:"-"`
	IsActivated    bool      `gorm:"default:false" json:"is_activated"`
	ActivationLink string    `gorm:"column:activation_link" json:"-"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	Roles []Role `gorm:"many2many:user_roles;" json:"roles"`
}

func (User) TableName() string {
	return "users"
}
