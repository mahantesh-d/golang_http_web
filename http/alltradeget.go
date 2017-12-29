package main

import (
	//"net/http"
	//"log"
	//"io/ioutil"
	"bufio"
	"fmt"
	"os"
	//	"net/http"
	//	"log"
	//	"io/ioutil"
)

func main() {
	//var b []string
	b := filter()
	fmt.Print(b)
	//	httpresp(b)
}

func filter() []string {
	var c []string
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the value for adjust_stock_no")
	a, _ := scanner.ReadString
	fmt.Print("Enter the value for location_code")
	b, _ := scanner.ReadString
	c = []string{a, b}
	return c
}

/*func httpresp(a string, b string){
	c := "http://10.138.32.76:9001/v1/alltrade/Stock_Adjustment?filters=(adjustStockNo="+a+")"
	url := string(c)
	resp, err := http.Get(url)
	if err != nil{
		log.Fatal("Could not reach:"+url)
	}
	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if  err != nil {
		log.Fatal("Error")
	}
	respString := string(respData)
	fmt.Print(respString)
}*/
