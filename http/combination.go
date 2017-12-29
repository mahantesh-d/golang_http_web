package main

import "fmt"

func GenerateCombinations(alphabet string, length int) <-chan string {
	c := make(chan string)

	// Starting a separate goroutine that will create all the combinations,
	// feeding them to the channel c
	go func(c chan string) {
		defer close(c) // Once the iteration function is finished, we close the channel

		AddLetter(c, "", alphabet, length) // We start by feeding it an empty string
	}(c)

	return c // Return the channel to the calling function
}

// AddLetter adds a letter to the combination to create a new combination.
// This new combination is passed on to the channel before we call AddLetter once again
// to add yet another letter to the new combination in case length allows it
func AddLetter(c chan string, combo string, alphabet string, length int) {
	// Check if we reached the length limit
	// If so, we just return without adding anything
	if length <= 0 {
		return
	}

	var newCombo string
	for _, ch := range alphabet {
		//fmt.Print("combo: ", combo, "\n")
		//fmt.Print("string(ch): ", string(ch), "\n")
		if string(ch) != combo {
			newCombo = combo + string(ch)
			fmt.Print("newCombo: ", newCombo, "\n")
		}

		c <- newCombo
		fmt.Print(c)
		//AddLetter(c, newCombo, alphabet, length-1)
	}
}

func main() {

	for combination := range GenerateCombinations("ab", 2) {
		fmt.Println(combination) // This is instead of process(combination)
	}

	fmt.Println("Done!")
}
