package main

import "os"

func main() {
	//var data = [][] string{{"line3", "line4"},{"aaa", "bbb"}}
	//s := string(byteA())
	file, err := os.OpenFile("result.csv", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if _, err = file.WriteString("Hello World"); err != nil {
		panic(err)
	}
}
