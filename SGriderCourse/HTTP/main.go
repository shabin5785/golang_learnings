package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logwriter struct{}

func main() {
	resp, err := http.Get("http://www.google.com")

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)

	// resp.Body.Read(bs)

	// fmt.Println(string(bs))

	lw := logwriter{}

	// io.Copy(os.Stdout, resp.Body)

	// using custom writer
	io.Copy(lw, resp.Body)

}

func (logwriter) Write(bs []byte) (int, error) {
	fmt.Println("hahaha")
	fmt.Println(string(bs))
	return len(bs), nil
}
