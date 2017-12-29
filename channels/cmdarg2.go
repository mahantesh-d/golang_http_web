package main

import (
	"fmt"
	//"strings"
	"os"
	"strings"
	"time"
)

func main(){
	fmt.Println(strings.Join(os.Args[1:],time.Minute,))
	//fmt.Println(os.Args[0])
}
