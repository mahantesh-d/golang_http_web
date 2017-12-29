package main

import "fmt"

func generateCombination1(alphabet string, length int) <- chan string {
	c := make(chan string)
	go func(c chan string) {
		defer close(c)
		addLetter1(c, "", alphabet, length)
	}(c)
	return c
}

func addLetter1(c chan string, combo string, alphabet string, length int)  {
	if length <=0 {
		return
	}
	var newcombo string
	for _, ch := range alphabet{
		if combo != string(ch) {
			newcombo = combo + string(ch)
			c <- newcombo
			addLetter1(c, newcombo, alphabet, length - 1)
		}
	}

}


func main()  {
	for generateCombo := range generateCombination1("abc", 3){
		fmt.Print("*******************\n")
		fmt.Print(/*"final: ",*/generateCombo, "\n")
		fmt.Print("*******************\n")
	}
	fmt.Print("Done....!")
}