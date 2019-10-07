package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Welcome)
	err := http.ListenAndServeTLS(":8080", "https-server.crt", "https-server.key", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"status":"200OK"}`)
}
