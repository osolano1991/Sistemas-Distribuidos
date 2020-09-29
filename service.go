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
    Book{BookId: "1", PublisherId: "1", Title: "Libro 1",
		Copyright: "2012", Edition: "5th", Pages: "976"},
	Book{BookId: "2", PublisherId: "1", Title: "Libro 2",
		Copyright: "2010", Edition: "9th", Pages: "1500"},
	Book{BookId: "3", PublisherId: "1", Title: "Libro 3",
		Copyright: "2010", Edition: "9th", Pages: "1500"},
	Book{BookId: "4", PublisherId: "2", Title: "Libro 4",
		Copyright: "2010", Edition: "11th", Pages: "722"},
}

func findBook(x string) int {
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
    i := findBook(id)
    if i == -1 {
        return empty, err
    }
    book = books[i]
    return book, nil
}
func (s bookservice) DeleteBook(ctx context.Context, id string) (string, error) {
    var err error
    msg := ""
    i := findBook(id)
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
    i := 0findBook(book.BookId)
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
    Author{AuthorId: "1", BookId: "3", Name: "OSCAR", Nationality: "Costa Rica",
		Birth: "1991", Genere: "First"},
	Author{AuthorId: "2", BookId: "2", Name: "MARIO", Nationality: "Costa Rica",
		Birth: "1991", Genere: "Second"},
	Author{AuthorId: "3", BookId: "2", Name: "LUIS", Nationality: "Panama",
		Birth: "1994", Genere: "Third"},
	Author{AuthorId: "4", BookId: "1", Name: "JOUSER", Nationality: "FRANCES",
		Birth: "1980", Genere: "Fourth"},
	Author{AuthorId: "5", BookId: "4", Name: "Ronald", Nationality: "Ingles",
		Birth: "1967", Genere: "Fiveth"},
}

func findAuthor(x string) int {
    for i, author := range authors {
        if x == author.AuthorId {
            return i
        }
    }
    return -1
}

func NewServiceAuthor(logger log.Logger) AuthorService {
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

//==========================================================================================================================================================
//                                                                          PUBLISHERS
//==========================================================================================================================================================
type Publisher struct {
    PublisherId string `json:"publisherId,omitempty"`
	Name string `json:"name,omitempty"`
	Country string `json:"country,omitempty"`
	Founded string `json:"founded,omitempty"`
	Genere string `json:"genere,omitempty"`
}

type publisherservice struct {
    logger log.Logger
}

// Service describes the Publisher service.
type PublisherService interface {
    CreatePublisher(ctx context.Context, publisher Publisher) (string, error)
    GetPublisherById(ctx context.Context, id string) (interface{}, error)
    UpdatePublisher(ctx context.Context, publisher Publisher) (string, error)
    DeletePublisher(ctx context.Context, id string) (string, error)
}

var publishers = []Publisher{
    Publisher{PublisherId: "1", Name: "Yensie", Country: "Inglaterra", Founded: "Costa Rica",
		Genere: "First"},
	Publisher{PublisherId: "2", Name: "Tatiana", Country: "Italia", Founded: "Costa Rica",
		Genere: "Second"},
	Publisher{PublisherId: "3", Name: "ANGIE", Country: "CHINA", Founded: "CHINA",
		Genere: "Third"},
	Publisher{PublisherId: "4", Name: "NADIA", Country: "JAPON", Founded: "JAPON",
		Genere: "Sixth"},
	Publisher{PublisherId: "5", Name: "FLORY", Country: "NICARAGUA", Founded: "NICARAGUA",
		Genere: "Seventh"},
}

func findPublisher(x string) int {
    for i, publisher := range publishers {
        if x == publisher.PublisherId {
            return i
        }
    }
    return -1
}

func NewServicePublisher(logger log.Logger) PublisherService {
    return &publisherservice{
        logger: logger,
    }
}

func (s publisherservice) CreatePublisher(ctx context.Context, publisher Publisher) (string, error) {
    var msg = "success"
    publishers = append(publishers, publisher)
    return msg, nil
}

func (s publisherservice) GetPublisherById(ctx context.Context, id string) (interface{}, error) {
    var err error
    var publisher interface{}
    var empty interface{}
    i := findPublisher(id)
    if i == -1 {
        return empty, err
    }
    publisher = publishers[i]
    return publisher, nil
}
func (s publisherservice) DeletePublisher(ctx context.Context, id string) (string, error) {
    var err error
    msg := ""
    i := findPublisher(id)
    if i == -1 {
        return "", err
    }
    copy(publishers[i:], publishers[i+1:])
    publishers[len(publishers)-1] = Publisher{}
    publishers = publishers[:len(publishers)-1]
    return msg, nil
}
func (s publisherservice) UpdatePublisher(ctx context.Context, publisher Publisher) (string, error) {
    var empty = ""
    var err error
    var msg = "success"
    i := findPublisher(publisher.PublisherId)
    if i == -1 {
        return empty, err
    }
    publishers[i] = publisher
    return msg, nil
}