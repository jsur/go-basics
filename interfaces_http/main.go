package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	ln := logWriter{}
	// destination is a Writer, src is a Reader
	io.Copy(ln, resp.Body)
}

// logWriter implements the Writer interface now
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
