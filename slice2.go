package main

import "fmt"

func main() {
	var names []string
	var distance []int
	var data []uint8
	var exchange []float64
	var status []bool
	if len(names) == 0 {
		fmt.Print("true")
	}

	fmt.Printf("Names    :%T %d %t\n", names, len(names), names == nil)
	fmt.Printf("Distance     :%T %d %t\n", distance, len(distance), distance == nil)
	fmt.Printf("Data         :%T %d %t\n", data, len(distance), distance == nil)
	fmt.Printf("Exchange     :%T %d %t\n", exchange, len(exchange), exchange == nil)
	fmt.Printf("Status       :%T %d %t\n", status, len(status), status == nil)
}
