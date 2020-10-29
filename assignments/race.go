package main

import (
	"fmt"
	"time"
)

func incrValue(x *int, val int) {
	*x = *x + val
	return
}

func main() {
	x := 0
	val1 := 3
	val2 := 9
	for {
		go incrValue(&x, val1)
		go incrValue(&x, val2)
		fmt.Println("Expect: ", val1+val2)
		fmt.Println("Actual: ", x)
		time.Sleep(500 * time.Millisecond)
		x = 0
	}

}
