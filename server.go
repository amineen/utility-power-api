package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/amineen/utility-api/handlers"
	"github.com/joho/godotenv"
)

func main() {

	serverMux := http.NewServeMux()

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	fmt.Printf("Server running on port %s\n", port)

	serverMux.HandleFunc("/", handlers.Home)

	// Define ANSI escape codes for colors
	green := "\033[32m"
	reset := "\033[0m"

	fmt.Printf("%sServer available at http://localhost:%s%s\n", green, port, reset)
	log.Fatal(http.ListenAndServe(":"+port, serverMux))
}
