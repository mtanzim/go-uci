package main

import (
	"fmt"
	"sync"
	"time"
)

func incrValue(x *int, val int, wg *sync.WaitGroup) {
	// TODO: this still creates a race condition
	// since *x is non-deterministic
	defer wg.Done()
	*x = *x + val
	return
}

func main() {
	x := 0
	numRoutines := 50

	vals := make([]int, numRoutines)
	for i := 0; i < numRoutines; i++ {
		vals[i] = 2 * numRoutines
	}
	var wg sync.WaitGroup
	for {
		wg.Add(numRoutines)
		for i := 0; i < numRoutines; i++ {
			go incrValue(&x, vals[i], &wg)
		}
		wg.Wait()
		sum := 0
		for _, val := range vals {
			sum += val
		}
		fmt.Println("Expect: ", sum)
		fmt.Println("Actual: ", x)
		if sum != x {
			panic("Values must equal each other!")
		}
		time.Sleep(500 * time.Millisecond)
		x = 0
	}

}
