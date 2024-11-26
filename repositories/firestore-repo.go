package repositories

import (
	"awesomeBackend/entities"
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
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

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed to save post: %v", err)
	}

	return post, nil
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
		fmt.Println(doc)

		if err == iterator.Done {
			break
		}

		if err != nil {
			fmt.Println(len(posts))
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
