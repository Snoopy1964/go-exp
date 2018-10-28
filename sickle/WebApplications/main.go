package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello Folks!")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})

	http.ListenAndServe(":8000", nil)
}
