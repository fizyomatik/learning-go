package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	const max = 5
	var uniques [max]int
loop:

	for f := 0; f < max; f++ {
		n := rand.Intn(max) + 1
		fmt.Print(n, "")
		for _, u := range uniques {
			if u == n {
				continue loop
			}
		}
		uniques[f] = n
		f++
	}
	fmt.Println("\n\nuniques:", uniques)

}
