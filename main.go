package main

import (
	"fmt"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("/mnt/shared"))

	mux := http.NewServeMux()

	mux.Handle("/", fileServer)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}

}
