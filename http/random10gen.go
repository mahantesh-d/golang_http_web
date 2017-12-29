package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(1000000000)
	c := strconv.Itoa(a)
	e := len(c)
	fmt.Println(e)
	fmt.Println(c)
	f := 10 - e
	g := "0"
	h := c
	for j := 0; j < f; j++ {
		h = g + h
	}
	fmt.Println(h)

}
