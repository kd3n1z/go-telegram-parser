package telegramparser

import (
	"testing"
)

const token = "123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11"

func TestParser_Parse_Valid(t *testing.T) {
	parser := CreateParser(token)

	data, err := parser.Parse("user=%7B%22first_name%22%3A%22test%22%7D&auth_date=1&hash=c8a1a1fcd1bc0ffd66f4202a82fc92afa3078d97668eb392dd4ea7df2655e6c9")

	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}

	if data.User.FirstName != "test" {
		t.Errorf("Expected 'test' but got %v", data.User.FirstName)
	}
}

func TestParser_Parse_InvalidHash(t *testing.T) {
	parser := CreateParser(token)

	_, err := parser.Parse("user=%7B%22first_name%22%3A%22test%22%7D&auth_date=1&hash=c8a1a1fcd1bc0ffd66f4202a82fc92afa3078d97668eb392dd4ea7df2655e6c")

	if err == nil {
		t.Errorf("Expected an error but got nil")
	}
}
