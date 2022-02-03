package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	namesC := strings.Split(namesA, ",")
	sort.Strings(namesC)
	sort.Strings(namesB)

	fmt.Printf("%q\n", namesC)
	fmt.Printf("%q\n", strings.Split(namesB[1], ","))

}
