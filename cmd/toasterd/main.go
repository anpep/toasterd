package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from <tt>toasterd</tt>!</h1>")
	fmt.Fprintf(w, "Powered by KernOS (:")
}

func main() {
	fmt.Println("toasterd version 1.0")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8090", nil)
}
