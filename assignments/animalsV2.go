package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Speak()
	Move()
}

type Cow struct {
	name string
}
type Bird struct {
	name string
}
type Snake struct {
	name string
}

func (_ Cow) Eat() {
	fmt.Println("grass")
}
func (_ Bird) Eat() {
	fmt.Println("worms")
}
func (_ Snake) Eat() {
	fmt.Println("mice")
}

func (_ Cow) Speak() {
	fmt.Println("moo")
}
func (_ Bird) Speak() {
	fmt.Println("peep")
}
func (_ Snake) Speak() {
	fmt.Println("hss")
}

func (_ Cow) Move() {
	fmt.Println("walk")
}
func (_ Bird) Move() {
	fmt.Println("fly")
}
func (_ Snake) Move() {
	fmt.Println("slither")
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	animalMap := make(map[string]Animal)
	for {
		fmt.Println("")
		fmt.Println("Command?")
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		commandsText := strings.Trim(text, "\n")
		commands := strings.Split(commandsText, " ")
		cmd, animalName := commands[0], commands[1]

		switch cmd {
		case "newanimal":
			animalType := commands[2]
			var newAnimal Animal
			switch animalType {
			case "cow":
				newAnimal = Cow{animalName}
			case "bird":
				newAnimal = Bird{animalName}
			case "snake":
				newAnimal = Snake{animalName}
			default:
				fmt.Println("Invalid animal type")
				continue
			}
			animalMap[animalName] = newAnimal
			fmt.Println("Created it!")

		case "query":
			animalAction := commands[2]
			foundAnimal, ok := animalMap[animalName]
			if !ok {
				fmt.Println("Animal not found!")
				continue
			}
			switch animalAction {
			case "eat":
				foundAnimal.Eat()
			case "move":
				foundAnimal.Move()
			case "speak":
				foundAnimal.Speak()
			default:
				fmt.Println("Invalid action type")
				continue
			}

		default:
			fmt.Println("Invalid command")

		}
	}

}
