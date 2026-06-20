package repositories

import (
	"project-management-be/models"
	"project-management-be/utils"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByID (id uint64) (*models.User, error)
	FindByPublicID (publicID string) (*models.User, error)
	GetAllPaginate (filter, sort string, limit, offset int) ([]models.User, int64, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) FindByID (id uint64) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepository) FindByPublicID (publicID string) (*models.User, error) {
	var user models.User
	err := r.db.Where("public_id = ?", publicID).First(&user).Error
	return &user, err
}

func (r *userRepository) GetAllPaginate (filter, sort string, limit, offset int) ([]models.User, int64, error) {
	var users []models.User

	db := r.db.Model(&models.User{})

	if filter != "" {
		filterPattern := "%" + filter + "%"
		db = db.Where("name ILIKE ? OR email ILIKE ?", filterPattern, filterPattern)
	}

	if sort == "-id" {
		sort = "-internal_id"
	} else if sort == "id" {
		sort = "internal_id"
	}

	count, err := utils.Paginate(db, limit, offset, sort, &users)

	return users, count, err
}