package main

import (
	"fmt"
	"net/http"
)

func main() {
	//устанавливаем роутер, то есть, роут “/”
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	http.ListenAndServe(":80", nil)
}
