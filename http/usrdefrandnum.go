package main

import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
	"log"
)

func main(){
	a := numinput2()
	b := genmultiple1(a)
	//fmt.Print(b)
	c := randomnumber1(b, a)
	fmt.Print(c)

}
func numinput2() int {
	var i int
	fmt.Print("Enter the number of digit")
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println(err)
	}
	return i
}
func genmultiple1(a int) int {
	var x string
	b := "0"
	j := "1"
	for i:=0;i<a;i++  {
		x = x + b
	}
	j = j + x
	u, err := strconv.Atoi(j)
	if err != nil {
		log.Fatal("Error")
	}
	return u
}
func randomnumber1(z int, x int) string {
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
