package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter text: ")
	scanner.Scan()
	text := scanner.Text()
	fmt.Println(text)

}
