package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Hello World")
	origSize := 3
	sli := make([]int, origSize)
	var val string
	idx := 0
	for {
		fmt.Println("Please enter a number")

		fmt.Scanln(&val)
		if val == "X" {
			break
		}
		i, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("invalid integer")
		} else {
			// TODO: fix placement!!
			if idx < origSize {
				sli[idx] = i
			} else {
				sli = append(sli, i)
			}
			sort.Slice(sli, func(i, j int) bool {
				return sli[i] < sli[j]
			})

			fmt.Println(sli)
		}
		idx++
	}
}
