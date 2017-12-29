package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "http://10.138.32.76:9001/v1/alltrade/Stock_Adjustment?filters=(adjustStockNo=AJ0000726538072)(locationCode=1771)"
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal("Could not reach host")
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Could nod read the body()")
	}
	respString := string(responseData)
	fmt.Println(respString)
}
