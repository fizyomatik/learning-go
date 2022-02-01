package main

import (
	"fmt"
	"os"
	"strings"
)

const bro = "bir berber bir berbere gel beraber berber dükkani acalim demis"

func main() {
	words := strings.Fields(bro)
	bra := os.Args[1:]
moin:
	for _, q := range bra {
		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d : %q\n", i+1, w)
				continue moin
			}
		}
	}

}
