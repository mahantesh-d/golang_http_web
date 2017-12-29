package main

import (
	"bufio"
	"fmt"
	"os"
)

func ParseLines(filePath string, parse func(string) (string,bool)) ([]string, error) {
	inputFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	var results []string
	for scanner.Scan() {
		if output, add := parse(scanner.Text()); add {
			results = append(results, output)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: line_parser <path>")
		return
	}

	lines, err := ParseLines(os.Args[1], func(s string)(string,bool){
		return s, true
	})
	if err != nil {
		fmt.Println("Error while parsing file", err)
		return
	}

	for _, l := range lines {
		fmt.Println(l)
	}
}