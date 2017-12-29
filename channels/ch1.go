package main

import (
	"fmt"
)

func main(){
	c := make(chan string)
	go func(c chan string) {
		arr := []string{"a","b"}
		for i:=0; i<2; i++{
			c <- arr[i]
		}
	}(c)
	for i:=0;i<2;i++  {
		d := <- c
		fmt.Print(d, "\n")
	}
}
