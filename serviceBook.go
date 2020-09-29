package main

import (
    "context"
    "github.com/go-kit/kit/log"
)

// BOOK
type Book struct {
    BookId    string `json:"bookId,omitempty"`
    Title     string `json:"title,omitempty"`
    Edition   string `json:"edition,omitempty"`
    Copyright string `json:"copyright,omitempty"`
    Language  string `json:"language,omitempty"`
    Pages     string `json:"pages,omitempty"`
    AuthorId    string `json:"author,omitempty"`
    PublisherId string `json:"publisher,omitempty"`
}

type bookservice struct {
    logger log.Logger
}

// Service describes the Book service.
type BookService interface {
    CreateBook(ctx context.Context, book Book) (string, error)
    GetBookById(ctx context.Context, id string) (interface{}, error)
    GetBookAuthorsById(ctx context.Context, id string) (interface{}, error)
    GetBookPublishersById(ctx context.Context, id string) (interface{}, error)
    UpdateBook(ctx context.Context, book Book) (string, error)
    DeleteBook(ctx context.Context, id string) (string, error)
}

var books = []Book{
    Book{BookId: "1", Title: "Operating System Concepts", Edition: "9th",
        Copyright: "2012", Language: "ENGLISH", Pages: "976",
        AuthorId: "1", PublisherId: "1"},
    Book{BookId: "2", Title: "Computer Networks", Edition: "5th",
        Copyright: "2010", Language: "ENGLISH", Pages: "960",
        AuthorId: "2", PublisherId: "2"},
    Book{BookId: "3", Title: "OSCAR", Edition: "5th",
        Copyright: "2010", Language: "ENGLISH", Pages: "960",
        AuthorId: "2", PublisherId: "2"},
}

func find(x string) int {
    for i, book := range books {
        if x == book.BookId {
            return i
        }
    }
    return -1
}
// /books/{id}/authors
func findAuthors(x string) int {
    for _, book := range books {
        if x == book.BookId {
            for i, author := range authors {
                if book.AuthorId == author.AuthorId {
                    return i
                }
            }
        }
    }
    return -1
}
// /books/{id}/publishers
func findPublishers(x string) int {
    for _, book := range books {
        if x == book.BookId {
            for i, publisher := range publishers {
                if book.PublisherId == publisher.PublisherId {
                    return i
                }
            }
        }
    }
    return -1
}

func NewService(logger log.Logger) BookService {
    return &bookservice{
        logger: logger,
    }
}

func (s bookservice) CreateBook(ctx context.Context, book Book) (string, error) {
    var msg = "success"
    books = append(books, book)
    return msg, nil
}

func (s bookservice) GetBookById(ctx context.Context, id string) (interface{}, error) {
    var err error
    var book interface{}
    var empty interface{}
    i := find(id)
    if i == -1 {
        return empty, err
    }
    book = books[i]
    return book, nil
}
// /books/{id}/authors
func (s bookservice) GetBookAuthorsById(ctx context.Context, id string) (interface{}, error) {
    var err error
    var author interface{}
    var empty interface{}
    i := findAuthors(id)
    if i == -1 {
        return empty, err
    }
    author = authors[i]
    return author, nil
}
// /books/{id}/publishers
func (s bookservice) GetBookPublishersById(ctx context.Context, id string) (interface{}, error) {
    var err error
    var publisher interface{}
    var empty interface{}
    i := findAuthors(id)
    if i == -1 {
        return empty, err
    }
    publisher = publishers[i]
    return publisher, nil
}
func (s bookservice) DeleteBook(ctx context.Context, id string) (string, error) {
    var err error
    msg := ""
    i := find(id)
    if i == -1 {
        return "", err
    }
    copy(books[i:], books[i+1:])
    books[len(books)-1] = Book{}
    books = books[:len(books)-1]
    return msg, nil
}
func (s bookservice) UpdateBook(ctx context.Context, book Book) (string, error) {
    var empty = ""
    var err error
    var msg = "success"
    i := find(book.BookId)
    if i == -1 {
        return empty, err
    }
    books[i] = book
    return msg, nil
}