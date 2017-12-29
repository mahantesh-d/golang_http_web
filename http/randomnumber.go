package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	b := numinput()
	heading1("Random Number")
	subheading1(b)
	if b == "10" {
		a := num10()
		fmt.Println(a)
	} else if b == "6" {
		a := num6()
		fmt.Println(a)
	} else if b == "4" {
		a := num4()
		fmt.Println(a)
	} else if b == "3" {
		a := num3()
		fmt.Println(a)
	} else {
		fmt.Println("err")
	}

}

func numinput() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter text: ")
	scanner.Scan()
	text := scanner.Text()
	return text
}

func num10() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(10000000000)
	c := strconv.Itoa(a)
	f := 10 - len(c)
	g := "0"
	h := c
	for j := 0; j < f; j++ {
		h = g + h
	}
	return h
}

func num6() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(1000000)
	c := strconv.Itoa(a)
	f := 6 - len(c)
	g := "0"
	h := c
	for j := 0; j < f; j++ {
		h = g + h
	}
	return h
}

func num4() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(10000)
	c := strconv.Itoa(a)
	f := 4 - len(c)
	g := "0"
	h := c
	for j := 0; j < f; j++ {
		h = g + h
	}
	return h
}

func num3() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(1000)
	c := strconv.Itoa(a)
	f := 3 - len(c)
	g := "0"
	h := c
	for j := 0; j < f; j++ {
		h = g + h
	}
	return h
}

func heading1(b string) {
	fmt.Println("-----------------------------------------------------------")
	fmt.Println(b)
	fmt.Println("-----------------------------------------------------------")
}

func subheading1(c string) {
	fmt.Println("-----" + c + "---------------------------------------------")

}
