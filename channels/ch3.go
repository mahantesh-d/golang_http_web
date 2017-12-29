package main

import (
	"fmt"
	"time"
)

func pinger(c chan string){
	for i:=0; ; i++{
		c <- "ping"
	}
}
func printer(c chan string)  {
	for{
		msg := <-c
		fmt.Print(msg, "\n")
		time.Sleep(time.Second * 3)
	}
}
func main() {
	c := make(chan string)
	go pinger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)
}