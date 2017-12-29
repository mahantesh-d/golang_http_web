package main

import (
	"strconv"
	"time"
	"fmt"
	"bufio"
	"os"
)

func main(){
	StringRandomInt()
}

func StringRandomInt() {
	b := input()
	heading("String Random")
	subheading(b)
	if b == "RT" || b == "TS" || b == "SS" || b == "RW" || b == "OS" {
		a := test13(b)
		fmt.Println(a)
	} else if b == "IN" {
		a := test14(b)
		fmt.Println(a)
	} else if b == "Pu" || b == "C" {
		a := test10(b)
		fmt.Println(a)
	} else {
		fmt.Println("err")
	}

	func input() string{
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter text: ")
		scanner.Scan()
		text := scanner.Text()
		return text
	}

	func test13(i string) string{
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		a := r1.Intn(10000000000000)
		c := strconv.Itoa(a)
		f := 13 - len(c)
		g := "0"
		h := c
		for j := 0; j<f; j++{
		h = g + h
	}
		return i+h
	}

	func test14(i string) string{
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		a := r1.Intn(100000000000000)
		c := strconv.Itoa(a)
		f := 14 - len(c)
		g := "0"
		h := c
		for j := 0; j<f; j++{
		h = g + h
	}
		return i+h
	}

	func test10(i string) string{
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		a := r1.Intn(10000000000)
		c := strconv.Itoa(a)
		f := 10 - len(c)
		g := "0"
		h := c
		for j := 0; j<f; j++{
		h = g + h
	}
		return i+h
	}

	func heading(b string){
	fmt.Println("-----------------------------------------------------------")
	fmt.Println(b)
	fmt.Println("-----------------------------------------------------------")
	}

	func subheading(c string){
	fmt.Println("-----" + c + "---------------------------------------------")

	}
}