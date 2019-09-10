package main

import (
	"fmt"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
	return ":" + p
	}
	return ":3000"
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server Starting...")
	fmt.Println("Running in port", getPort())
	http.ListenAndServe(getPort(), nil)

}
