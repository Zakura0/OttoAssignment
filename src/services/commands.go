package services

import (
	"OttoAssignment/src/models"
	"fmt"
	"strconv"
	"strings"
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

var Color = "\033[32m"
var Reset = "\033[0m"

func PrintPosts(userID int, posts []models.Post, filter string) {
	switchColor := true
	for _, post := range posts {
		fmt.Printf("Post ID: %d\nTitle: %s\n\n", post.ID, post.Title)
		fmt.Printf("%s\n\n", post.Body)
		for _, comment := range post.Comments {
			if switchColor {
				Color = "\033[93m"
			} else {
				Color = "\033[33m"
			}
			if filter != "" {
				if strings.Contains(comment.Body, filter) {
					fmt.Printf(Color+"Author: %s \nTitle: %s \n\n%s \n\n", comment.Email, comment.Name, comment.Body+Reset)
				}
			} else {
				fmt.Printf(Color+"Author: %s \nTitle: %s \n\n%s \n\n", comment.Email, comment.Name, comment.Body+Reset)

			}
			switchColor = !switchColor
		}
	}
}
