package main

import (
	"log"
	"net/http"

	handlers "github.com/pradyumna001/blog-app/Handlers"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8081"
)

func init() {
	//Routing
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/posts", handlers.PostsHandler)
	http.HandleFunc("/users/search", handlers.UserSearchHandler)
	http.HandleFunc("/users", handlers.UsersHandler)
}

func main() {
	//Server
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
