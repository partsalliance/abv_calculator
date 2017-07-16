package main

import (
	"io"
	"net/http"
)

func abv(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "abv calculator")
}


func main() {

	http.Handle("/abv", http.HandlerFunc(abv))

	http.ListenAndServe(":8080", nil)
}

