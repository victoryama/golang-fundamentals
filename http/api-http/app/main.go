package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello: ", w)
}

func headers(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Headers: ", w)

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	fmt.Println("Initializing Server")

	http.HandleFunc("GET /hello", hello)
	http.HandleFunc("GET /headers", headers)

	http.ListenAndServe(":8090", nil)
}
