package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	FOUND := "Found"
	NOT_FOUND := "Not Found"

	fmt.Println("Please enter a string")
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')

	sl := strings.TrimSpace(strings.ToLower(line))

	i := strings.IndexByte(sl, 'i')
	a := strings.IndexByte(sl, 'a')
	n := strings.LastIndexByte(sl, 'n')

	if i < 0 || a < 0 || n < 0 {
		fmt.Println(NOT_FOUND)
		return
	}

	if i == 0 && i < a && a < n && n == len(sl)-1 {
		fmt.Println(FOUND)
		return
	}
	fmt.Println(NOT_FOUND)
	return

}
