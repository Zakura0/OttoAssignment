package services

import (
	"OttoAssignment/src/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// FetchPosts ruft die Beiträge eines bestimmten Nutzers von der API ab
func FetchPosts(userID int) ([]models.Post, error) {
	//Anfrage an die API
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return nil, fmt.Errorf("error fetching posts: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	var posts []models.Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling of JSON-Files: %v", err)
	}

	var userPosts []models.Post
	for _, post := range posts {
		if post.UserID == userID {
			userPosts = append(userPosts, post)
		}
	}
	return userPosts, nil
}
