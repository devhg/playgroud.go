package m_http

//package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	//302重定向
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", "http://baidu.com")
		w.WriteHeader(302)
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
