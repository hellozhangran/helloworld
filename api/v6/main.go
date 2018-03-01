package main

import (
	"helloworld/api/v6/demo"
	"log"
	"net/http"
)

func main() {
	router := demo.NewRouter()
	log.Fatal(http.ListenAndServe(":8083", router))
}
