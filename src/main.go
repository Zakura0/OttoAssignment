package main

import (
	"OttoAssignment/src/services"
	"fmt"
	"os"
)

func main() {
	userID, err := services.ParseUserID(os.Args)
	if len(os.Args) > 3 {
		fmt.Printf("Invalid arguments\nUsage: go run main.go <userID> <filter>")
		return
	}
	filter := ""
	if len(os.Args) == 3 {
		filter = os.Args[2]
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	posts, err := services.FetchPosts(userID)
	if err != nil {
		fmt.Printf("Error fetching posts: %v\n", err)
		return
	}
	services.PrintPosts(userID, posts, filter)
}
