package main

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func main() {

	fileServer := http.FileServer(http.Dir(filepath.Dir(`c:\Downloads`)))

	mux := http.NewServeMux()

	mux.Handle("/", fileServer)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}

}
