package main

import (
	"fmt"
	"math/rand"
)

var comp_choice int
var user_choice int

func main() {
	fmt.Println("Welcome to Rock Paper Scissors")
	fmt.Println("[1] rock")
	fmt.Println("[2] paper")
	fmt.Println("[3] scissor")
	fmt.Printf("Enter your choice:- ")
	fmt.Scanf("%d", &user_choice)
	comp_choice := rand.Intn(3)
	switch user_choice {
	case 1:
		if comp_choice == user_choice {
			fmt.Println("It is draw")
		} else if comp_choice == 2 {
			fmt.Println("You lost")
		} else {
			fmt.Println("You won")
		}
	case 2:
		if comp_choice == user_choice {
			fmt.Println("It is draw")
		} else if comp_choice == 3 {
			fmt.Println("You lost")
		} else {
			fmt.Println("You won")
		}
	case 3:
		if comp_choice == user_choice {
			fmt.Println("It is draw")
		} else if comp_choice == 1 {
			fmt.Println("You lost")
		} else {
			fmt.Println("You won")
		}
	default:
		fmt.Println("Invalid choice")
	}
}
