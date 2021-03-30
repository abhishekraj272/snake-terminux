package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	board       [][]int
	score       int
	length      int
	round       int
	headX       int
	headY       int
	tailX       int
	tailY       int
	boardHeight int
	boardWidth  int
	foodX       int
	foodY       int
	direction   string
}

func NewGame() *Game {
	g := new(Game)

	var boardHeight int
	var boardWidth int

	fmt.Scan(boardHeight)
	fmt.Scan(boardWidth)

	g.boardHeight = boardHeight
	g.boardWidth = boardWidth
	g.direction = "right"

	return g
}

func (g *Game) getFood() {
	rand.Seed(time.Now().UnixNano())

	i, j := g.getRand()

	// 1 in game array means snake.
	if g.board[i][j] == 1 {
		g.getFood()
	} else {
		// 2 in game array means food.
		g.board[i][j] = 2
		g.foodX = i
		g.foodY = j
	}
}

func (g *Game) checkFood() {
	if g.foodX == g.headX && g.foodY == g.headY {
		g.score++
		g.length++
		g.board[g.foodX][g.foodY] = 0

		switch g.direction {

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

		g.getFood()
	}
}

func (g *Game) checkMove() bool {
	switch g.direction {

	case "up":
		if g.board[g.headX-1][g.headY] == 1 || g.headX == 0 {
			return false
		}
		return true

	case "right":
		if g.board[g.headX][g.headY+1] == 1 || g.headY == g.boardWidth {
			return false
		}
		return true

	case "down":
		if g.board[g.headX+1][g.headY] == 1 || g.headX == g.boardHeight {
			return false
		}
		return true

	case "left":
		if g.board[g.headX][g.headY-1] == 1 || g.headY == 0 {
			return false
		}
		return true

	default:
		panic("OMG!! Direction is wrong")
	}

}

func (g *Game) resetGame() {
	g.board = make([][]int, g.boardHeight)

	for i := 0; i < g.boardHeight; i++ {
		g.board[i] = make([]int, g.boardWidth)

		for j := 0; j < g.boardWidth; j++ {
			g.board[i][j] = 0
		}
	}

	g.score = 0
	g.length = 1
	g.round = 0
	g.headX = 0
	g.headY = 0
	g.tailX = 0
	g.tailY = 0

	g.getFood()
}

func (g *Game) moveUp() {
	g.headX--
	g.board[g.headX][g.headY] = 1
	g.board[g.tailX][g.tailY] = 0
	g.getTail()
}

func (g *Game) moveRight() {
	g.headY++
	g.board[g.headX][g.headY] = 1
	g.board[g.tailX][g.tailY] = 0
	g.getTail()
}

func (g *Game) moveDown() {
	g.headX++
	g.board[g.headX][g.headY] = 1
	g.board[g.tailX][g.tailY] = 0
	g.getTail()
}

func (g *Game) moveLeft() {
	g.headY--
	g.board[g.headX][g.headY] = 1
	g.board[g.tailX][g.tailY] = 0
	g.getTail()
}
