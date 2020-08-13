package main

import (
    "encoding/json"
    "net/http"
    "path"
    "fmt"

    //"github.com/gorilla/mux"
)

func find(x string) int {
    for i, book := range books {
        if x == book.Id {
            return i
        }
    }
    return -1
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
    id := path.Base(r.URL.Path)
    fmt.Println("Id:",id)
    if id =="book"{
       /* w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(books)*/
        w.Header().Set("Content-Type", "application/json")
        dataJson, _ := json.Marshal(books)
        w.Write(dataJson)
    }
    checkError("Parse error", err)
   // fmt.Println("Id:",id)
    i := find(id)
    if i == -1 {
        fmt.Println("Id:",i)  
        if id != "book"{
            fmt.Fprintf(w,"No se encuentra el ID %v",id) 
        }
        return
    }
    dataJson, err := json.Marshal(books[i])
    w.Header().Set("Content-Type", "application/json")
    w.Write(dataJson)
    return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
    len := r.ContentLength
    body := make([]byte, len)
    r.Body.Read(body)
    book := Book{}
    json.Unmarshal(body, &book)
    books = append(books, book)
    w.WriteHeader(200)
    return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
    w.Header().Set("Content-Type", "application/json")
    id := path.Base(r.URL.Path)
   // var Book Book
  /*  var Book Book

    reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &Book)

    for i, item := range books {
		if item.Id == id {            
            item.Title = Book.Title
            item.Edition = Book.Edition
            item.Copyright = Book.Copyright
            item.Language = Book.Language
            item.Pages = Book.Pages
            item.Author = Book.Author
            item.Publisher = Book.Publisher
            books  = append(books [:i], item)
            json.NewEncoder(w).Encode(&item)
           // var Book Book
			//_ = json.NewDecoder(r.Body).Decode(&Book)
			//json.NewDecoder(r.Body).Decode(&Book)
			//Book.Id = id
			//Book.Id = params["id"]
		//	books = append(books, Book)
			//json.NewEncoder(w).Encode(&Book)
		//	return
        }
    }*/

    //=======================================
  //  i := find(id)
   /* if i == -1 {
        // return
         fmt.Println("Id Invalido")
        // fmt.Fprintf(w,"Id Invalido %v",i)
    }
    books = append(books[:i], books[i+1:]...)  
    book.Id = i
    books = append(books, book)*/

    //Funcionando
    for index, item := range books {
		if item.Id == id {
			books = append(books[:index], books[index+1:]...)
            var Book Book
			_ = json.NewDecoder(r.Body).Decode(&Book)
			Book.Id = id
			books = append(books, Book)
			json.NewEncoder(w).Encode(&Book)
			return
        }
    }

    w.WriteHeader(200)
    return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
     //OTRA FORMA DE HACERLO
     /* id := path.Base(r.URL.Path) 
    for index, item := range books {
		if item.Id == id {
			books = append(books[:index], books[index+1:]...)
			break
		}
    }*/
    id := path.Base(r.URL.Path)
    i := find(id)
    if i == -1 {
        fmt.Println("Id Invalido")
       // fmt.Fprintf(w,"Id Invalido %v",i)
    }
    books = append(books[:i], books[i+1:]...)    
    w.WriteHeader(200)
    return
}

