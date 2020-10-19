package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Go Docker Service")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola Mundo desde GO lang")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
