package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main(){
	data, err := ioutil.ReadFile("sample.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	file := string(data) // convert content to a 'string'
	temp := strings.Split(file,"\n")
	a := []string{"Keyspace", "Table (index)","Space used (live)","Space used (total)", "Table"}
	for _,item := range temp {
		for _,key := range a {
			if strings.Contains(item, key) {
				fmt.Println(item)
			}
		}
	}

}


