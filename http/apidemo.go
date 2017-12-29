package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	url := fmt.Sprintf("https://httpbin.org/get")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Sprint(req.Body)

}
