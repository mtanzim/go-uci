package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(i int, sli []int) error {

	if i < 0 || i > len(sli)-1 {
		return errors.New("Invalid slice indices")
	}
	sli[i], sli[i+1] = sli[i+1], sli[i]
	return nil
}

func BubbleSort(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				Swap(j, numbers)
			}
		}

	}
}

func GetNumbers() []int {
	fmt.Println("Enter up to 10 integers, seperated by a space")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	numbersText := strings.Trim(text, "\n")
	numbersStr := strings.Split(numbersText, " ")
	if len(numbersStr) > 10 {
		panic(errors.New("Up to 10 numbers are allowed"))
	}
	numbers := make([]int, len(numbersStr))
	for i, numStr := range numbersStr {
		numbers[i], _ = strconv.Atoi(numStr)
	}
	return numbers
}

func main() {

	numbers := GetNumbers()
	fmt.Println("Unsorted: ")
	fmt.Println(numbers)
	BubbleSort(numbers)
	fmt.Println("Sorted: ")
	fmt.Println(numbers)

}
