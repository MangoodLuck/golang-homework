package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func healthzGetRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	io.WriteString(w, "ok")
}

func getRequestHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	defer fmt.Printf("时间: %s HTTP返回码: %d", now.Format("2006-01-02 15:04:05"), 200)
	for k, v := range r.Header {
		w.Header().Set(k, v[0])
	}
	w.Header().Set("VERSION", os.Getenv("VERSION"))
}

func main() {
	http.HandleFunc("/req/get", getRequestHandler)
	http.HandleFunc("/healthz", healthzGetRequestHandler)
	err := http.ListenAndServe(":8005", nil)
	if err != nil {
		log.Fatal(err)
	}
}

