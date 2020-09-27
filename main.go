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
    http.Handle("/", r)
    http.Handle("/book", CreateBookHandler)
    http.Handle("/book/update", UpdateBookHandler)
    r.Handle("/book/{bookid}", GetByBookIdHandler).Methods("GET")
    r.Handle("/book/{bookid}", DeleteBookHandler).Methods("DELETE")
//=============================================================================================================
//                                                     Author
//=============================================================================================================
var svc AuthorService
    svc = NewService(logger)

    // svc = loggingMiddleware{logger, svc}
    // svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}

    CreateAuthorHandler := httptransport.NewServer(
        makeCreateAuthorEndpoint(svc),
        decodeCreateAuthorRequest,
        encodeResponse,
    )
    GetByAuthorIdHandler := httptransport.NewServer(
        makeGetAuthorByIdEndpoint(svc),
        decodeGetAuthorByIdRequest,
        encodeResponse,
    )
    DeleteAuthorHandler := httptransport.NewServer(
        makeDeleteAuthorEndpoint(svc),
        decodeDeleteAuthorRequest,
        encodeResponse,
    )
    UpdateAuthorHandler := httptransport.NewServer(
        makeUpdateAuthorendpoint(svc),
        decodeUpdateAuthorRequest,
        encodeResponse,
    )
    http.Handle("/", r)
    http.Handle("/author", CreateAuthorHandler)
    http.Handle("/author/update", UpdateAuthorHandler)
    r.Handle("/author/{authorid}", GetByAuthorIdHandler).Methods("GET")
    r.Handle("/author/{authorid}", DeleteAuthorHandler).Methods("DELETE")



//=============================================================================================================
    // http.Handle("/metrics", promhttp.Handler())
    logger.Log("msg", "HTTP", "addr", ":"+os.Getenv("PORT"))
    logger.Log("err", http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}