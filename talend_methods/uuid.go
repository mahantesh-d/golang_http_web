package main

import (
	"container/list"
	"fmt"
)

func main(){
	l := list.New()
	e4 := l.PushBack("a")
	e1 := l.PushFront("b")
	e2 := l.InsertBefore(3, e4)
	e3 := l.InsertAfter(5, e2)
	l.InsertBefore(10, e3)
	l.InsertAfter(2, e1)


	for e := l.Front(); e != nil; e = e.Next(){
	//for e := l.Front(); e != nil; e = e.Next(){

		fmt.Println(e.Value)
	}
}