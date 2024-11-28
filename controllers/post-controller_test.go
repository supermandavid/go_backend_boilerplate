package controllers

import (
	"awesomeBackend/entities"
	router "awesomeBackend/http"
	repositories "awesomeBackend/repositories/post"
	postSrv "awesomeBackend/services/post"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	IDENTIFIER int64  = 111
	TITLE      string = "Test Title 12"
	TEXT       string = "Test Text 12"
)

var (
	testPostRepository repositories.PostRepository = repositories.NewSQLiteRepository()
	testPostService    postSrv.PostService         = postSrv.NewPostService(testPostRepository)
	testPostController PostController              = NewPostController(testPostService)
	httpRouter                                     = router.NewGinRouter()
)

func TestAddPost(t *testing.T) {
	//Create a new HTTP Post request
	var input = []byte(`{
		"title": "` + TITLE + `",
		"text": "` + TEXT + `"
	}`)

	rPath := "/posts"
	req, _ := http.NewRequest("POST", rPath, bytes.NewBuffer(input))

	//Assign HTTP Handler function
	httpRouter.POST(rPath, testPostController.AddPost)

	//Record HTTP Response
	response := httptest.NewRecorder()

	//Dispatch the HTTP request
	httpRouter.ServeRequest(response, req)

	//Log
	t.Logf("status: %d", response.Code)
	t.Logf("response: %s", response.Body.String())

	//Add Assertions on the HTTP status code and the response
	status := response.Code
	if status != http.StatusCreated {
		t.Errorf("Response should be %d but got %d", http.StatusCreated, status)
	}

	//Decode the HTTP response
	var post entities.Post

	var result map[string]interface{}

	err := json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		t.Error("Could not decode response body")
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		t.Error("Could not find \"data\" in response body")
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Error("Could not encode response body")
	}

	err = json.Unmarshal(jsonData, &post)
	if err != nil {
		t.Error("Could not decode to post entity")
	}

	// Assert HTTP RESPONSE

	assert.NotNil(t, post.ID)
	assert.Equal(t, TITLE, post.Title)
	assert.Equal(t, TEXT, post.Text)

	//clean up database
	cleanUp(&post)
}

func TestGetPost(t *testing.T) {

	//Create a new HTTP Post request
	setup()

	rPath := "/posts"
	req, _ := http.NewRequest("GET", rPath, nil)

	//Assign HTTP Handler function
	httpRouter.GET(rPath, testPostController.GetPosts)

	//Record HTTP Response
	response := httptest.NewRecorder()

	//Dispatch the HTTP request
	httpRouter.ServeRequest(response, req)

	//Log
	t.Logf("status: %d", response.Code)
	t.Logf("response: %s", response.Body.String())

	//Add Assertions on the HTTP status code and the response
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Response should be %d but got %d", http.StatusCreated, status)
	}

	//Decode the HTTP response
	var posts []entities.Post

	var result map[string]interface{}

	err := json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		t.Error("Could not decode response body")
	}

	data, ok := result["data"].([]interface{})
	if !ok {
		t.Error("Could not find \"data\" in response body")
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Error("Could not encode resp onse body")
	}

	err = json.Unmarshal(jsonData, &posts)
	if err != nil {
		t.Error("Could not decode to post entity")
	}

	// Assert HTTP RESPONSE

	assert.NotNil(t, posts[0].ID)
	assert.Equal(t, TITLE, posts[0].Title)
	assert.Equal(t, TEXT, posts[0].Text)

	//clean up database
	cleanUp(&posts[0])

}

func setup() {
	var post = entities.Post{
		Title: TITLE,
		Text:  TEXT,
		ID:    IDENTIFIER,
	}
	_, err := testPostRepository.Save(&post)

	if err != nil {
		fmt.Println("error with saving post")
		panic(err)
	} else {
		fmt.Println("saved post successfully")
	}
}

func cleanUp(post *entities.Post) {
	err := testPostRepository.Delete(post)
	if err != nil {
		panic(err)
	}
}
