/*
Exercise 1.9: Modify fetch to also print the HTTP status code,
found in resp.Status.
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		io.Copy(os.Stdout, resp.Body)

		fmt.Println("")
		fmt.Println("Returned status code:", resp.Status)
		resp.Body.Close()

	}
}
