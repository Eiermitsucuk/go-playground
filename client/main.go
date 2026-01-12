package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
