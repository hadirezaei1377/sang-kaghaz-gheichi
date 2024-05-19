package main

import (
	"fmt"
	"math/rand"
	"sang-kaghaz-gheichi/database"
	"sang-kaghaz-gheichi/game"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	redisClient := database.NewRedisClient()
	db := database.NewSQLiteDB()

	player := game.NewPlayer("User")
	computer := game.NewPlayer("Computer")

	// Load previous scores
	player.Score = game.LoadScore(redisClient, player.Name)
	computer.Score = game.LoadScore(redisClient, computer.Name)
	fmt.Printf("Previous Scores: You %d - Computer %d\n", player.Score, computer.Score)

	for player.Score < 3 && computer.Score < 3 {
		fmt.Println("Enter your choice (rock, paper, scissors):")
		var choice string
		fmt.Scanln(&choice)

		computerChoice := game.RandomChoice()

		result := game.PlayRound(choice, computerChoice)
		fmt.Println("Computer chose:", computerChoice)
		fmt.Println(result)

		if result == "You win!" {
			player.Score++
		} else if result == "You lose!" {
			computer.Score++
		}

		fmt.Printf("Score: You %d - Computer %d\n", player.Score, computer.Score)
		game.SaveScore(redisClient, player, computer)
	}

	if player.Score == 3 {
		fmt.Println("Congratulations! You win the game.")
	} else {
		fmt.Println("Sorry! You lose the game.")
	}

	game.SaveGame(db, player, computer)

	// Reset scores in Redis for next game
	game.ResetScores(redisClient, player.Name, computer.Name)
}
