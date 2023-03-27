package main

import (
	"fmt"
	"math/rand"
)

const (
	Red    = "red"
	Green  = "green"
	Yellow = "yellow"
	Blue   = "blue"
)

const (
	NumPlayers       = 2
	NumTokens        = 4
	NumSpace         = 52
	StartSpaceRed    = 0
	StartSpaceGreen  = 13
	StartSpaceYellow = 26
	StartSpaceBlue   = 39
)

type Token struct {
	Color string
	Space int
}

type Player struct {
	Name   string
	Color  string
	Tokens []*Token
}

type Game struct {
	Players []*Player
	Board   []string
	Die     *rand.Rand
}

func NewGame() *Game {
	redTokens := make([]*Token, NumTokens)
	greenTokens := make([]*Token, NumTokens)
	yellowTokens := make([]*Token, NumTokens)
	blueTokens := make([]*Token, NumTokens)

	for i := 0; i < NumTokens; i++ {
		redTokens[i] = &Token{Color: Red, Space: -1}
		greenTokens[i] = &Token{Color: Red, Space: -1}
		yellowTokens[i] = &Token{Color: Red, Space: -1}
		blueTokens[i] = &Token{Color: Red, Space: -1}
	}

	players := []*Player{
		{Name: "player 1", Color: Red, Tokens: redTokens},
		{Name: "player 2", Color: Green, Tokens: greenTokens},
	}

	board := make([]string, NumSpace)
	for i := 0; i < NumSpace; i++ {
		board[i] = ""
	}

	game := &Game{
		Players: players,
		Board:   board,
		Die:     rand.New(rand.NewSource(99)),
	}
	return game
}

func (g *Game) RollDie() int {
	return g.Die.Intn(6) + 1
}

func (g *Game) Play() {
	for {
		for i := range g.Players {
			player := g.Players[i]

			fmt.Printf("Player %d's turn\n", i+1)
			fmt.Println("Press enter to roll the die")
			fmt.Scanln()

			roll := g.RollDie()
			fmt.Printf("You rolled a %d\n", roll)

			var tokenIndex int
			fmt.Println("Select a token to move")
			fmt.Scanln(&tokenIndex)
			token := player.Tokens[tokenIndex-1]

			if token.Space < 0 {
				if roll == 6 {
					token.Space = getStartSpace(player.Color)
					g.Board[token.Space] = token.Color
				}
			} else {
				g.Board[token.Space] = ""

				nextSpace := token.Space + roll
				if nextSpace > 51 {
					nextSpace = 51 - (nextSpace - 51)
				}

				if g.Board[nextSpace] == "" {
					token.Space = nextSpace
					g.Board[nextSpace] = token.Color
				} else {
					for _, otherPlayer := range g.Players {
						if otherPlayer.Color == g.Board[nextSpace] {
							otherToken := otherPlayer.Tokens[0]
							otherToken.Space = getStartSpace(otherPlayer.Color)
							g.Board[otherToken.Space] = otherToken.Color
						}
					}
					token.Space = nextSpace
					g.Board[nextSpace] = token.Color
				}
			}

			if isWinner(player) {
				fmt.Printf("Player %d has won the game!\n", i+1)
				return
			}

			fmt.Println(g.Board)
		}
	}
}

func getStartSpace(color string) int {
	switch color {
	case Red:
		return StartSpaceRed
	case Green:
		return StartSpaceGreen
	case Yellow:
		return StartSpaceYellow
	case Blue:
		return StartSpaceBlue
	default:
		return -1
	}
}

func isWinner(player *Player) bool {
	for _, token := range player.Tokens {
		if token.Space < 51 {
			return false
		}
	}
	return true
}

func main() {
	game := NewGame()
	game.Play()
}
