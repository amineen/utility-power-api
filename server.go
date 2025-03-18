
package main

import (
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
)


func main() {

	// serverMux := http.NewServeMux()

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port:=os.Getenv("PORT")
	fmt.Printf("Server running on port %s\n", port)
	
	// http.HandleFunc("/", greet)

	// log.Printf("serving http://%s\n", *addr)
	// log.Fatal(http.ListenAndServe(*addr, nil))
}
