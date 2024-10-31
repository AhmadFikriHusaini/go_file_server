package main

import (
	"fmt"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./resources"))

	mux := http.NewServeMux()

	mux.Handle("/", fileServer)

	server := http.Server{
		Addr:    "localhost:80",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Server is running")

}
