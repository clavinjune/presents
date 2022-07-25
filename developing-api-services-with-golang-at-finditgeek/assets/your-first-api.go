package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // HL1
		name := r.URL.Query().Get("name") // HL1
		if name == "" {                   // HL1
			name = "World" // HL1
		} // HL1

		fmt.Fprintf(w, "Hello, %s!", name) // HL1
	}) // HL1

	log.Println("Listening on port 8000")
	http.ListenAndServe(":8000", nil)

	// You can access it at http://localhost:8000/?name=John
}
