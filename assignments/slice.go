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
		iVal, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Invalid integer")
		} else {
			// TODO: fix placement!!
			if idx < origSize-1 {
				sli[idx] = iVal
			} else if idx == origSize-1 {
				if iVal != 0 {
					zeroAt := 0
					for k := 0; k < origSize; k++ {
						if sli[k] == 0 {
							zeroAt = k
							break
						}
					}
					sli[zeroAt] = iVal
				} else {
					// do nothing
				}

			} else {
				sli = append(sli, iVal)
			}
			sort.Slice(sli, func(i, j int) bool {
				return sli[i] < sli[j]
			})

			fmt.Println(sli)
		}
		idx++
	}
}
