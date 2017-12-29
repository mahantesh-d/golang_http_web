package main

import "fmt"

type empdetail struct {
	name string
	age int
}

func main(){
	fmt.Print(empdetail{"Mahantesh", 32})

	fmt.Print("\n",empdetail{name:"Askshay", age:24})

	fmt.Print("\n",empdetail{name:"Nishith"})

	fmt.Print("\n",&empdetail{name:"Rohini", age:29})

	s := empdetail{name:"Kavitha", age:30}
	fmt.Print("\n",s.name)

	sp := &s
	fmt.Print("\n",sp.age)

	sp.age = 32
	fmt.Print("\n",sp.age)
}
