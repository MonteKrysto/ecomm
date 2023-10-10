package repository

import (
	"github.com/Montekrysto/ecomm/domain/entity"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user entity.User) error
	GetAll() ([]entity.User, error)
	GetByID(id uuid.UUID) (entity.User, error)
	Update(user entity.User) error
	Delete(id uuid.UUID) error
}
