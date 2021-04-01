package termsnake

import (
	"math/rand"
	"time"
)

func (g *Game) min() int {
	if g.boardHeight < g.boardWidth {
		return g.boardHeight - 1
	}
	return g.boardWidth - 1
}

func (g *Game) getRand() (int, int) {
	rand.Seed(time.Now().UnixNano())
	rand1 := rand.Intn(g.min() + 1)

	rand.Seed(time.Now().UnixNano())
	rand2 := rand.Intn(g.min() + 1)

	return rand2, rand1
}

func (g *Game) getInstructLoc() (int, int) {
	var boardEndX int = 2 + g.boardWidth*2
	var boardEndY int = 3

	return boardEndX + 2, boardEndY
}
