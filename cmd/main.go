package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Print("starting server on port 3000")
	hellohandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "General Kenobi\n")
	}

	http.HandleFunc("/hellothere", hellohandler)
	http.ListenAndServe(":3000", nil)
}