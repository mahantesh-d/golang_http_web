package main

import "fmt"

type addrdetails struct {
	city string
	street string
}

func main(){
	names := []string{"Rohini", "Ravi"}
	companies := []string{"Unotech", "Entronica"}
	m := make(map[string]map[string]addrdetails)
	for _, company := range companies{
		for _,name := range names{
			//fmt.Print("*********", company, name)
			m[company] = make(map[string]addrdetails)
			m[company][name] = addrdetails{city:"Bengaluru", street:"Rajkumar Road"}
			//fmt.Println("\n",m)
			break
		}
		fmt.Println("\n------",m)
	}
	//fmt.Println("\n",m)
}
