package main

import (
	"io"
	"io/ioutil"
	"os"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	bs, err := ioutil.ReadFile("/content/index.html")

	if err != nil {
		fmt.Printf("Couldn't read index.html: %v", err);
		os.Exit(1)
	}

	io.WriteString(w, string(bs[:]))
}

func main() {
	http.HandleFunc("/", index)
	port := ":8000"
	fmt.Printf("Starting to service on port %s\n", port);
	http.ListenAndServe(port, nil)
}
