package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, val := range resp.Header {
		fmt.Println(val)
	}
	fmt.Printf("%v\n", resp.Body)
	resp.Body.Close()
}
