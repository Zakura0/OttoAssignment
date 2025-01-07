package main

import (
	"OttoAssignment/src/services"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <userID>")
		return
	}
	inputID := os.Args[1]
	userID64, err := strconv.ParseInt(inputID, 10, 0)
	if err != nil {
		fmt.Printf("Invalid UserID: %v\nmake sure to only use numbers", inputID)
		return
	}
	userID := int(userID64)
	if userID < 0 {
		fmt.Printf("Invalid UserID: %v\n only positive numbers allowed", inputID)
		return
	}
	fmt.Printf("Fetching posts for user ID: %d\n", userID)

	posts, err := services.FetchPosts(userID)
	if err != nil {
		fmt.Printf("error fetching posts: %v\n", err)
		return
	}

	for _, post := range posts {
		fmt.Printf("Post ID: %d\nTitle: %s\n", post.ID, post.Title)
		fmt.Printf("Body: %s\n\n", post.Body)
	}
}
