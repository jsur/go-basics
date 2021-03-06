package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	// destination is a Writer, src is a Reader
	io.Copy(os.Stdout, resp.Body)
}
