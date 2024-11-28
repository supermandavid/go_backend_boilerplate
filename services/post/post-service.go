package services

import (
	"awesomeBackend/entities"
	postRepo "awesomeBackend/repositories/post"
	"errors"
	"math/rand"
)

type PostService interface {
	Validate(post *entities.Post) error
	Create(post *entities.Post) (*entities.Post, error)
	FindAll() ([]entities.Post, error)
}

type service struct{}

var (
	repo postRepo.PostRepository
)

// NewPostService
func NewPostService(postRepository postRepo.PostRepository) PostService {
	repo = postRepository
	return &service{}
}

func (s service) Validate(post *entities.Post) error {
	if post == nil {
		return errors.New("post is nil")
	}
	if post.Title == "" {
		return errors.New("post title is empty")
	}
	return nil
}

func (s service) Create(post *entities.Post) (*entities.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}

func (s service) FindAll() ([]entities.Post, error) {
	return repo.FindAll()

}
