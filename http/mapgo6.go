package main

import (
	"strings"
	"io/ioutil"
	"fmt"
)

type tablemetrics struct {
	keyspace string
	tablename string
	spacelive string
	spaceused string
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main()  {
//	var z[]string
	data, err := ioutil.ReadFile("cfstats.out")
	check(err)
	file := string(data)
	temp := strings.Split(file,"\n")
	a := []string{"Keyspace","Table:","Table (index)","Space used (live)","Space used (total)"}
	for _,lines := range temp{
		for _,key := range a {
			if strings.Contains(lines, key){

				fmt.Print("\n",lines)
			}

		}
	}
}
