package repositories

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/supermandavid/go_backend_boilerplate/entities"
	"google.golang.org/api/iterator"
	"log"
	"strconv"
)

type repo struct {
}

// NewFirestoreRepository
func NewFirestoreRepository() PostRepository {
	return &repo{}
}

const (
	projectID      = "test-a2054"
	collectionName = "posts"
)

func (*repo) Save(post *entities.Post) (*entities.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		log.Fatalf("Failed to create a firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()

	_, err = client.Collection(collectionName).Doc(strconv.Itoa(int(post.ID))).Set(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed to save post: %v", err)
	}

	return post, nil
}
func (*repo) Delete(post *entities.Post) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		log.Fatalf("Failed to create a firestore Client: %v", err)
		return err
	}
	defer client.Close()

	_, err = client.Collection(collectionName).Doc(strconv.Itoa(int(post.ID))).Delete(ctx)

	if err != nil {
		log.Fatalf("Failed to delete post: %v", err)
	}

	return nil
}

func (*repo) FindAll() ([]entities.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	defer client.Close()

	if err != nil {
		log.Fatalf("Failed to create a firestore Client: %v", err)
		return nil, err
	}

	var posts []entities.Post
	iter := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}
		post := entities.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}

	return posts, nil

}
