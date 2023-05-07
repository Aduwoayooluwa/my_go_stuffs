package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// creating a tictactoe app using golang
	fmt.Println("Hello")
	options := [3]string{"rock", "paper", "scissors"}

	randomIndex := rand.Intn(3)
	computerGame := options[randomIndex]
	fmt.Println(computerGame)

	var user string
	fmt.Scanln(&user)
	fmt.Println("Enter your input")

	if user == computerGame {
		fmt.Println("Correct")
	} else {
		fmt.Println("Wrong")
	}

	if user == "rock" && computerGame =="scissors" {
		fmt.Println("User wins")
	}

	else if user == "paper" && computerGame =="rock" {
		fmt.Println("Computer wins")
	}

	else if user == "rock" && computerGame =="paper" {
		fmt.Println("User wins")
	}
	else {
		fmt.Println("Not working")
	}

}
