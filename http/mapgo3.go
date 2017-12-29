package main

import (
	"io/ioutil"
	"strings"
	"fmt"
)
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main(){
//	b := make(map[string]map[string]map[string]map[string]map[string]string)
	var z []string
	data, err := ioutil.ReadFile("cfstats.out") // just pass the file name
	check(err)
	file := string(data) // convert content to a 'string'
	temp := strings.Split(file,"\n")
	a := []string{"Keyspace","Table:","Table (index)","Space used (live)","Space used (total)"}
	for _,lines := range temp {
		for _,key := range a {
			if strings.Contains(lines, key) {
				z = []string{lines,"\n"}
				fmt.Print(z)
			}
		}
	}
}
