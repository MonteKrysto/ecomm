package usecase

import "github.com/google/uuid"

type CRUDInterface[T any] interface {
    Create(item T) error
    // GetAll() ([]T, error)
    GetByID(id uuid.UUID) (T, error)
    // Update(id uuid.UUID, item T) error
    // Delete(id uuid.UUID) error
}
