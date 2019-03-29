package main

import (
	"log"
	"net/http"

	"github.com/rfsx0829/security-code/server"
)

func main() {
	http.HandleFunc("/check", server.Check)
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
