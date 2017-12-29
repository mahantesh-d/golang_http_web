package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://httpbin.org/post?b=mantu"
	resp, err := http.Post(url, "", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	responseString := string(responseData)
	fmt.Println(responseString)
}
