package repositories

import (
	"awesomeBackend/entities"
)

type PostRepository interface {
	Save(post *entities.Post) (*entities.Post, error)
	FindAll() ([]entities.Post, error)
	Delete(post *entities.Post) error
}
