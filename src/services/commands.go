package services

import (
	"OttoAssignment/src/models"
	"fmt"
	"strconv"
)

func ParseUserID(args []string) (int, error) {
	if len(args) < 2 {
		return 0, fmt.Errorf("usage: go run main.go <userID>")
	}

	userID64, err := strconv.ParseInt(args[1], 10, 0)
	if err != nil {
		return 0, fmt.Errorf("invalid UserID: %v\nmake sure to only use numbers", args[1])
	}

	userID := int(userID64)
	if userID < 0 {
		return 0, fmt.Errorf("invalid UserID: %v\nonly positive numbers allowed", args[1])
	}

	return userID, nil
}

func PrintPosts(userID int, posts []models.Post) {
	fmt.Printf("Fetching posts for user ID: %d\n\n", userID)
	for _, post := range posts {
		fmt.Printf("Post ID: %d\nTitle: %s\n", post.ID, post.Title)
		fmt.Printf("Body: %s\n\n", post.Body)
	}
}
