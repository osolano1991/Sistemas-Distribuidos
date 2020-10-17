package main

import (
    "net/http"
    "os"
)

func handler(writer http.ResponseWriter, request *http.Request) {
    var err error
    readData("Mass Shootings Dataset Ver 5.csv")
    switch request.Method {
    case "GET":
        err = handleGet(writer, request)
    case "POST":
        err = handlePost(writer, request)
    case "PUT":
        err = handlePut(writer, request)
    case "DELETE":
        err = handleDelete(writer, request)
    }
    writeData("Mass Shootings Dataset Ver 5.csv")
    if err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
    }
}

func main() {
    http.HandleFunc("/book/", handler)
    http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
