package main

import "fmt"

func main(){
	m := map[string]map[string]map[string]int{
		"a": map[string]map[string]int{
			"b": map[string]int{
				"one":1,
				"two":2,
			},
			"c": map[string]int{
				"three":3,
				"four":4,
			},
		},
		"aa": map[string]map[string]int{
			"bb":map[string]int{
				"five":5,
				"six":6,
			},
			"cc":map[string]int{
				"seven":7,
				"eight":8,
			},
		},

	}
	fmt.Print("\n",m)
}