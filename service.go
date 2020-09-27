package main

import (
    "context"
    "github.com/go-kit/kit/log"
)

type Book struct {
    BookId    string `json:"bookId,omitempty"`
    Title     string `json:"title,omitempty"`
    Edition   string `json:"edition,omitempty"`
    Copyright string `json:"copyright,omitempty"`
    Language  string `json:"language,omitempty"`
    Pages     string `json:"pages,omitempty"`
    Author    string `json:"author,omitempty"`
    Publisher string `json:"publisher,omitempty"`
}

type bookservice struct {
    logger log.Logger
}

// Service describes the Book service.
type BookService interface {
    CreateBook(ctx context.Context, book Book) (string, error)
    GetBookById(ctx context.Context, id string) (interface{}, error)
    UpdateBook(ctx context.Context, book Book) (string, error)
    DeleteBook(ctx context.Context, id string) (string, error)
}

var books = []Book{
    Book{BookId: "Book1", Title: "Operating System Concepts", Edition: "9th",
        Copyright: "2012", Language: "ENGLISH", Pages: "976",
        Author: "Abraham Silberschatz", Publisher: "John Wiley & Sons"},
    Book{BookId: "Book3", Title: "Computer Networks", Edition: "5th",
        Copyright: "2010", Language: "ENGLISH", Pages: "960",
        Author: "Andrew S. Tanenbaum", Publisher: "Andrew S. Tanenbaum"},
}

func find(x string) int {
    for i, book := range books {
        if x == book.BookId {
            return i
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
//==========================================================================================================================================================
//                                                                          AUTHORS
//==========================================================================================================================================================
type Author struct {
    AuthorId string `json:"authorId,omitempty"`
	BookId string `json:"bookId,omitempty"`
	Name string `json:"name,omitempty"`
	Nationality string `json:"nationality,omitempty"`
	Birth string `json:"birth,omitempty"`
	Genere string `json:"genere,omitempty"`
}

type authorservice struct {
    logger log.Logger
}

// Service describes the Author service.
type AuthorService interface {
    CreateAuthor(ctx context.Context, author Author) (string, error)
    GetAuthorById(ctx context.Context, id string) (interface{}, error)
    UpdateAuthor(ctx context.Context, author Author) (string, error)
    DeleteAuthor(ctx context.Context, id string) (string, error)
}

var authors = []Author{
    Author{AuthorId: "1", BookId: "1", Name: "Author 1",
        Nationality: "Costa Rica", Birth: "julio 1991", Genere: "Masculino"},
    Author{AuthorId: "2", BookId: "2", Name: "Author 2",
        Nationality: "Espanol", Birth: "agosto 2000", Genere: "Masculino"},
}

func find(x string) int {
    for i, author := range authors {
        if x == author.AuthorId {
            return i
        }
    }
    return -1
}

func NewService(logger log.Logger) AuthorService {
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
    i := find(id)
    if i == -1 {
        return empty, err
    }
    author = authors[i]
    return author, nil
}
func (s authorservice) DeleteAuthor(ctx context.Context, id string) (string, error) {
    var err error
    msg := ""
    i := find(id)
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
    i := find(author.AuthorId)
    if i == -1 {
        return empty, err
    }
    authors[i] = author
    return msg, nil
}