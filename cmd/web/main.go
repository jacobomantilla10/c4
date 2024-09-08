package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	routes := routes(mux)
	fmt.Println("Listening on port 3000...")
	http.ListenAndServe(":3000", routes)
}
