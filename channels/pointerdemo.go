package main

import "fmt"

func main(){
	x := 1
	p := &x
	fmt.Print("value of x: ", x, "\nAddress of x: ", p)
	x = 3
	*p = 2
	fmt.Print("\nNew value of x: ", x)
	var y int
	z := &y
	fmt.Print("\nAddress y: ", z, "\nValue of y: ", *z)

}
