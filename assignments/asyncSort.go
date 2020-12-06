package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func getNumbers() []int {
	fmt.Println("Enter a series of integers, seperated by a space")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	numbersText := strings.Trim(text, "\n")
	numbersStr := strings.Split(numbersText, " ")
	numbers := make([]int, len(numbersStr))
	for i, numStr := range numbersStr {
		numbers[i], _ = strconv.Atoi(numStr)
	}
	return numbers
}

func splitNumbers(numbers []int) [][]int {
	sizeOfEach := len(numbers) / 4
	if sizeOfEach < 1 {
		return [][]int{numbers}
	}

	return [][]int{
		numbers[:sizeOfEach],
		numbers[sizeOfEach : sizeOfEach*2],
		numbers[sizeOfEach*2 : sizeOfEach*3],
		numbers[sizeOfEach*3:]}

}

func mergePartitionPairs(a []int, b []int) []int {
	temp := make([]int, 0)

	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			temp = append(temp, a[i])
			i++
		} else {
			temp = append(temp, b[j])
			j++
		}
	}

	for i < len(a) {
		temp = append(temp, a[i])
		i++
	}
	for j < len(b) {
		temp = append(temp, b[j])
		j++
	}

	return temp

}

func mergeAllPartitions(partitions [][]int) []int {
	temp := partitions[0]
	for i := 1; i < len(partitions); i++ {
		temp = mergePartitionPairs(temp, partitions[i])
	}
	return temp
}

func sortEachPartition(numbers []int, jobs <-chan int, results chan<- bool) {

	for j := range jobs {
		fmt.Println("Started job", j)
		fmt.Println("numbers", numbers)
		time.Sleep(time.Second * 3)
		sort.Ints(numbers)
		results <- sort.IntsAreSorted(numbers)
		fmt.Println("Completed job", j)
	}
}

// https://gobyexample.com/worker-pools
func sortPartitions(partitions [][]int) {
	numPartitions := len(partitions)
	jobs := make(chan int, numPartitions)
	results := make(chan bool, numPartitions)
	for i := 0; i < numPartitions; i++ {
		go sortEachPartition(partitions[i], jobs, results)
	}
	for j := 0; j < numPartitions; j++ {
		jobs <- j
	}
	close(jobs)
	// wait for results
	for k := 0; k < numPartitions; k++ {
		isSorted := <-results
		if !isSorted {
			panic("Not sorted")
		}
	}
}

func main() {

	numbers := getNumbers()
	fmt.Println("Original array:")
	fmt.Println(numbers)

	partitions := splitNumbers(numbers)
	fmt.Println("Unsorted partitions:")
	fmt.Println(partitions)

	sortPartitions(partitions)
	fmt.Println("Sorted partitions:")
	fmt.Println(partitions)

	sortedNumbers := mergeAllPartitions(partitions)
	if !sort.IntsAreSorted(sortedNumbers) {
		panic("Array was not sorted!")
	}
	fmt.Println("Sorted array:")
	fmt.Println(sortedNumbers)

}
