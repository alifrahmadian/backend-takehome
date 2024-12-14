package main

import (
	"fmt"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, World!")
// }

func main() {
	app := NewApp()

	err := app.Router.Run(":8080")
	if err != nil {
		fmt.Printf("Error running the server: %v\n", err)
	}
	// http.HandleFunc("/", handler)
	// fmt.Println("Server is running on http://localhost:8080")
	// http.ListenAndServe(":8080", nil)
}
