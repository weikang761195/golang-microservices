package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "HELLO WORLD!")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

