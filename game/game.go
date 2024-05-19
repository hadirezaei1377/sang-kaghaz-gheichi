package game

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomChoice() string {
	choices := []string{"rock", "paper", "scissors"}
	return choices[rand.Intn(len(choices))]
}

func PlayRound(playerChoice, computerChoice string) string {
	if playerChoice == computerChoice {
		return "It's a tie!"
	}
	switch playerChoice {
	case "rock":
		if computerChoice == "scissors" {
			return "You win!"
		}
	case "paper":
		if computerChoice == "rock" {
			return "You win!"
		}
	case "scissors":
		if computerChoice == "paper" {
			return "You win!"
		}
	}
	return "You lose!"
}
