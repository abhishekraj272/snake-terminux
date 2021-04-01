package termsnake

import (
	"fmt"
	"math/rand"
	"time"
)

// Game stores all the game data
type Game struct {
	board       [][]int8
	score       int
	length      int
	Round       int
	State       int
	headX       int
	headY       int
	tailX       int
	tailY       int
	boardHeight int
	boardWidth  int
	foodX       int
	foodY       int
	Direction   string
}

// NewGame creates a new Game instance and returns it
func NewGame() *Game {
	g := new(Game)

	var boardHeight int
	var boardWidth int
	fmt.Println("Enter Height of Board: ")
	fmt.Scan(&boardHeight)

	fmt.Println("Enter Width of Board: ")
	fmt.Scan(&boardWidth)

	g.boardHeight = boardHeight
	g.boardWidth = boardWidth
	g.Direction = "right"
	g.resetGame()

	return g
}

//  resetGame resets to game to its initial stage
//  and assigns default values to game variables
func (g *Game) resetGame() {

	g.board = make([][]int8, g.boardHeight)

	for i := 0; i < g.boardHeight; i++ {
		g.board[i] = make([]int8, g.boardWidth)

		for j := 0; j < g.boardWidth; j++ {
			g.board[i][j] = 0
		}
	}
	g.board[0][0] = 1
	g.score = 0
	g.length = 1
	g.Round = 0
	g.headX = 0
	g.headY = 0
	g.tailX = 0
	g.tailY = 0
	g.State = 1
	g.GetFood()
}

// GetFood spawns a food on the board
func (g *Game) GetFood() {
	rand.Seed(time.Now().UnixNano())

	i, j := g.getRand()

	// 1 in game array means snake.
	if g.board[i][j] == 1 {
		// Check if food will spawn on snake
		g.GetFood()
	} else {
		// 2 in game array means food.
		g.board[i][j] = 2
		g.foodX = i
		g.foodY = j
	}
}

// CheckFood checks if snake head is on the food
func (g *Game) CheckFood() {
	if g.foodX == g.headX && g.foodY == g.headY {
		g.score++
		g.length++
		g.board[g.foodX][g.foodY] = 1

		switch g.Direction {

		case "up":
			g.tailX++
			g.board[g.tailX][g.tailY] = 1

		case "right":
			g.tailY--
			g.board[g.tailX][g.tailY] = 1

		case "down":
			g.tailX--
			g.board[g.tailX][g.tailY] = 1

		case "left":
			g.tailY++
			g.board[g.tailX][g.tailY] = 1

		default:
			panic("OMG, Direction is wrong")
		}

		g.GetFood()
	}
}

// CheckMove checks if the move of snake is valid, it returns
// true if valid else false
func (g *Game) CheckMove() bool {
	switch g.Direction {

	case "up":
		if g.headX == 0 || g.board[g.headX-1][g.headY] == 1 {
			return false
		}
		return true

	case "right":
		if g.headY == g.boardHeight-1 || g.board[g.headX][g.headY+1] == 1 {
			return false
		}
		return true

	case "down":
		if g.headX == g.boardWidth-1 || g.board[g.headX+1][g.headY] == 1 {
			return false
		}
		return true

	case "left":
		if g.headY == 0 || g.board[g.headX][g.headY-1] == 1 {
			return false
		}
		return true

	default:
		panic("OMG!! Direction is wrong")
	}

}

// MoveUp moves the snake horizontally up
func (g *Game) MoveUp() {
	g.headX--
	g.board[g.headX][g.headY] = 1
	g.board[g.tailX][g.tailY] = 0
	g.getTail()
}

// MoveRight moves the snake vertically right
func (g *Game) MoveRight() {
	g.headY++
	g.board[g.headX][g.headY] = 1
	g.board[g.tailX][g.tailY] = 0
	g.getTail()
}

// MoveDown moves the snake horizontally down
func (g *Game) MoveDown() {
	g.headX++
	g.board[g.headX][g.headY] = 1
	g.board[g.tailX][g.tailY] = 0
	g.getTail()
}

// MoveLeft moves the snake vertically left
func (g *Game) MoveLeft() {
	g.headY--
	g.board[g.headX][g.headY] = 1
	g.board[g.tailX][g.tailY] = 0
	g.getTail()
}

// getTail fetches the latest location of the tail after a move
func (g *Game) getTail() {
	if g.length == 1 {
		g.tailX = g.headX
		g.tailY = g.headY
		return
	}

	if g.tailX != 0 && g.board[g.tailX-1][g.tailY] == 1 {
		g.tailX--
		return
	}

	if g.tailX != g.boardHeight-1 && g.board[g.tailX+1][g.tailY] == 1 {
		g.tailX++
		return
	}

	if g.tailY != 0 && g.board[g.tailX][g.tailY-1] == 1 {
		g.tailY--
		return
	}

	if g.tailY != g.boardWidth-1 && g.board[g.tailX][g.tailY+1] == 1 {
		g.tailY++
		return
	}

}
