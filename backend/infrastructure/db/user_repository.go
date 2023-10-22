package db

import (
	"github.com/Montekrysto/ecomm/domain/entity"
	"github.com/Montekrysto/ecomm/domain/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (repo *UserRepositoryImpl) Create(user entity.User) error {
	err := repo.db.Create(&user).Error

	if err != nil {
		// Handle specific errors as per your requirements, for example:
		// if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
		// 	return fmt.Errorf("email already in use")
		// }
		return err
	}

	return nil
}

func (repo *UserRepositoryImpl) GetAll() ([]entity.User, error) {
	var users []entity.User
	err := repo.db.Find(&users).Error
	return users, err
}

func (repo *UserRepositoryImpl) GetByID(id uuid.UUID) (entity.User, error) {
	var user entity.User
	err := repo.db.First(&user, "id = ?", id).Error
	return user, err
}

func (repo *UserRepositoryImpl) Update(user entity.User) error {
	return repo.db.Save(&user).Error
}

func (repo *UserRepositoryImpl) Delete(id uuid.UUID) error {
	return repo.db.Delete(&entity.User{}, "id = ?", id).Error
}
