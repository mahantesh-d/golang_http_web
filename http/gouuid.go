package main

import (
	//"fmt"
	//"github.com/satori/go.uuid"
	"encoding/csv"
	"log"
	"os"
)

func main() {
	//a := uuid.NewV4()
	//c := a.String()
	//fmt.Print(c)
	//var data  = [][] string{{"UUID"},{c}}
	var data = [][]string{{"line1", "line2"}, {"abc", "xyz"}}
	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)

	for _, value := range data {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}

	defer writer.Flush()
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
