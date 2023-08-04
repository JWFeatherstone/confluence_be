package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	http.ListenAndServe(":8080", nil)


}