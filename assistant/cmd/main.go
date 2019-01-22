package main

import (
	"net/http"

	"github.com/budougumi0617/sandbox-cloud-functions/assistant"
)

func main() {
	http.HandleFunc("/", assistant.Hello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
