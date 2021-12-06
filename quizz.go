package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Quizz")
	fmt.Printf(("Enter your name: "))

	var name string
	var age uint
	// this way it gets input from user
	fmt.Scan(&name)
	fmt.Printf("hello, %v, Welcome to the Game...\n", name)
	fmt.Printf("Enter Your Age: ")
	fmt.Scan(&age)
	if age < 10 {
		fmt.Printf("Sorry kid.. The minimum age for playing this game is 10 but you are %v ", age)
		return
	} else {
		fmt.Printf("Ready when you are!")
	}
	fmt.Printf("\n Exeis faei pote avocado me meli? [Y/n]? ")
	var avocado string
	fmt.Scanln(&avocado)

	if strings.ToLower(avocado) == "y" || avocado == "" {
		fmt.Printf("eisi ouraios!")
		return
	} else if strings.ToLower(avocado) == "n" {
		fmt.Printf("Διεν ξερς τι χανς")
	}
}
