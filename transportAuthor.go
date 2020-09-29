package main

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/go-kit/kit/endpoint"
    "github.com/gorilla/mux"
    "net/http"
)
// AUTHOR
func makeCreateAuthorEndpoint(s AuthorService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(CreateAuthorRequest)
        msg, err := s.CreateAuthor(ctx, req.author)
        return CreateAuthorResponse{Msg: msg, Err: err}, nil
    }
}
func makeGetAuthorByIdEndpoint(s AuthorService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(GetAuthorByIdRequest)
        authorDetails, err := s.GetAuthorById(ctx, req.Id)
        if err != nil {
            return GetAuthorByIdResponse{Author: authorDetails, Err: "Id not found"}, nil
        }
        return GetAuthorByIdResponse{Author: authorDetails, Err: ""}, nil
    }
}
// /author/{id}/books
func makeGetAuthorBooksByIdEndpoint(s AuthorService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(GetAuthorBooksByIdRequest)
        authorDetails, err := s.GetAuthorBooksById(ctx, req.Id)
        if err != nil {
            return GetAuthorBooksByIdResponse{Book: authorDetails, Err: "Id not found"}, nil
        }
        return GetAuthorBooksByIdResponse{Book: authorDetails, Err: ""}, nil
    }
}
func makeDeleteAuthorEndpoint(s AuthorService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(DeleteAuthorRequest)
        msg, err := s.DeleteAuthor(ctx, req.Authorid)
        if err != nil {
            return DeleteAuthorResponse{Msg: msg, Err: err}, nil
        }
        return DeleteAuthorResponse{Msg: msg, Err: nil}, nil
    }
}
func makeUpdateAuthorendpoint(s AuthorService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(UpdateAuthorRequest)
        msg, err := s.UpdateAuthor(ctx, req.author)
        return msg, err
    }
}

func decodeCreateAuthorRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req CreateAuthorRequest
    fmt.Println("-------->>>>into Decoding")
    if err := json.NewDecoder(r.Body).Decode(&req.author); err != nil {
        return nil, err
    }
    return req, nil
}

func decodeGetAuthorByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req GetAuthorByIdRequest
    fmt.Println("-------->>>>into GetById Decoding")
    vars := mux.Vars(r)
    req = GetAuthorByIdRequest{
        Id: vars["authorid"],
    }
    return req, nil
}
// /author/{id}/books
func decodeGetAuthorBooksByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req GetAuthorBooksByIdRequest
    fmt.Println("-------->>>>into GetById Decoding")
    vars := mux.Vars(r)
    req = GetAuthorBooksByIdRequest{
        Id: vars["authorid"],
    }
    return req, nil
}
func decodeDeleteAuthorRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Delete Decoding")
    var req DeleteAuthorRequest
    vars := mux.Vars(r)
    req = DeleteAuthorRequest{
        Authorid: vars["authorid"],
    }
    return req, nil
}
func decodeUpdateAuthorRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Update Decoding")
    var req UpdateAuthorRequest
    if err := json.NewDecoder(r.Body).Decode(&req.author); err != nil {
        return nil, err
    }
    return req, nil
}

type (
    CreateAuthorRequest struct {
        author Author
    }
    CreateAuthorResponse struct {
        Msg string `json:"msg"`
        Err error  `json:"error,omitempty"`
    }
    GetAuthorByIdRequest struct {
        Id string `json:"authorid"`
    }
    // /author/{id}/books
    GetAuthorBooksByIdRequest struct {
        Id string `json:"authorid"`
    }
    GetAuthorByIdResponse struct {
        Author interface{} `json:"author,omitempty"`
        Err  string      `json:"error,omitempty"`
    }
    // /author/{id}/books
    GetAuthorBooksByIdResponse struct {
        Book interface{} `json:"book,omitempty"`
        Err  string      `json:"error,omitempty"`
    }
    DeleteAuthorRequest struct {
        Authorid string `json:"authorid"`
    }

    DeleteAuthorResponse struct {
        Msg string `json:"response"`
        Err error  `json:"error,omitempty"`
    }
    UpdateAuthorRequest struct {
        author Author
    }
    UpdateAuthorResponse struct {
        Msg string `json:"status,omitempty"`
        Err error  `json:"error,omitempty"`
    }
)