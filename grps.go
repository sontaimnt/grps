package main

import (
	"fmt"
	"math/rand"
	"os"
)

var game_choice int
var comp_choice int
var user_choice int
var user_yn int
var player1name string
var player2name string
var player1_choice int
var player2_choice int

func play_with_computer() {
	fmt.Println("[1] rock")
	fmt.Println("[2] paper")
	fmt.Println("[3] scissor")
	fmt.Printf("enter your choice:- ")
	fmt.Scanf("%d", &user_choice)
	comp_choice = rand.Intn(3)
	switch user_choice {
	case 1:
		if user_choice == comp_choice {
			fmt.Println("this game is a draw")
		} else if comp_choice == 2 {
			fmt.Println("you lost the game")
		} else {
			fmt.Println("you won the game")
		}
	case 2:
		if user_choice == comp_choice {
			fmt.Println("this game is a draw")
		} else if comp_choice == 3 {
			fmt.Println("you lost the game")
		} else {
			fmt.Println("you won the game")
		}
	case 3:
		if user_choice == comp_choice {
			fmt.Println("this game is a draw")
		} else if comp_choice == 1 {
			fmt.Println("you lost the game")
		} else {
			fmt.Println("you won the game")
		}
	}
	y_n()
}

func play_with_friend() {
	fmt.Printf("Hey Player 1, enter your name:- ")
	fmt.Scanf("%s", &player1name)
	fmt.Printf("Hey Player 2, enter your name:- ")
	fmt.Scanf("%s", &player2name)
	fmt.Println("[1] rock")
	fmt.Println("[2] paper")
	fmt.Println("[3] scissor")
	fmt.Printf("%s", player1name+" enter your choice:- ")
	fmt.Scanf("%d", &player1_choice)
	fmt.Printf("%s", player2name+" enter your choice:- ")
	fmt.Scanf("%d", &player2_choice)
	if player1_choice == player2_choice {
		fmt.Println("it is a draw")
	} else if player1_choice == 1 {
		if player2_choice == 2 {
			fmt.Println(player1name + " has lost the game")
		} else {
			fmt.Println(player1name + "has won the game")
		}
	} else if player1_choice == 2 {
		if player2_choice == 3 {
			fmt.Println(player1name + " has lost the game")
		} else {
			fmt.Println(player1name + "has won the game")
		}
	} else {
		if player2_choice == 1 {
			fmt.Println(player1name + " has lost the game")
		} else {
			fmt.Println(player1name + "has won the game")
		}
	}
	y_n()
}

func y_n() {
	fmt.Println("[1] yes")
	fmt.Println("[2] no")
	fmt.Printf("would you continue to play? ")
	fmt.Scanf("%d", &user_yn)
	switch user_yn {
	case 1:
		main()
	case 2:
		os.Exit(0)
	default:
		y_n()
	}
}

func main() {
	fmt.Println("Welcome to grps")
	fmt.Println("[1] play with computer ")
	fmt.Println("[2] play with friend")
	fmt.Println("[3] exit")
	fmt.Printf("enter your choice:- ")
	fmt.Scanf("%d", &game_choice)
	switch game_choice {
	case 1:
		play_with_computer()
	case 2:
		play_with_friend()
	case 3:
		os.Exit(0)
	default:
		fmt.Println("error invalid choice")
		os.Exit(1)
	}
}
