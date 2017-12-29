package main

import "fmt"

func generateCombination(alphabet string, length int) <- chan string {
	c := make(chan string)
	go func(c chan string) {
		defer close(c)
		addLetter(c, "", alphabet, length)
	}(c)
	return c
}

func addLetter(c chan string, combo string, alphabet string, length int)  {
	if length <=0 {
		return
	}
	var newcombo string
	/*fmt.Print("*******************\n")
	fmt.Print("newcombo: ",newcombo,"\n")
	fmt.Print("combo: ", combo, "\n")
	fmt.Print("alphabet: ",alphabet, "\n")
	fmt.Print("length: ", length, "\n")
	fmt.Print("*******************\n")*/
	for _, ch := range alphabet{
		//fmt.Print("String(ch): ", string(ch), "\n")
		//for ch:=0;ch<1;ch++{
		if combo != string(ch){
			newcombo = combo + string(ch)
		}
		c <- newcombo
		addLetter(c, newcombo, alphabet, length-1)
	}
}


func main()  {
	for generateCombo := range generateCombination("ab", 2){
		//fmt.Print("*******************\n")
		fmt.Print(/*"final: ",*/generateCombo, "\n")
		//fmt.Print("*******************\n")
	}
	fmt.Print("Done....!")
}