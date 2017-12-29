package main

import "fmt"

func main(){
	a := []int{5,1,6,2,4,3}
	j := len(a)
	for d:=0;d<j;d++ {
		for i := 0; i < j - 1; i++ {
			//if a[i] < a[i + 1] {
			if a[i] > a[i + 1] {
				tmp := a[i]
				a[i] = a[i + 1]
				a[i + 1] = tmp
			}
		}
	}
	for i:=0;i<j ;i++  {
		fmt.Print(a[i])
	}
}
