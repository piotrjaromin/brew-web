package main

import (
	"net/http"
	"log"
)

func main() {

	mux := http.NewServeMux()
	mux.Handle("/",  http.FileServer(http.Dir("public")))

	log.Println("Listening...")
	err := http.ListenAndServe(":3001", mux)
	if err != nil {
		log.Println("Error occured %+v", err)
	}

}