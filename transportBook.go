package main

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/go-kit/kit/endpoint"
    "github.com/gorilla/mux"
    "net/http"
)
// BOOK
func makeCreateBookEndpoint(s BookService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(CreateBookRequest)
        msg, err := s.CreateBook(ctx, req.book)
        return CreateBookResponse{Msg: msg, Err: err}, nil
    }
}
func makeGetBookByIdEndpoint(s BookService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(GetBookByIdRequest)
        bookDetails, err := s.GetBookById(ctx, req.Id)
        if err != nil {
            return GetBookByIdResponse{Book: bookDetails, Err: "Id not found"}, nil
        }
        return GetBookByIdResponse{Book: bookDetails, Err: ""}, nil
    }
}
// /books/{id}/authors
func makeGetBookAuthorsByIdEndpoint(s BookService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(GetBookAuthorsByIdRequest)
        bookDetails, err := s.GetBookAuthorsById(ctx, req.Id)
        if err != nil {
            return GetBookAuthorsByIdResponse{Author: bookDetails, Err: "Id not found"}, nil
        }
        return GetBookAuthorsByIdResponse{Author: bookDetails, Err: ""}, nil
    }
}
// /books/{id}/publishers
func makeGetBookPublishersByIdEndpoint(s BookService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(GetBookPublishersByIdRequest)
        bookDetails, err := s.GetBookPublishersById(ctx, req.Id)
        if err != nil {
            return GetBookPublishersByIdResponse{Publisher: bookDetails, Err: "Id not found"}, nil
        }
        return GetBookPublishersByIdResponse{Publisher: bookDetails, Err: ""}, nil
    }
}
func makeDeleteBookEndpoint(s BookService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(DeleteBookRequest)
        msg, err := s.DeleteBook(ctx, req.Bookid)
        if err != nil {
            return DeleteBookResponse{Msg: msg, Err: err}, nil
        }
        return DeleteBookResponse{Msg: msg, Err: nil}, nil
    }
}
func makeUpdateBookendpoint(s BookService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(UpdateBookRequest)
        msg, err := s.UpdateBook(ctx, req.book)
        return msg, err
    }
}

func decodeCreateBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req CreateBookRequest
    fmt.Println("-------->>>>into Decoding")
    if err := json.NewDecoder(r.Body).Decode(&req.book); err != nil {
        return nil, err
    }
    return req, nil
}

func decodeGetBookByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req GetBookByIdRequest
    fmt.Println("-------->>>>into GetById Decoding")
    vars := mux.Vars(r)
    req = GetBookByIdRequest{
        Id: vars["bookid"],
    }
    return req, nil
}
// /books/{id}/authors
func decodeGetBookAuthorsByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req GetBookAuthorsByIdRequest
    fmt.Println("-------->>>>into GetAuthorsById Decoding")
    vars := mux.Vars(r)
    req = GetBookAuthorsByIdRequest{
        Id: vars["bookid"],
    }
    return req, nil
}
// /books/{id}/publishers
func decodeGetBookPublishersByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req GetBookPublishersByIdRequest
    fmt.Println("-------->>>>into GetPublishersById Decoding")
    vars := mux.Vars(r)
    req = GetBookPublishersByIdRequest{
        Id: vars["bookid"],
    }
    return req, nil
}
func decodeDeleteBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Delete Decoding")
    var req DeleteBookRequest
    vars := mux.Vars(r)
    req = DeleteBookRequest{
        Bookid: vars["bookid"],
    }
    return req, nil
}
func decodeUpdateBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Update Decoding")
    var req UpdateBookRequest
    if err := json.NewDecoder(r.Body).Decode(&req.book); err != nil {
        return nil, err
    }
    return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    fmt.Println("into Encoding <<<<<<----------------")
    return json.NewEncoder(w).Encode(response)
}

type (
    CreateBookRequest struct {
        book Book
    }
    CreateBookResponse struct {
        Msg string `json:"msg"`
        Err error  `json:"error,omitempty"`
    }
    GetBookByIdRequest struct {
        Id string `json:"bookid"`
    }
    // /books/{id}/authors
    GetBookAuthorsByIdRequest struct {
        Id string `json:"bookid"`
    }
    // /books/{id}/publishers
    GetBookPublishersByIdRequest struct {
        Id string `json:"bookid"`
    }
    GetBookByIdResponse struct {
        Book interface{} `json:"book,omitempty"`
        Err  string      `json:"error,omitempty"`
    }
    // /books/{id}/authors
    GetBookAuthorsByIdResponse struct {
        Author interface{} `json:"author,omitempty"`
        Err  string      `json:"error,omitempty"`
    }
    // /books/{id}/publishers
    GetBookPublishersByIdResponse struct {
        Publisher interface{} `json:"publisher,omitempty"`
        Err  string      `json:"error,omitempty"`
    }

    DeleteBookRequest struct {
        Bookid string `json:"bookid"`
    }

    DeleteBookResponse struct {
        Msg string `json:"response"`
        Err error  `json:"error,omitempty"`
    }
    UpdateBookRequest struct {
        book Book
    }
    UpdateBookResponse struct {
        Msg string `json:"status,omitempty"`
        Err error  `json:"error,omitempty"`
    }
)