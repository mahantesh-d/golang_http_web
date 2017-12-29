package main

import "fmt"

func main(){
	a := make(map[string]map[string]map[string]int)
	a["b"] = make(map[string]map[string]int)
	a["b"]["c"] = make(map[string]int)
	a["b"]["c"]["d"] = 123
	fmt.Print(a)
}
