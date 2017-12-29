package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"os"
)

func main() {

	file, err := os.Open("cfstats.out")
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Keyspace"){
			fmt.Print("\n",scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		}
	}
}