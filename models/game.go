package models

type Game struct {
	ID            uint `gorm:"primaryKey"`
	PlayerName    string
	PlayerScore   int
	ComputerScore int
}
