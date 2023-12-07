package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler (w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w,"404",http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello")
}

func formHandler (w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "post request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name = %s", name)
	fmt.Fprintf(w, "address = %s", address)

}


func main() {
	fileserver := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("starting server")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}