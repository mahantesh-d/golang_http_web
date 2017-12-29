package main

import "fmt"

type parmeter struct {
	height int
	width int
}

func (r *parmeter)rectarea() int{
	return r.height * r.width
}

func (r parmeter)perim() int{
	return 2*r.width + 2*r.height
}

func main(){
	r := parmeter{height:10, width:5}

	fmt.Println("Area of rect: ", r.rectarea() )
	fmt.Println("Area of perim: ", r.perim())

	rp := &r

	fmt.Println("1Area of rect: ", rp.rectarea())
	fmt.Println("1Area of perim: ", rp.perim())
}
