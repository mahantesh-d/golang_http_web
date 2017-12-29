package main

import "fmt"

func main() {
	heading1("Set")
	subheading1("Set")
	a := setg()
	fmt.Print(a)
}

func setg() []string {
	s := make([]string, 3)
	s[0] = "Asus"
	s[1] = "Device"
	s[2] = "Handset"
	return s
}

func heading1(b string) {
	fmt.Println("-----------------------------------------------------------")
	fmt.Println(b)
	fmt.Println("-----------------------------------------------------------")
}

func subheading1(c string) {
	fmt.Println("-----" + c + "---------------------------------------------")

}
