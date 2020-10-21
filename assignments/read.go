package main

import (
	"bufio"
	"fmt"
	"os"
)

type person struct {
	fname string
	lname string
}

func main() {

	persons := make([]person, 0)
	newPerson := person{fname: "Tanzim", lname: "Mokammel"}
	persons = append(persons, newPerson)
	fmt.Println(persons)
	fmt.Println("Please provide filename")
	var filename string
	fmt.Scanln(&filename)
	fmt.Println(filename)
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	f.Close()

}
