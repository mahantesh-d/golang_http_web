package main

import (
	"os"
	"bufio"
	"fmt"
)

func contLines(f *os.File, counts map[string]int)  {
	input := bufio.NewScanner(f)
	for input.Scan(){
		counts[input.Text()]++
	}
}

func main(){
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0{
		contLines(os.Stdin, counts)
	}else{
		for _,arg := range files{
			f,err := os.Open(arg)
			if err != nil{
				fmt.Fprint(os.Stderr,"dup2 %v\n", err)
				continue
			}
			contLines(f, counts)
			f.Close()
		}
	}
	for line,n := range counts{
		if n>1{
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}