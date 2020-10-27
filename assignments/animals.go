package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}
func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}
func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {

	cow := &Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := &Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := &Animal{food: "mice", locomotion: "slither", noise: "hiss"}

	var animalIn string
	var actionIn string
	for {
		fmt.Println("Animal?")
		fmt.Print("> ")
		fmt.Scanln(&animalIn)

		var curAnimal *Animal
		switch {
		case animalIn == "cow":
			curAnimal = cow
		case animalIn == "bird":
			curAnimal = bird
		case animalIn == "snake":
			curAnimal = snake
		default:
			fmt.Println("Invalid input")
			continue
		}

		fmt.Println("Action?")
		fmt.Print("> ")
		fmt.Scanln(&actionIn)
		switch {
		case actionIn == "eat":
			curAnimal.Eat()
		case actionIn == "move":
			curAnimal.Move()
		case actionIn == "speak":
			curAnimal.Speak()
		default:
			fmt.Println("Invalid input")
			continue
		}

	}

}
