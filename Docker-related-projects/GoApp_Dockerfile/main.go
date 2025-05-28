package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("Welcome, You are at DevOps Solutions!"))
	case "/health":
		w.WriteHeader(http.StatusNoContent)
	case "/ping":
		fmt.Println("In ping")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("pong"))
	default:
		http.Error(w, "Invalid Request!", http.StatusBadRequest)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.HandleFunc("/", handler)
	address := fmt.Sprintf(":%s", port)
	fmt.Printf("App listening at http://%s\n", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		fmt.Println(err)
	}
}