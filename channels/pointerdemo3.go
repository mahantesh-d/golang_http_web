package main

import "fmt"

func demof(p *int) int{
	*p++
	return *p
}

func main(){
	q := 1
	r :=demof(&q)
	fmt.Print(r)
	z := q+1
	fmt.Print("\nZ: ", z)
}
