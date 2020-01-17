package main

import (
	"log"
	"net/http"
)

func main() {
	// initDb()
	// defer db.Close()
	// http.HandleFunc("/api/index", indexHandler)
	// http.HandleFunc("/api/repo/", repoHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// func indexHandler(w http.ResponseWriter, r *http.Request) {
//     //...
// }

// func repoHandler(w http.ResponseWriter, r *http.Request) {
//     //...
// }
