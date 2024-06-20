package main

import (
	"fmt"
	"log"
	"net/http"
)

var counter int = 0

func main() {
	port := 3000
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Printf("Starting request with counter %v\n", counter)
		counter++
		res.Header().Set("Access-Control-Allow-Origin", "*")
		h := http.FileServer(http.Dir("files"))
		fmt.Printf("Starting request with counter %v\n", counter)
		h.ServeHTTP(res, req)
	})

	fmt.Printf("Starting server on %v\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
