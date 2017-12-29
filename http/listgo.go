package main

import (
	//"container/list"
	"fmt"

	"strings"
)

func main() {

	/*l := list.New()
	e1 := l.PushBack("Mumbai")
	e2 := l.InsertBefore(1, e1)
	l.PushFront("Unotech")
	l.InsertBefore("Software", e1)
	//fmt.Print(e1)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}*/
	a := []string{"abc", "xyz", "1", "2"}
	fmt.Println(strings.Join(a, ","))
}
