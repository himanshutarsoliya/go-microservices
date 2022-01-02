package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		println("Hello World !")
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oooops !!", http.StatusBadRequest)
		}
		fmt.Fprintf(rw, "Data recieved in body is %s:", data)
	})
	http.HandleFunc("/bye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Good Bye !")
	})
	http.ListenAndServe(":9090", nil)
}
