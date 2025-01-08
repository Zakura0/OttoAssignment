package services

import (
	"OttoAssignment/src/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// FetchPosts zieht sich alle Posts der JSONPlaceholder API um anschließend die Posts einer bestimmten Benutzer-ID rauszufiltern.
//
// Args:
//
//	userID (int): Die Benutzer-ID, dessen Posts abgerufen werden sollen.
//
// Return:
//
//	([]models.Post, error): Die angeforderten Posts oder ein Fehler, wenn der Abruf fehlerhaft war.
func FetchPosts(userID int) ([]models.Post, error) {
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

	fmt.Printf("Fetching posts for user ID: %d\n\n", userID)

	var userPosts []models.Post
	for _, post := range posts {
		comments, _ := fetchComments(post.ID)
		post.Comments = append(post.Comments, comments...)
		if post.UserID == userID {
			userPosts = append(userPosts, post)
		}
	}
	if len(userPosts) == 0 {
		return nil, fmt.Errorf("no posts found for user ID: %d", userID)
	}
	return userPosts, nil
}

// fetchComments zieht sich alle Kommentare der JSONPlaceholder API um anschließend die Kommentare einer bestimmten Post-ID rauszufiltern.
//
// Args:
//
//	postID (int): Die Post-ID, dessen Kommentare abgerufen werden sollen.
//
// Return:
//
//	([]models.Post, error): Die angeforderten Kommentare oder ein Fehler, wenn der Abruf fehlerhaft war.
func fetchComments(postID int) ([]models.Comment, error) {
	response, err := http.Get("https://jsonplaceholder.typicode.com/comments")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var comments []models.Comment
	err = json.Unmarshal(body, &comments)
	if err != nil {
		return nil, err
	}

	var postComments []models.Comment
	for _, comment := range comments {
		if comment.PostID == postID {
			postComments = append(postComments, comment)
		}
	}

	return postComments, nil
}
