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

    var bookSvc BookService
    bookSvc = NewService(logger)

    var authorSvc AuthorService
    authorSvc = NewAuthorService(logger)

    var publisherSvc PublisherService
    publisherSvc = NewPublisherService(logger)



    // BOOK
    CreateBookHandler := httptransport.NewServer(
        makeCreateBookEndpoint(bookSvc),
        decodeCreateBookRequest,
        encodeResponse,
    )
    GetByBookIdHandler := httptransport.NewServer(
        makeGetBookByIdEndpoint(bookSvc),
        decodeGetBookByIdRequest,
        encodeResponse,
    )
    // /books/{id}/authors
    GetAuthorsByBookIdHandler := httptransport.NewServer(
        makeGetBookAuthorsByIdEndpoint(bookSvc),
        decodeGetBookAuthorsByIdRequest,
        encodeResponse,
    )
    // /books/{id}/publishers
    GetPublishersByBookIdHandler := httptransport.NewServer(
        makeGetBookPublishersByIdEndpoint(bookSvc),
        decodeGetBookPublishersByIdRequest,
        encodeResponse,
    )
    DeleteBookHandler := httptransport.NewServer(
        makeDeleteBookEndpoint(bookSvc),
        decodeDeleteBookRequest,
        encodeResponse,
    )
    UpdateBookHandler := httptransport.NewServer(
        makeUpdateBookendpoint(bookSvc),
        decodeUpdateBookRequest,
        encodeResponse,
    )
    // AUTHOR
    CreateAuthorHandler := httptransport.NewServer(
        makeCreateAuthorEndpoint(authorSvc),
        decodeCreateAuthorRequest,
        encodeResponse,
    )
    GetByAuthorIdHandler := httptransport.NewServer(
        makeGetAuthorByIdEndpoint(authorSvc),
        decodeGetAuthorByIdRequest,
        encodeResponse,
    )
    // /author/{id}/books
    GetBooksByAuthorIdHandler := httptransport.NewServer(
        makeGetAuthorBooksByIdEndpoint(authorSvc),
        decodeGetAuthorBooksByIdRequest,
        encodeResponse,
    )
    DeleteAuthorHandler := httptransport.NewServer(
        makeDeleteAuthorEndpoint(authorSvc),
        decodeDeleteAuthorRequest,
        encodeResponse,
    )
    UpdateAuthorHandler := httptransport.NewServer(
        makeUpdateAuthorendpoint(authorSvc),
        decodeUpdateAuthorRequest,
        encodeResponse,
    )
    // PUBLISHER
    CreatePublisherHandler := httptransport.NewServer(
        makeCreatePublisherEndpoint(publisherSvc),
        decodeCreatePublisherRequest,
        encodeResponse,
    )
    GetByPublisherIdHandler := httptransport.NewServer(
        makeGetPublisherByIdEndpoint(publisherSvc),
        decodeGetPublisherByIdRequest,
        encodeResponse,
    )
    // /publisher/{id}/books
    GetBooksByPublisherIdHandler := httptransport.NewServer(
        makeGetPublisherBooksByIdEndpoint(publisherSvc),
        decodeGetPublisherBooksByIdRequest,
        encodeResponse,
    )
    DeletePublisherHandler := httptransport.NewServer(
        makeDeletePublisherEndpoint(publisherSvc),
        decodeDeletePublisherRequest,
        encodeResponse,
    )
    UpdatePublisherHandler := httptransport.NewServer(
        makeUpdatePublisherendpoint(publisherSvc),
        decodeUpdatePublisherRequest,
        encodeResponse,
    )

    http.Handle("/", r)
    // BOOK
    http.Handle("/book", CreateBookHandler)
    http.Handle("/book/update", UpdateBookHandler)
    r.Handle("/book/{bookid}", GetByBookIdHandler).Methods("GET")
    r.Handle("/book/{bookid}/authors", GetAuthorsByBookIdHandler).Methods("GET")
    r.Handle("/book/{bookid}/publishers", GetPublishersByBookIdHandler).Methods("GET")
    r.Handle("/book/{bookid}", DeleteBookHandler).Methods("DELETE")
    // AUTHOR
    http.Handle("/author", CreateAuthorHandler)
    http.Handle("/author/update", UpdateAuthorHandler)
    r.Handle("/author/{authorid}", GetByAuthorIdHandler).Methods("GET")
    r.Handle("/author/{authorid}/books", GetBooksByAuthorIdHandler).Methods("GET")
    r.Handle("/author/{authorid}", DeleteAuthorHandler).Methods("DELETE")
    // PUBLISHER
    http.Handle("/publisher", CreatePublisherHandler)
    http.Handle("/publisher/update", UpdatePublisherHandler)
    r.Handle("/publisher/{publisherid}", GetByPublisherIdHandler).Methods("GET")
    r.Handle("/publisher/{publisherid}/books", GetBooksByPublisherIdHandler).Methods("GET")
    r.Handle("/publisher/{publisherid}", DeletePublisherHandler).Methods("DELETE")

    // http.Handle("/metrics", promhttp.Handler())
    logger.Log("msg", "HTTP", "addr", ":"+os.Getenv("PORT"))
    logger.Log("err", http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
