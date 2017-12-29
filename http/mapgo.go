package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	heading("Map")
	subheading("Map")
	mapg()
}

func mapg() {
	a := make(map[string]string)
	a["A"] = "AIS"
	a["U"] = "Unotech"
	a["E"] = "Entornica"
	b, err := json.Marshal(a)
	fmt.Print(a)
	fmt.Print(err)
	fmt.Print(b)
}
func heading(b string) {
	fmt.Println("-----------------------------------------------------------")
	fmt.Println(b)
	fmt.Println("-----------------------------------------------------------")
}

func subheading(c string) {
	fmt.Println("-----" + c + "---------------------------------------------")

}
