package main

import (
	"fmt"
	"os"
)

func main(){
	//fmt.Print(os.Args[1])
	var s, sep string
	for i:=1;i<len(os.Args);i++{
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Print(s)
}