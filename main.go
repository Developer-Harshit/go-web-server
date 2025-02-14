package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "embed"
)

//go:embed page.html
var page string

func main() {
	http.HandleFunc("/test", test)
	http.HandleFunc("/", serveHello)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server listening on http://localhost:" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Error starting server:", err)
	}
}
func test(rw http.ResponseWriter, req *http.Request) {
    body, err := io.ReadAll(req.Body)
    if err != nil {
        panic(err)
    }
    log.Println(string(body))
    fmt.Fprint(rw,"hi")

}

func serveHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, page)
}
