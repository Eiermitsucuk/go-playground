package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	//targetURL := mylib.ServerURL
	c := http.Client{Timeout: time.Duration(10) * time.Second}
	resp, err := c.Get("https://www.google.com/")
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Printf("Body : %s", body)

}
