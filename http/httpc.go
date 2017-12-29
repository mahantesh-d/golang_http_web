package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	url := "http://10.252.169.12:7788/v1/sales/salesStatistic/stockAdjustment.json?filter=(&(transactionId=200911))"
	response, err := http.Get(url)
	//response, err := http.Post(url, "", nil)
	if err != nil {
		log.Fatal(err)
	}
	//defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)
	fmt.Println(responseString)
	s := "resultCode"
	fmt.Scanf(responseString)
	fmt.Println(strings.Contains(s,responseString))
}
