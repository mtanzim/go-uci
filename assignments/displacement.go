package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInitVars() (float64, float64, float64) {
	fmt.Println("Enter acceleration, initial velocity, and initial displacement seperated by a space")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	inText := strings.Trim(text, "\n")
	inStr := strings.Split(inText, " ")
	if len(inStr) != 3 {
		panic(errors.New("Only 3 variables are allowed"))
	}
	inputVars := make([]float64, len(inStr))
	for i, iVar := range inStr {
		var err error
		inputVars[i], err = strconv.ParseFloat(iVar, 64)
		if err != nil {
			panic(err)
		}
	}
	acc, iVel, iDisp := inputVars[0], inputVars[1], inputVars[2]
	return acc, iVel, iDisp
}

func GetTime() float64 {
	fmt.Println("Enter time")
	var text string
	fmt.Scanln(&text)
	t, err := strconv.ParseFloat(text, 64)
	if err != nil {
		panic(err)
	}
	return t
}

func GenDisplaceFn(acc, iVel, iDisp float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*acc*t*t + iVel*t + iDisp
	}
}

func main() {

	acc, iVel, iDisp := GetInitVars()
	fn := GenDisplaceFn(acc, iVel, iDisp)
	for {
		t := GetTime()
		d := fn(t)
		fmt.Println("Displacement is", d)
	}

}
