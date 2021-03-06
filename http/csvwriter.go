package main

import (
	"encoding/csv"
	"log"
	"os"
)

var data = [][]string{{"Line1", "Line2","Line3"}, {"Hello1 Readers of:", "golangcode.com"}}

func main() {
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
