package main

import (
	"fmt"
	"net/http"
	"time"

	"my_microposts/routes"
)

func main() {
	fmt.Println(`My Microposts`, `started at`, config.Address)

	mux := http.NewServeMux()

	// make root as micropost route
	mux.HandleFunc(`/`, routes.Microposts.Index)
	mux.HandleFunc(`/microposts/new`, routes.Microposts.New)

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
