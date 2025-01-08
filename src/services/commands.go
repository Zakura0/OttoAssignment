package services

import (
	"OttoAssignment/src/models"
	"fmt"
	"strconv"
	"strings"
)

// ParseUserID analysiert die Benutzer-ID aus den Eingaben des Nutzers.
// Es erwartet mindestens zwei Argumente, wobei das zweite Argument die Benutzer-ID ist.
// Gibt die analysierte Benutzer-ID als Ganzzahl zurück oder einen Fehler, wenn die Benutzer-ID ungültig ist.
//
// Args:
//
//	args ([]string): Die Befehlszeilenargumente.
//
// Return:
//
//	(int, error): Die analysierte Benutzer-ID oder ein Fehler, wenn die Benutzer-ID ungültig ist.
func ParseUserID(args []string) (int, error) {
	if len(args) < 2 {
		return 0, fmt.Errorf("usage: main.exe <userID> <filter>")
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

// PrintPosts printed die Posts und deren Kommentare für eine gegebene Benutzer-ID zum STDOUT.
// Es wechselt die Farbe der Kommentare für eine bessere Lesbarkeit.
// Wenn ein Filter angegeben wird, werden nur Kommentare gedruckt, die den Filter enthalten.
//
// Args:
//
//	userID (int): Die Benutzer-ID.
//	posts ([]models.Post): Die Liste der zu druckenden Beiträge.
//	filter (string): Der anzuwendende Filter.
func PrintPosts(userID int, posts []models.Post, filter string) {
	var Color = "\033[32m"
	var Reset = "\033[0m"
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
				} else {
					continue
				}
			} else {
				fmt.Printf(Color+"Author: %s \nTitle: %s \n\n%s \n\n", comment.Email, comment.Name, comment.Body+Reset)

			}
			switchColor = !switchColor
		}
	}
}
