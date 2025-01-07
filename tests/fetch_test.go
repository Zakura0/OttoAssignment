package test

import (
	"OttoAssignment/src/services"
	"testing"
)

func TestFetchPosts(t *testing.T) {
	userID := 1

	posts, err := services.FetchPosts(userID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(posts) == 0 {
		t.Fatalf("expected some posts, got none")
	}

	for _, post := range posts {
		if post.UserID != userID {
			t.Errorf("expected userID %d, got %d", userID, post.UserID)
		}
	}

	userID = -1
	posts, err = services.FetchPosts(userID)
	if err == nil {
		t.Fatalf("expected error, got none")
	}

	if len(posts) != 0 {
		t.Fatalf("expected 0 posts")
	}
}
