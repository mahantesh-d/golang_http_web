package main

import (
	"fmt"
	"time"
)

func main()  {
	t := 1
	switch t {
	case 1:
		fmt.Print("Its one\n")
	case 2:
		fmt.Print("Its two\n")
	default:
		fmt.Print("Nothing to print\n")
	}

	r := time.Now().Weekday()
	switch r {
	case time.Saturday, time.Sunday:
		fmt.Print("Its Weekend\n")
	default:
		fmt.Print("Its Weekday\n")
	}

	s := time.Now()
	switch {
	case s.Hour() < 12:
		fmt.Print("Its Morning")
	default:
		fmt.Print("Its Noon")
	}
	whatIam := func(i interface{}){

		switch t := i.(type) {
		case bool:
			fmt.Print("Boolean")
		case int:
			fmt.Print("Integer")
		case string:
			fmt.Print("String ", t)
		}
	}
	whatIam(false)
	whatIam(1)
	whatIam("Hello")
}