package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"io/ioutil"
)
/******************************Generate the Combination********************************/
func generatecomb(alpha []string, length int) <-chan string {
	c := make(chan string)
	go func(c chan string) {
		defer close(c)
		Addlet(c, "", alpha, length)
	}(c)
	return c
}
/******************************Add the fields*******************************************/
func Addlet(c chan string, combo string, alpha []string, length int) {
	if length <= 0 {
		return
	}
	var newcombo string
	for _, ch := range alpha {
		newcombo = combo + string(ch)
		c <- newcombo
		Addlet(c, newcombo, alpha, length-1)
	}
}
/*******************************Create the URL from above two function**********************/
func urlcreation(b string) {

	a := "http://10.138.32.76:9001/v1/alltrade/obtainDetail?filters=(" + b + ")"
	url := string(a)
	httpreq(url)
	//fmt.Print(url, "\n")
}
/********************************Send the HTTP request using above URL*********************/
func httpreq(c string) {
	fmt.Print("******************", c, "**************************\n")
	resp, err := http.Get(c)
	if err != nil {
		log.Fatal("Could not read the host")
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	scanner.Split(bufio.ScanRunes)
	var buf bytes.Buffer
	for scanner.Scan() {
		buf.WriteString(scanner.Text())
	}
	writetofile(buf.String())
}
/*********************************Write the data to file**************************************/
func check(err error) {
	if err != nil {
		panic(err)
	}
}
func writetofile(a string){
	f, err := os.Create("data.txt")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = fmt.Fprintf(w, "%v\n", a)
	check(err)
	w.Flush()
}
/***********************************Search for count in the file********************************/
func searchstring(){
	var text string
	//fmt.Print("Enter text: ")
	// get the sub string to search from the user
	//fmt.Scanln(&text)
	text = "AJ0000952226834"

	// read the whole file at once
	b, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	s := string(b)
	// //check whether s contains substring text
	fmt.Println(strings.Contains(s, text))
}
/***********************************Function Main***********************************************/
func main() {
	var arr = []string{"(obtainNo=PU1177160500009)","(locationCode=1177)"}
	count := 0
	var inr int
	for combination := range generatecomb(arr, 2) {
		urlcreation(combination)
		inr = count + 1
		count++
	}
	searchstring()
	fmt.Print("Total Combination: ", inr, "\n")
	fmt.Print("Done!")
}
