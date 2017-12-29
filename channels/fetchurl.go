package main

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
)

func main(){
	for _, arg := range os.Args[1:]{
		resp, err := http.Get(arg)
		if err != nil{
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b,err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil{
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", arg, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
