package main

import "fmt"

func main() {

	var (
		names     []string
		distances []int
		data      []byte
		ratios    []float64
		alives    []bool
	)
	names = []string{"diana", "mehmet", "sonja"}
	distances = []int{110, 144, 180, 200, 200}
	data = []byte{2, 3, 4, 5, 6}
	ratios = []float64{13.4, 15.5}
	alives = []bool{true, true, true, true}

	fmt.Printf("names    : %T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data     : %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios   : %T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives   : %T %d %t\n", alives, len(alives), alives == nil)
	if len(distances) == len(data) {
		fmt.Println("The lenght of the distance and data slices are same ")
	}
}
