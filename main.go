package main

import (
    "github.com/go-kit/kit/log"
    httptransport "github.com/go-kit/kit/transport/http"
    "github.com/gorilla/mux"
    "net/http"
    "os"
)

func main() {
    logger := log.NewLogfmtLogger(os.Stderr)

    r := mux.NewRouter()

    var svc BookService
    svc = NewService(logger)

    // svc = loggingMiddleware{logger, svc}
    // svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}

    CreateBookHandler := httptransport.NewServer(
        makeCreateBookEndpoint(svc),
        decodeCreateBookRequest,
        encodeResponse,
    )
    GetByBookIdHandler := httptransport.NewServer(
        makeGetBookByIdEndpoint(svc),
        decodeGetBookByIdRequest,
        encodeResponse,
    )
    DeleteBookHandler := httptransport.NewServer(
        makeDeleteBookEndpoint(svc),
        decodeDeleteBookRequest,
        encodeResponse,
    )
    UpdateBookHandler := httptransport.NewServer(
        makeUpdateBookendpoint(svc),
        decodeUpdateBookRequest,
        encodeResponse,
    )
    
//=============================================================================================================
//                                                     Author
//=============================================================================================================
var svcAuthor AuthorService
    svcAuthor = NewServiceAuthor(logger)

    // svcAuthor = loggingMiddleware{logger, svcAuthor}
    // svcAuthor = instrumentingMiddleware{requestCount, requestLatency, countResult, svcAuthor}

    CreateAuthorHandler := httptransport.NewServer(
        makeCreateAuthorEndpoint(svcAuthor),
        decodeCreateAuthorRequest,
        encodeResponseAuthor,
    )
    GetByAuthorIdHandler := httptransport.NewServer(
        makeGetAuthorByIdEndpoint(svcAuthor),
        decodeGetAuthorByIdRequest,
        encodeResponseAuthor,
    )
    DeleteAuthorHandler := httptransport.NewServer(
        makeDeleteAuthorEndpoint(svcAuthor),
        decodeDeleteAuthorRequest,
        encodeResponseAuthor,
    )
    UpdateAuthorHandler := httptransport.NewServer(
        makeUpdateAuthorendpoint(svcAuthor),
        decodeUpdateAuthorRequest,
        encodeResponseAuthor,
    )
   
    
//=============================================================================================================
//                                                     PUBLISHER
//=============================================================================================================
var svcPublisher PublisherService
    svcPublisher = NewServicePublisher(logger)

    // svcPublisher = loggingMiddleware{logger, svcPublisher}
    // svcPublisher = instrumentingMiddleware{requestCount, requestLatency, countResult, svcPublisher}

    CreatePublisherHandler := httptransport.NewServer(
        makeCreatePublisherEndpoint(svcPublisher),
        decodeCreatePublisherRequest,
        encodeResponsePublisher,
    )
    GetBypublisheridHandler := httptransport.NewServer(
        makeGetPublisherByIdEndpoint(svcPublisher),
        decodeGetPublisherByIdRequest,
        encodeResponsePublisher,
    )
    DeletePublisherHandler := httptransport.NewServer(
        makeDeletePublisherEndpoint(svcPublisher),
        decodeDeletePublisherRequest,
        encodeResponsePublisher,
    )
    UpdatePublisherHandler := httptransport.NewServer(
        makeUpdatePublisherendpoint(svcPublisher),
        decodeUpdatePublisherRequest,
        encodeResponsePublisher,
    )
    
    http.Handle("/", r)
    http.Handle("/books", CreateBookHandler)
    http.Handle("/books/update", UpdateBookHandler)
    r.Handle("/books/{bookid}", GetByBookIdHandler).Methods("GET")
    r.Handle("/books/{bookid}/authors", GetByBookIdHandler).Methods("GET")
    r.Handle("/books/{bookid}/publishers", GetByBookIdHandler).Methods("GET")
    r.Handle("/books/{bookid}", DeleteBookHandler).Methods("DELETE")
    
    http.Handle("/author", CreateAuthorHandler)
    http.Handle("/author/update", UpdateAuthorHandler)
    r.Handle("/author/{authorid}", GetByAuthorIdHandler).Methods("GET")
    r.Handle("/author/{authorid}/books", GetByAuthorIdHandler).Methods("GET")
    r.Handle("/author/{authorid}", DeleteAuthorHandler).Methods("DELETE")    
    
    http.Handle("/publisher", CreatePublisherHandler)
    http.Handle("/publisher/update", UpdatePublisherHandler)
    r.Handle("/publisher/{publisherid}", GetBypublisheridHandler).Methods("GET")
    r.Handle("/publisher/{publisherid}/books", GetBypublisheridHandler).Methods("GET")
    r.Handle("/publisher/{publisherid}", DeletePublisherHandler).Methods("DELETE")

    // http.Handle("/metrics", promhttp.Handler())
    logger.Log("msg", "HTTP", "addr", ":"+os.Getenv("PORT"))
    logger.Log("err", http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}