package main

import (
	"fmt"
	"net/http"
)

func main() {
	response, _ := http.Get("http://youtube.com")

	res := make([]byte, 1024)

	response.Body.Read(res)
	fmt.Println(res)
}
