package main
import (
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)
func main()  {
	counts := make(map[string]int)
	//files := os.Args[1:]
	for _, filename := range os.Args[1:]{
		data,err := ioutil.ReadFile(filename)
		if err != nil{
			fmt.Fprint(os.Stderr, "dups3%v\n", err)
			continue
		}else{
			for _,line := range strings.Split(string(data), "\n"){
				counts[line]++
			}
		}
		for line, n := range counts{
			if n>1{
				fmt.Printf("%d\t%s", n, line)
			}
		}
	}
}
