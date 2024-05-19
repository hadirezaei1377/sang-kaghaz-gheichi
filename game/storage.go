package game

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func SaveScore(client *redis.Client, player *Player, computer *Player) {
	ctx := context.Background()
	client.Set(ctx, player.Name, player.Score, 0)
	client.Set(ctx, computer.Name, computer.Score, 0)
}

func LoadScore(client *redis.Client, name string) int {
	ctx := context.Background()
	score, err := client.Get(ctx, name).Result()
	if err != nil {
		return 0
	}
	s, _ := strconv.Atoi(score)
	return s
}

func ResetScores(client *redis.Client, playerName, computerName string) {
	ctx := context.Background()
	client.Set(ctx, playerName, 0, 0)
	client.Set(ctx, computerName, 0, 0)
}

func SaveGame(db *gorm.DB, player *Player, computer *Player) {
	db.Create(&GameRecord{PlayerName: player.Name, PlayerScore: player.Score, ComputerScore: computer.Score})
}

type GameRecord struct {
	gorm.Model
	PlayerName    string
	PlayerScore   int
	ComputerScore int
}
