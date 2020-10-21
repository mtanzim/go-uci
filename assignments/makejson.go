package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	var personMap map[string]string
	personMap = make(map[string]string)
	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please enter a name")
		name, _ := in.ReadString('\n')

		fmt.Println("Please enter an address")
		address, _ := in.ReadString('\n')

		personMap["name"] = strings.Trim(name, "\n")
		personMap["address"] = strings.Trim(address, "\n")
		jsonPerson, _ := json.Marshal(personMap)
		fmt.Println(string(jsonPerson))

	}

}
