package main

import (
	"go/token"
	"go/parser"
	"fmt"
)

func main(){
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "randomnumber.go", nil, parser.ImportsOnly)
	if err != nil{
		fmt.Print(err)
		return
	}
	for _, s := range f.Imports {
		fmt.Println(s.Path.Value)
	}
}
