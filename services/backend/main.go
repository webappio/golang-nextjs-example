package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var resp []byte
		if req.URL.Path == "/status" {
			resp = []byte(`{"status": "ok"}`)
		} else if req.URL.Path == "/username" {
			resp = []byte(`{"username": "colin"}`)
		} else {
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.Header().Set("Content-Length", fmt.Sprint(len(resp)))
		rw.Write(resp)
	})

	log.Println("Server is available at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
