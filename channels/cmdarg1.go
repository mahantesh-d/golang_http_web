package main

import (
	"os"
	"fmt"
)

func main()  {
	s,sep:= " ", " "
	for a,arg := range os.Args[1:]{
		s += sep + arg
		sep = " "
		fmt.Print(a, s, "\n")
	}
	//fmt.Print(s)
}
