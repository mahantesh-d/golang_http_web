package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"os"
)
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main(){
	data, err := ioutil.ReadFile("cfstats.out") // just pass the file name
	check(err)
	file := string(data) // convert content to a 'string'
	temp := strings.Split(file,"\n")
	a := []string{"Keyspace","Table (index)","Space used (live)","Space used (total)"}
	f, err := os.Create("result.txt")
	check(err)
	for _,item := range temp {
			for _,key := range a {
				if strings.Contains(item, key) {
				_, err := f.WriteString(item)
				check(err)
				fmt.Println(item)
			}
		}
	}

}
