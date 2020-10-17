package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	//"github.com/gorilla/mux"
)

func find(x string) int {
	for i, shooting := range shootings {
		if x == shooting.Id {
			return i
		}
	}
	return -1
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	fmt.Println("Id:", id)
	if id == "shooting" {
		/* w.Header().Set("Content-Type", "application/json")
		   json.NewEncoder(w).Encode(shootings)*/
		w.Header().Set("Content-Type", "application/json")
		dataJson, _ := json.Marshal(shootings)
		w.Write(dataJson)
	}
	checkError("Parse error", err)
	// fmt.Println("Id:",id)
	i := find(id)
	if i == -1 {
		fmt.Println("Id:", i)
		if id != "shooting" {
			fmt.Fprintf(w, "No se encuentra el ID %v", id)
		}
		return
	}
	dataJson, err := json.Marshal(shootings[i])
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	shooting := Shooting{}
	json.Unmarshal(body, &shooting)
	shootings = append(shootings, shooting)
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Content-Type", "application/json")
	id := path.Base(r.URL.Path)
	// var Shooting Shooting
	/*  var Shooting Shooting

	    reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
		}
		json.Unmarshal(reqBody, &Shooting)

	    for i, item := range shootings {
			if item.Id == id {
	            item.Title = Shooting.Title
	            item.Edition = Shooting.Edition
	            item.Copyright = Shooting.Copyright
	            item.Language = Shooting.Language
	            item.Pages = Shooting.Pages
	            item.Author = Shooting.Author
	            item.Publisher = Shooting.Publisher
	            shootings  = append(shootings [:i], item)
	            json.NewEncoder(w).Encode(&item)
	           // var Shooting Shooting
				//_ = json.NewDecoder(r.Body).Decode(&Shooting)
				//json.NewDecoder(r.Body).Decode(&Shooting)
				//Shooting.Id = id
				//Shooting.Id = params["id"]
			//	shootings = append(shootings, Shooting)
				//json.NewEncoder(w).Encode(&Shooting)
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
	   shootings = append(shootings[:i], shootings[i+1:]...)
	   shooting.Id = i
	   shootings = append(shootings, shooting)*/

	//Funcionando
	for index, item := range shootings {
		if item.Id == id {
			shootings = append(shootings[:index], shootings[index+1:]...)
			var Shooting Shooting
			_ = json.NewDecoder(r.Body).Decode(&Shooting)
			Shooting.Id = id
			shootings = append(shootings, Shooting)
			json.NewEncoder(w).Encode(&Shooting)
			return
		}
	}

	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	//OTRA FORMA DE HACERLO
	/* id := path.Base(r.URL.Path)
	    for index, item := range shootings {
			if item.Id == id {
				shootings = append(shootings[:index], shootings[index+1:]...)
				break
			}
	    }*/
	id := path.Base(r.URL.Path)
	i := find(id)
	if i == -1 {
		fmt.Println("Id Invalido")
		// fmt.Fprintf(w,"Id Invalido %v",i)
	}
	shootings = append(shootings[:i], shootings[i+1:]...)
	w.WriteHeader(200)
	return
}
