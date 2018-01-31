package main

import (
	"log"
	"net/http"
	"helloworld/api/v6/demo"
)

func main() {
	router := demo.NewRouter()
	log.Fatal(http.ListenAndServe(":8083", router))
}
