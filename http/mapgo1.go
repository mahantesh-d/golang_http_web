package main

import "fmt"

func main(){
	c := make(map[string]int)
	c["a"] = 1
	c["b"] = 2
	c["c"] = 3
	fmt.Print(c)
}
