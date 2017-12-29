package main

import "fmt"

func f() *int{
	v := 1
	return &v
}
func main()  {
	p := f()
	fmt.Print(&p)
	//fmt.Print("\n",p)
	fmt.Print("\n",*p)
}
