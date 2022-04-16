package main

import (
	// "net/http"
	// "net/url"
	"fmt"
	"api_client/client_tool"
	"log"
	"time"
	// "errors"
)
func main() {
	defaultTimeout := time.Second * 10
	client := client_tool.NewClient(
		"https://jsonplaceholder.typicode.com",
		defaultTimeout,
	)
	data, err := client.GetPosts()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("ID: %d, title is ... %v\n", data[i].ID, data[i].Title)
	}
}