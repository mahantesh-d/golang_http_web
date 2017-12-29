package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	b := []string{"AIS", "AWN", "WDS"}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(len(b))
	fmt.Println(b[a])
}
