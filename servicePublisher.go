package main

import (
    "context"
    "github.com/go-kit/kit/log"
)

// PUBLISHER
type Publisher struct {
    PublisherId    string `json:"publisherId,omitempty"`
    Name     string `json:"name,omitempty"`
    Country   string `json:"country,omitempty"`
    Founded string `json:"founded,omitempty"`
    Genere  string `json:"genere,omitempty"`
}

type publisherservice struct {
    logger log.Logger
}

// Service describes the Publisher service.
type PublisherService interface {
    CreatePublisher(ctx context.Context, publisher Publisher) (string, error)
    GetPublisherById(ctx context.Context, id string) (interface{}, error)
    // /publisher/{id}/books
    GetPublisherBooksById(ctx context.Context, id string) (interface{}, error)
    UpdatePublisher(ctx context.Context, publisher Publisher) (string, error)
    DeletePublisher(ctx context.Context, id string) (string, error)
}

var publishers = []Publisher{
    Publisher{PublisherId: "1", Name: "Publisher 1", Country: "CR",
        Founded: "1990-12-20", Genere: "Male"},
    Publisher{PublisherId: "2", Name: "Publisher 2", Country: "USA",
        Founded: "1990-12-20", Genere: "Female"},
}

func findPublisher(x string) int {
    for i, publisher := range publishers {
        if x == publisher.PublisherId {
            return i
        }
    }
    return -1
}
// /publisher/{id}/books
func findPublisherBooks(x string) int {
    for _, publisher := range publishers {
        if x == publisher.PublisherId {
            for i, book := range books {
                if book.PublisherId == publisher.PublisherId {
                    return i
                }
            }
        }
    }
    return -1
}

func NewPublisherService(logger log.Logger) PublisherService {
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
// /publisher/{id}/books
func (s publisherservice) GetPublisherBooksById(ctx context.Context, id string) (interface{}, error) {
    var err error
    var book interface{}
    var empty interface{}
    i := findPublisherBooks(id)
    if i == -1 {
        return empty, err
    }
    book = books[i]
    return book, nil
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
