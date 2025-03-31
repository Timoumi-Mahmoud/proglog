package main

import (
	"log"

	"github.com/Timoumi-Mahmoud/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":6969")
	log.Println("Starting the server on port 6969 ......")
	log.Fatal(srv.ListenAndServe())
}
