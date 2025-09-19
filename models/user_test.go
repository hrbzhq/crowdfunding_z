package models

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUserJSONOmitsPassword(t *testing.T) {
	u := User{
		Username: "alice",
		Email:    "alice@example.com",
		Password: "supersecret",
	}

	b, err := json.Marshal(u)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}

	s := string(b)
	if strings.Contains(s, "supersecret") || strings.Contains(s, "Password") {
		t.Fatalf("expected password to be omitted from JSON, got: %s", s)
	}
}
