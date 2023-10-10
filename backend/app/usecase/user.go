package usecase

import (
	"github.com/Montekrysto/ecomm/domain/entity"
	"github.com/Montekrysto/ecomm/domain/repository"
	"github.com/google/uuid"
)

type UserUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: repo,
	}
}

func (u *UserUseCase) Create(user entity.User) error {
	return u.userRepo.Create(user)
}

func (u *UserUseCase) GetAll() ([]entity.User, error) {
	return u.userRepo.GetAll()
}

func (u *UserUseCase) GetByID(id uuid.UUID) (entity.User, error) {
	return u.userRepo.GetByID(id)
}

func (u *UserUseCase) Update(user entity.User) error {
	return u.userRepo.Update(user)
}

func (u *UserUseCase) Delete(id uuid.UUID) error {
	return u.userRepo.Delete(id)
}
