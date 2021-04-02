package main

import (
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	seedUserData("data/user_seed_data.json")

	// register routes
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/users", UsersHandler)
	http.HandleFunc("/subscribe", SubscribeHandler)
	http.HandleFunc("/wiki", WikiHandler)

	// serve up static content
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	// start the web server
	http.ListenAndServe(":8080", nil)
}
