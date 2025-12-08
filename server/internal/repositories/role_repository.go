package repositories

import (
	"AdvAuthGo/internal/models"

	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(role *models.Role) error
	GetByID(id int) (*models.Role, error)
	GetByName(name string) (*models.Role, error)
	GetAll() ([]models.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository { return &roleRepository{db: db} }

func (r roleRepository) Create(role *models.Role) error {
	return r.db.Create(&models.Role{Name: role.Name}).Error
}

func (r roleRepository) GetByID(id int) (*models.Role, error) {
	var role models.Role
	err := r.db.First(&role, id).Error
	return &role, err
}

func (r roleRepository) GetByName(name string) (*models.Role, error) {
	var role models.Role
	err := r.db.Where("name = ?", name).First(&role).Error
	return &role, err
}

func (r roleRepository) GetAll() ([]models.Role, error) {
	roles := []models.Role{}
	err := r.db.Find(&roles).Error
	return roles, err
}
