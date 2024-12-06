package repositories

import (
	"github.com/supermandavid/go_backend_boilerplate/entities"
)

type PostRepository interface {
	Save(post *entities.Post) (*entities.Post, error)
	FindAll() ([]entities.Post, error)
	Delete(post *entities.Post) error
}
