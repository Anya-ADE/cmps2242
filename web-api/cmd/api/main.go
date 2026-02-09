package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Shapes API"))
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Anya Andrews"))
}

func getServerTime(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("2026-02-08 15:04:05")
	w.Write([]byte(currentTime))
}

func random(w http.ResponseWriter, r *http.Request) {
	res := fmt.Sprintf("Random Number: %d", rand.Intn(100))
	w.Write([]byte(res))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/health", health)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/time", getServerTime)
	mux.HandleFunc("/random", random)

	fmt.Println("Starting server on :4000...")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
