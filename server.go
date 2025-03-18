package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/amineen/utility-api/handlers"
)

func main() {

	// serverMux := http.NewServeMux()

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	fmt.Printf("Server running on port %s\n", port)

	http.HandleFunc("/", Home)

	// log.Printf("serving http://%s\n", *addr)
	// log.Fatal(http.ListenAndServe(*addr, nil))
}
