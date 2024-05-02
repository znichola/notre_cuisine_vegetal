package main

import (
	"fmt"
	"net/http"
	// "fmt"
	// "log"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("../../static")))
	fmt.Println("server started on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
