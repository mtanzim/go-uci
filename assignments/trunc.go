package main

import (
	"fmt"
)

func main() {
	var a float64
	fmt.Println("Please enter a number")
	fmt.Scan(&a)
	fmt.Println("Read %d", int64(a))
}
