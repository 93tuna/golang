package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HelloWorldResponse struct {
	Message string `json: "message"`
}

type HelloWorldRequest struct {
	Name string `json: "name"`
}

const port = 8080

func main() {
	server()
}

func server() {
	http.HandleFunc("/helloworld", HelloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var request HelloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := HelloWorldResponse{Message: "Hello" + request.Name}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
