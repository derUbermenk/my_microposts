package main

import (
	"fmt"
	"net/http"
	"time"

	"my_microposts/handlers"
)

func main() {
	fmt.Println(`My Microposts`, `started at`, config.Address)

	mux := http.NewServeMux()

	// make root as micropost route
	mux.HandleFunc(`/`, handlers.Microposts.Index)
	mux.HandleFunc(`/microposts/new`, handlers.Microposts.New)
	mux.HandleFunc(`/microposts/create`, handlers.Microposts.Create)

	// starting up the server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
