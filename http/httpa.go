package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Print("Start \n")
	if len(os.Args) != 2 {
		fmt.Print("if")
		fmt.Fprint(os.Stderr, "Usage: %s URL\n", os.Args[0])
		os.Exit(1)
	}
	fmt.Print("if/or \n")
	response, err := http.Get(os.Args[1])
	fmt.Print("uuu")
	if err != nil {
		fmt.Print("if1 \n")
		log.Fatal(err)
	} else {
		fmt.Print("else \n")
		//defer response.Body.Close()

		_, err := io.Copy(os.Stdout, response.Body)
		if err != nil {
			fmt.Print("if3 \n")
			log.Fatal(err)
		}
	}
}
