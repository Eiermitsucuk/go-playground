package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const ServerPort = 8080
const ServerURL = "http://localhost:8080/"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(strconv.Itoa(ServerPort), nil))
}
