package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	resp, err := client.Get("http://localhost:8090/hello")

	if err != nil {
		fmt.Println("Erro during GET request: ", err)
		return
	}
	fmt.Println(resp.Status)

	respP, err := client.Get("http://localhost:8090/headers")
	if err != nil {
		fmt.Println("Erro during GET request: ", err)
		return
	}

	fmt.Println(respP.Status)
}
