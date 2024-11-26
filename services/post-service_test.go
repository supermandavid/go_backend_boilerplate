package services

import (
	"awesomeBackend/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) FindAll() ([]entities.Post, error) {
	args := m.Called()
	result := args.Get(0)
	return result.([]entities.Post), args.Error(1)
}

func (m *mockRepository) Save(post *entities.Post) (*entities.Post, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*entities.Post), args.Error(1)
}

func TestCreate(t *testing.T) {

	mockRepo := new(mockRepository)
	post := entities.Post{Title: "A", Text: "B"}

	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.Create(&post)

	mockRepo.AssertExpectations(t)

	assert.NotNil(t, result.ID)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "B", result.Text)
	assert.Nil(t, err)

}

func TestFindAll(t *testing.T) {

	var identifier int64 = 1

	mockRepo := new(mockRepository)
	post := entities.Post{ID: identifier, Title: "A", Text: "B"}

	//Setup expectations
	mockRepo.On("FindAll").Return([]entities.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.FindAll()

	//Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)
	assert.Equal(t, identifier, result[0].ID)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "post is nil")

}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := entities.Post{ID: 1, Title: "", Text: "B"}

	testService := NewPostService(nil)

	err := testService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "post title is empty", err.Error())
}
