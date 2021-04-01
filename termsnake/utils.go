package termsnake

import (
	"math/rand"
	"time"
)

// min returns the minimum of board Width and Height
func (g *Game) min() int {
	if g.boardHeight < g.boardWidth {
		return g.boardHeight
	}
	return g.boardWidth
}

// getRand returns two random integer between 0 and min
func (g *Game) getRand() (int, int) {
	rand.Seed(time.Now().UnixNano())
	rand1 := rand.Intn(g.min())

	rand.Seed(time.Now().UnixNano())
	rand2 := rand.Intn(g.min())

	return rand2, rand1
}

// getBoardEnd returns coords of board edges
func (g *Game) getBoardEnd() (int, int) {
	var boardEndX int = 2 + g.boardWidth*2
	var boardEndY int = 3

	return boardEndX + 2, boardEndY
}
