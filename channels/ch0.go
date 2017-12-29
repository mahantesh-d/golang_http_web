package main

import (
	"time"
	"fmt"
)

func main(){
	c := make(chan time.Time)
	go func(c chan time.Time) {
		c <- time.Now()
	}(c)
	d := <-c
	fmt.Print(d, "\n")

	a := make(chan []string)
	go func(a chan []string) {
		arr := []string{"a", "b", "c"}
		a <- arr

	}(a)
	e := <-a
	fmt.Print(e)



}
