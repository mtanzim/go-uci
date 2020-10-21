package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type person struct {
	fname string
	lname string
}

const (
	maxLength = 20
)

func (n *person) Set(first string, last string) {
	n.fname = first
	n.lname = last
	if len(first) > maxLength {
		n.fname = first[:maxLength]
	}
	if len(last) > maxLength {
		n.lname = last[:maxLength]
	}
}

func main() {

	persons := make([]person, 0)
	fmt.Println("Please provide filename")
	var filename string
	fmt.Scanln(&filename)
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fullname := scanner.Text()
		names := strings.Split(fullname, " ")
		newPerson := &person{}
		newPerson.Set(names[0], names[1])
		persons = append(persons, *newPerson)

	}

	for _, p := range persons {
		fmt.Println(p.fname, p.lname)
	}

	f.Close()

}
