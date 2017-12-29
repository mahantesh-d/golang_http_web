package main

import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
)

func main(){
	a := numinput1()
	b := genmultiple(a)
	c := randomnumber(b, a)
	fmt.Print(c)

}
func numinput1() int {
	var i int
	fmt.Print("Enter the number of digit")
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println(err)
	}
	return i
}
func genmultiple(a int) int {
	b := 1
	j := 10
	for i:=0;i<a;i++  {
		b = b * j
	}
	return b
}
func randomnumber(z int, x int) string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(z)
	c := strconv.Itoa(a)
	f := x - len(c)
	g := "0"
	h := c
	for j := 0; j < f; j++ {
		h = g + h
	}
	return h
}
