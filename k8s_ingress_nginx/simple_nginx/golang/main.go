package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	podName, _ := os.LookupEnv("HOSTNAME")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf("I'm Golang MS - POD: %s", podName))
	})

	http.ListenAndServe(":3000", nil)
	fmt.Println("Server running on 3000")
}
