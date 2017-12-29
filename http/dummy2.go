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
				if key == "Keyspace" {
					z = []string{">>---",lines,"---<<"}
					err := writer.Write(z)
					check(err)
				}else if key == "Table:" || key == "Table (index)" {
					z = []string{">---",lines,"---<"}
					err := writer.Write(z)
					check(err)
				}else {
					z = []string{lines}
					err := writer.Write(z)
					check(err)
				}
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
