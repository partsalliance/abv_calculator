package main

import (
	"io"
	"net/http"
	"strconv"
)

func abv(res http.ResponseWriter, req *http.Request) {

	og_string := req.FormValue("og")
	og, err := strconv.ParseFloat(og_string, 64)
	if err != nil {
		// do something sensible
	}

	fg_string := req.FormValue("fg")
	fg, err := strconv.ParseFloat(fg_string, 64)
	if err != nil {
		// do something sensible
	}
	abv_value := (og - fg) * 131.25
	string_abv_value := strconv.FormatFloat(abv_value, 'f', 2, 64)

	io.WriteString(res, string_abv_value + "%")
}


func main() {

	http.Handle("/abv", http.HandlerFunc(abv))

	http.ListenAndServe(":8080", nil)
}

