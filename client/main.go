package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080/?id=38")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, val := range resp.Header {
		fmt.Println(val)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Status)
	fmt.Printf("%s\n", body)
	resp.Body.Close()

}
