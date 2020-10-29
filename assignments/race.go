/*
A race condition is when the outcome of your program is dependent on the non-deterministic nature of interleaving.
This is almost never desired.
Synchronization mechanisms are generally in place to prevent race conditions.
An example of a race condition can be seen from the output of the provided script.

In the script, we have a shared pointer to x, and two goroutines modify it concurrently.
Therefore, the output in the main routine is dependent on the scheduler, and thus non-deterministic.

Expect:  12
Actual:  12
Expect:  12
Actual:  12
Expect:  12
Actual:  0
Expect:  12
Actual:  12
Expect:  12
Actual:  12
Expect:  12
Actual:  3
Expect:  12
Actual:  12
Expect:  12
Actual:  3
Expect:  12
Actual:  12
Expect:  12
Actual:  12
Expect:  12
Actual:  12
*/

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
