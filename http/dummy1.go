package main

import (
	"io/ioutil"
	"strings"
	"os"
	"encoding/csv"
)
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main(){
	var z []string
	data, err := ioutil.ReadFile("cfstats.out") // just pass the file name
	check(err)
	file := string(data) // convert content to a 'string'
	temp := strings.Split(file,"\n")
	a := []string{"Keyspace","Table:","Table (index)","Space used (live)","Space used (total)"}
	f, err := os.Create("result.csv")
	check(err)
	defer f.Close()
	writer := csv.NewWriter(f)
	for _,lines := range temp {
		for _,key := range a {
			if strings.Contains(lines, key) {
				z = []string{lines}
				err := writer.Write(z)
				check(err)
				//err := writer.WriteAll([][]string{{key},{lines}})
				//check(err)
				//fmt.Print("\n")
				//fmt.Print(lines)
			}
			//err := writer.Write(z)
			//check(err)
		}
		//err := writer.Write(z)
		//check(err)
	}
}

	/*x := len(z)
	fmt.Print(x)

	a := []string{"Keyspace","Table (index)","Space used (live)","Space used (total)"}
	f, err := os.Create("result.csv")
	check(err)
	defer f.Close()
	writer := csv.NewWriter(f)
	for _,item := range z {
		for _,key := range a {
			if strings.Contains(item, key) {
				//err := writer.Write(item)
				//check(err)
				fmt.Println(item)
			}
		}
	}*/

