package main

import (
    "context"
    "github.com/go-kit/kit/log"
)

// AUTHOR
type Author struct {
    AuthorId    string `json:"authorId,omitempty"`
    Name     string `json:"name,omitempty"`
    Nationality   string `json:"nationality,omitempty"`
    Birth string `json:"birth,omitempty"`
    Genere  string `json:"genere,omitempty"`
}

type authorservice struct {
    logger log.Logger
}

// Service describes the Author service.
type AuthorService interface {
    CreateAuthor(ctx context.Context, author Author) (string, error)
    GetAuthorById(ctx context.Context, id string) (interface{}, error)
    // /author/{id}/books
    GetAuthorBooksById(ctx context.Context, id string) (interface{}, error)
    UpdateAuthor(ctx context.Context, author Author) (string, error)
    DeleteAuthor(ctx context.Context, id string) (string, error)
}

var authors = []Author{
    Author{AuthorId: "1", Name: "Author 1", Nationality: "CR",
        Birth: "1990-12-20", Genere: "Male"},
    Author{AuthorId: "2", Name: "Author 2", Nationality: "USA",
        Birth: "1990-12-20", Genere: "Female"},
}

func findAuthor(x string) int {
    for i, author := range authors {
        if x == author.AuthorId {
            return i
        }
    }
    return -1
}
// /author/{id}/books
func findBooks(x string) int {
    for _, author := range authors {
        if x == author.AuthorId {
            for i, book := range books {
                if book.AuthorId == author.AuthorId {
                    return i
                }
            }
        }
    }
    return -1
}

func NewAuthorService(logger log.Logger) AuthorService {
    return &authorservice{
        logger: logger,
    }
}

func (s authorservice) CreateAuthor(ctx context.Context, author Author) (string, error) {
    var msg = "success"
    authors = append(authors, author)
    return msg, nil
}

func (s authorservice) GetAuthorById(ctx context.Context, id string) (interface{}, error) {
    var err error
    var author interface{}
    var empty interface{}
    i := findAuthor(id)
    if i == -1 {
        return empty, err
    }
    author = authors[i]
    return author, nil
}
// /author/{id}/books
func (s authorservice) GetAuthorBooksById(ctx context.Context, id string) (interface{}, error) {
    var err error
    var book interface{}
    var empty interface{}
    i := findBooks(id)
    if i == -1 {
        return empty, err
    }
    book = books[i]
    return book, nil
}
func (s authorservice) DeleteAuthor(ctx context.Context, id string) (string, error) {
    var err error
    msg := ""
    i := findAuthor(id)
    if i == -1 {
        return "", err
    }
    copy(authors[i:], authors[i+1:])
    authors[len(authors)-1] = Author{}
    authors = authors[:len(authors)-1]
    return msg, nil
}
func (s authorservice) UpdateAuthor(ctx context.Context, author Author) (string, error) {
    var empty = ""
    var err error
    var msg = "success"
    i := findAuthor(author.AuthorId)
    if i == -1 {
        return empty, err
    }
    authors[i] = author
    return msg, nil
}