package test

import (
	"OttoAssignment/src/services"
	"testing"
)

func TestParseUserID(t *testing.T) {
	args := []string{"application", "1", ""}
	id, err := services.ParseUserID(args)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if id != 1 {
		t.Fatalf("expected integer 1, got %v", id)
	}
}

func TestParseUserIDNegative(t *testing.T) {
	args := []string{"application", "-1", ""}
	_, err := services.ParseUserID(args)
	if err == nil {
		t.Fatalf("expected error, got none")
	}
	args = []string{"app", "string", ""}
	_, err = services.ParseUserID(args)
	if err == nil {
		t.Fatalf("expected error, got none")
	}
}
