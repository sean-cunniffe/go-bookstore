package utils

import (
	"github.com/sean-cunniffe/go-bookstore/src/pkg/models"
	"strings"
	"testing"
)

func TestParseBody(t *testing.T) {
	name := "Book1"
	author := "JohnDoe"
	publication := "Penguin"
	jsonBody := "{\n\t\"name\":\"" + name + "\",\n\t\"author\":\"" + author + "\",\n\t\"publication\":\"" + publication + "\"\n}"
	reader := strings.NewReader(jsonBody)
	var actual models.Book
	ParseBody(reader, &actual)
	if name != actual.Name {
		t.Fatalf("actual.name = %v, wanted %v", actual.Name, name)
	}
	if author != actual.Author {
		t.Fatalf("actual.author = %v, wanted %v", actual.Author, author)
	}
	if publication != actual.Publication {
		t.Fatalf("actual.publication = %v, wanted %v", actual.Publication, publication)
	}
}

func TestParseIllegalBody(t *testing.T) {
	jsonBody := "not a body"
	reader := strings.NewReader(jsonBody)
	var book models.Book
	err := ParseBody(reader, &book)
	if err == nil {
		t.Fatal("Expected an error to return from ParseBody(reader, &book)")
	}
}
